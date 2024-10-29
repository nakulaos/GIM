package service

import (
	"GIM/pkg/common/xants"
	"GIM/pkg/common/xlog"
	"GIM/pkg/common/xmysql"
	"GIM/pkg/constant"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/utils"
	"gorm.io/gorm"
)

func (s *chatService) removeChatMember(u *entity.MysqlUpdate, chatId int64, uidList []int64, chatType pb_enum.CHAT_TYPE) (rowsAffected int64, err error) {
	/*
		注意： 单聊和群聊一样，都是删除聊天成员，只是单聊只有两个成员，群聊有多个成员
		1. 如果是单聊，删除chat对应的两个成员，如果是群聊，删除chat对应的多个成员
		2. 缓存删除chat对应的成员信息
	*/

	var (
		htk      = utils.GetHashTagKey(chatId)
		key1     = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + htk
		key2     = constant.RK_SYNC_CHAT_MEMBER_INFO_HASH + htk
		uidCount = len(uidList)
		uid      int64
	)
	defer func() {
		if err != nil {
			xlog.Warn(err.Error())
		}
	}()

	err = xmysql.Transaction(func(tx *gorm.DB) (err error) {
		rowsAffected, err = s.chatMemberRepo.TxQuitChatMember(tx, u)
		if err != nil {
			return
		}
		switch chatType {
		case pb_enum.CHAT_TYPE_PRIVATE:
			if rowsAffected != 1 {
				err = ERR_CHAT_UPDATE_VALUE_FAILED
				return
			}
			if uidCount != 2 {
				err = ERR_CHAT_REQ_PARAM_ERR
				return
			}
			u.Reset()
			u.SetFilter("chat_id=?", chatId)
			u.SetFilter("owner_id=?", uidList[1])
			u.SetFilter("deleted_ts=?", 0)
			u.Set("status", int32(pb_enum.CHAT_STATUS_NON_CONTACT))
			rowsAffected, err = s.chatMemberRepo.TxQuitChatMember(tx, u)
			if err != nil {
				return
			}
			if rowsAffected != 1 {
				err = ERR_CHAT_UPDATE_VALUE_FAILED
				return
			}
		case pb_enum.CHAT_TYPE_GROUP:
			if int(rowsAffected) != len(uidList) {
				err = ERR_CHAT_UPDATE_VALUE_FAILED
				return
			}
		}
		return
	})
	if err != nil {
		return
	}
	if uidCount == 0 {
		return
	}
	xants.Submit(func() {
		var (
			removes = map[string][]string{}
			sKey    string
			field   string
			slot    = ":0"
		)
		for _, uid = range uidList {
			field = utils.ToString(uid)
			if chatType != pb_enum.CHAT_TYPE_PRIVATE {
				slot = utils.GetChatSlot(uid)
			}
			sKey = key1 + slot
			removes[sKey] = append(removes[sKey], field)
			removes[key2] = append(removes[key2], field)
		}
		terr := s.chatMemberCache.HDelChatMembers(removes)
		if terr != nil {
			xlog.Warnf("remove chat member failed. err: %s,removes: %v", terr.Error(), removes)
			_, _, terr = s.cacheProducer.Push(removes, constant.CONST_MSG_KEY_CACHE_REMOVE_CHAT_MEMBER)
			if terr != nil {
				xlog.Errorf("push remove chat member message failed. err: %s,removes: %v", terr.Error(), removes)
			}
		}
	})
	return
}

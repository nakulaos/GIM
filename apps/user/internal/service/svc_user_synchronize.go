package service

import (
	"GIM/pkg/constant"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_chat_member"
	"GIM/pkg/utils"
	"github.com/spf13/cast"
)

func (s *userService) updateChatMemberCacheInfo(uid int64) (err error) {
	var (
		w       = entity.NewMysqlQuery()
		members []*pb_chat_member.ChatMemberInfo
		member  *pb_chat_member.ChatMemberInfo
		uidStr  = cast.ToString(uid)
		jsonStr string
		key     string
		keys    []string
		vals    []string
		index   int
	)
	w.SetFilter("uid=?", uid)
	w.SetFilter("sync=?", constant.SYNCHRONIZE_USER_INFO)
	members, err = s.chatMemberRepo.ChatMemberList(w)
	if err != nil {
		//r.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		return
	}
	keys = make([]string, len(members))
	vals = make([]string, len(members))
	for index, member = range members {
		jsonStr, err = utils.Marshal(member)
		if err != nil {
			//r.Set(ERROR_CODE_USER_MARSHAL_FAILED, ERROR_USER_MARSHAL_FAILED)
			return
		}
		key = constant.RK_SYNC_CHAT_MEMBER_INFO_HASH + utils.GetHashTagKey(member.ChatId)
		keys[index] = key
		vals[index] = jsonStr
	}
	if len(keys) == 0 {
		return
	}
	// key 是  "CHAT:MEMBER_INFO_HASH:" + chatId 某个聊天的id
	// field 是 uid
	// val 是 ChatMemberInfo 的 序列化
	err = s.chatMemberCache.HSetDistChatMembers(keys, uidStr, vals)
	if err != nil {
		//r.Set(ERROR_CODE_USER_CACHE_CHAT_MEMBER_INFO_FAILED, ERROR_USER_CACHE_CHAT_MEMBER_INFO_FAILED)
	}
	return
}

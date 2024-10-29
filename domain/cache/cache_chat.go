package cache

import (
	"GIM/pkg/common/xlog"
	"GIM/pkg/common/xredis"
	"GIM/pkg/constant"
	"GIM/pkg/proto/pb_chat"
	"GIM/pkg/utils"
)

type ChatCache interface {
	GetGroupChatInfo(chatId int64) (info *pb_chat.ChatInfo, err error)
	SetGroupChatInfo(info *pb_chat.ChatInfo) (err error)
	DelChatInfo(chatId int64) (err error)
}

type chatCache struct {
}

func NewChatCache() ChatCache {
	return &chatCache{}
}

func (c *chatCache) GetGroupChatInfo(chatId int64) (info *pb_chat.ChatInfo, err error) {
	var (
		key = constant.RK_SYNC_GROUP_CHAT_INFO + utils.GetHashTagKey(chatId)
	)
	info = new(pb_chat.ChatInfo)
	err = Get(key, info)
	return
}

func (c *chatCache) SetGroupChatInfo(info *pb_chat.ChatInfo) (err error) {
	var (
		key = constant.RK_SYNC_GROUP_CHAT_INFO + utils.GetHashTagKey(info.ChatId)
	)
	err = Set(key, info, constant.CONST_DURATION_GROUP_CHAT_INFO_SECOND)
	return
}

func (c *chatCache) DelGroupChatInfo(chatId int64) (err error) {
	var (
		key = constant.RK_SYNC_GROUP_CHAT_INFO + utils.GetHashTagKey(chatId)
	)
	err = xredis.Unlink(key)
	return
}

func (c *chatCache) DelChatInfo(chatId int64) (err error) {
	var (
		key = constant.RK_SYNC_GROUP_CHAT_INFO + utils.GetHashTagKey(chatId)
	)
	err = xredis.Unlink(key)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_DELETE_FAILED, ERROR_CACHE_REDIS_DELETE_FAILED, err.Error())
	}
	return
}

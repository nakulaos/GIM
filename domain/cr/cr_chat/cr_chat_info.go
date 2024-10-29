package cr_chat

import (
	"GIM/domain/cache"
	"GIM/domain/po"
	"GIM/domain/repo"
	"GIM/pkg/common/xants"
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_chat"
	"github.com/jinzhu/copier"
)

func GetGroupChatInfo(chatCache cache.ChatCache, chatRepo repo.ChatRepository, chatId int64) (chatInfo *pb_chat.ChatInfo, err error) {
	chatInfo, _ = chatCache.GetGroupChatInfo(chatId)
	if chatInfo.ChatId > 0 {
		return
	}
	var (
		q    = entity.NewMysqlQuery()
		chat *po.Chat
	)
	q.SetFilter("chat_id = ?", chatId)
	chat, err = chatRepo.Chat(q)
	if err != nil {
		return
	}
	if chat.ChatId == 0 {
		err = ERROR_CR_CHAT_QUERY_FAILED
	}
	copier.Copy(chatInfo, chat)
	xants.Submit(func() {
		chatCache.SetGroupChatInfo(chatInfo)
	})
	return
}

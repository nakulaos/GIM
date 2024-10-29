package dto_chat

import "GIM/apps/interfaces/internal/dto/dto_kv"

type EditGroupChatReq struct {
	ChatId int64             `json:"chat_id" binding:"required,gt=0"` // chat ID
	Kvs    *dto_kv.KeyValues `json:"kvs" binding:"required"`          // 更新字段
}

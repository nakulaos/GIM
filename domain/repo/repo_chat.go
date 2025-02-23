package repo

import (
	"GIM/domain/po"
	"GIM/pkg/common/xmysql"
	"GIM/pkg/common/xsnowflake"
	"GIM/pkg/entity"
	"gorm.io/gorm"
)

type ChatRepository interface {
	Create(chat *po.Chat) (err error)
	TxCreate(tx *gorm.DB, chat *po.Chat) (err error)
	Chat(w *entity.MysqlQuery) (chat *po.Chat, err error)
	TxChat(tx *gorm.DB, w *entity.MysqlQuery) (chat *po.Chat, err error)
	UpdateChat(u *entity.MysqlUpdate) (chat *po.Chat, err error)
	TxUpdateChat(tx *gorm.DB, u *entity.MysqlUpdate) (err error)
}

type chatRepository struct {
}

func NewChatRepository() ChatRepository {
	return &chatRepository{}
}

func (r *chatRepository) Create(chat *po.Chat) (err error) {
	if chat.ChatId == 0 {
		chat.ChatId = xsnowflake.NewSnowflakeID()
	}
	db := xmysql.GetDB()
	err = db.Create(chat).Error
	return
}

func (r *chatRepository) TxCreate(tx *gorm.DB, chat *po.Chat) (err error) {
	if chat.ChatId == 0 {
		chat.ChatId = xsnowflake.NewSnowflakeID()
	}
	err = tx.Create(chat).Error
	return
}

func (r *chatRepository) Chat(w *entity.MysqlQuery) (chat *po.Chat, err error) {
	chat = new(po.Chat)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(chat).Error
	return
}

func (r *chatRepository) TxChat(tx *gorm.DB, w *entity.MysqlQuery) (chat *po.Chat, err error) {
	chat = new(po.Chat)
	err = tx.Where(w.Query, w.Args...).Find(chat).Error
	return
}

func (r *chatRepository) UpdateChat(u *entity.MysqlUpdate) (chat *po.Chat, err error) {
	chat = new(po.Chat)
	db := xmysql.GetDB()
	err = db.Model(po.Chat{}).Where(u.Query, u.Args...).Updates(u.Values).Find(chat).Error
	return
}

func (r *chatRepository) TxUpdateChat(tx *gorm.DB, u *entity.MysqlUpdate) (err error) {
	err = tx.Model(po.Chat{}).Where(u.Query, u.Args...).Updates(u.Values).Error
	return
}

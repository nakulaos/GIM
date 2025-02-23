package repo

import (
	"GIM/domain/po"
	"GIM/pkg/common/xmysql"
	"GIM/pkg/entity"
	"gorm.io/gorm"
)

type ChatMessageRepository interface {
	Create(message *po.Message) (err error)
	Message(w *entity.MysqlQuery) (message *po.Message, err error)
	TxMessage(tx *gorm.DB, w *entity.MysqlQuery) (message *po.Message, err error)
	UpdateMessage(u *entity.MysqlUpdate) (err error)
	TxUpdateMessage(tx *gorm.DB, u *entity.MysqlUpdate) (err error)
	HistoryMessages(w *entity.MysqlQuery) (list []*po.Message, err error)
}

type chatMessageRepository struct {
}

func NewChatMessageRepository() ChatMessageRepository {
	return &chatMessageRepository{}
}

func (r *chatMessageRepository) Create(message *po.Message) (err error) {
	db := xmysql.GetDB()
	return db.Create(message).Error
}

func (r *chatMessageRepository) Message(w *entity.MysqlQuery) (message *po.Message, err error) {
	message = new(po.Message)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Find(message).Error
	return
}

func (r *chatMessageRepository) TxMessage(tx *gorm.DB, w *entity.MysqlQuery) (message *po.Message, err error) {
	message = new(po.Message)
	err = tx.Where(w.Query, w.Args...).Find(message).Error
	return
}

func (r *chatMessageRepository) UpdateMessage(u *entity.MysqlUpdate) (err error) {
	db := xmysql.GetDB()
	return db.Model(po.Message{}).Where(u.Query, u.Args...).Updates(u.Values).Error
}

func (r *chatMessageRepository) TxUpdateMessage(tx *gorm.DB, u *entity.MysqlUpdate) (err error) {
	return tx.Model(po.Message{}).Where(u.Query, u.Args...).Updates(u.Values).Error
}

func (r *chatMessageRepository) HistoryMessages(w *entity.MysqlQuery) (list []*po.Message, err error) {
	list = make([]*po.Message, 0)
	db := xmysql.GetDB()
	err = db.Where(w.Query, w.Args...).Order(w.Sort).Find(&list).Error
	return
}

package repo

import (
	"GIM/domain/po"
	"gorm.io/gorm"
)

type RedEnvelopeReceiverRepository interface {
	TxCreateRedEnvelopeReceivers(tx *gorm.DB, list []*po.RedEnvelopeReceiver) (err error)
}

type redEnvelopeReceiverRepository struct {
}

func NewRedEnvelopeReceiverRepository() RedEnvelopeReceiverRepository {
	return &redEnvelopeReceiverRepository{}
}

func (r *redEnvelopeReceiverRepository) TxCreateRedEnvelopeReceivers(tx *gorm.DB, list []*po.RedEnvelopeReceiver) (err error) {
	if len(list) == 0 {
		return
	}
	err = tx.Create(list).Error
	return
}

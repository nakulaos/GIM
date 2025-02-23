package repo

import (
	"GIM/domain/po"
	"GIM/pkg/common/xmysql"
)

type UserLocationRepository interface {
	Save(loc *po.UserLocation) (err error)
}

type userLocationRepository struct {
}

func NewUserLocationRepository() UserLocationRepository {
	return &userLocationRepository{}
}

func (r *userLocationRepository) Save(loc *po.UserLocation) (err error) {
	db := xmysql.GetDB()
	err = db.Save(loc).Error
	return
}

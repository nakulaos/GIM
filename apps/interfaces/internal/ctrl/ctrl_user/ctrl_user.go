package ctrl_user

import (
	"GIM/apps/interfaces/internal/service/svc_user"
)

type UserCtrl struct {
	userService svc_user.UserService
}

func NewUserCtrl(userService svc_user.UserService) *UserCtrl {
	return &UserCtrl{userService: userService}
}

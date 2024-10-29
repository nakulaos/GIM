package do

import (
	"GIM/domain/po"
	"GIM/pkg/proto/pb_auth"
)

type SignUp struct {
	User         *po.User
	Avatar       *po.Avatar
	AccessToken  *pb_auth.Token
	RefreshToken *pb_auth.Token
	Code         int32
	Msg          string
	Err          error
}

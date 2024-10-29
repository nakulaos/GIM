package cr_user

import (
	"GIM/domain/cache"
	"GIM/domain/pdo"
	"GIM/domain/repo"
	"GIM/pkg/entity"
)

func GetOauthUserInfo(oauthUserCache cache.OauthUserCache, userRepo repo.OauthUserRepository, uid int64) (info *pdo.OauthUserInfo, err error) {
	info, _ = oauthUserCache.GetOauthUserInfo(uid)
	if info.Uid > 0 {
		return
	}
	var (
		q = entity.NewMysqlQuery()
	)
	q.SetFilter("uid=?", uid)
	info, err = userRepo.GetOAuthUserInfo(q)
	if err != nil {
		return
	}
	oauthUserCache.SetOauthUserInfo(info)
	return
}

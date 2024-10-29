package dig

import (
	"GIM/apps/auth/internal/config"
	"GIM/apps/auth/internal/server"
	"GIM/apps/auth/internal/server/auth"
	"GIM/apps/auth/internal/service"
	"GIM/business/biz_online"
	"GIM/domain/cache"
	"GIM/domain/repo"
	"go.uber.org/dig"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(auth.NewAuthServer)
	Provide(service.NewAuthService)
	Provide(repo.NewOauthUserRepository)
	Provide(repo.NewAvatarRepository)
	Provide(repo.NewUserRepository)
	Provide(repo.NewChatMemberRepository)
	Provide(cache.NewAuthCache)
	Provide(cache.NewUserCache)
	Provide(cache.NewServerMgrCache)
	Provide(biz_online.NewOnline)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}

func Provide(constructor interface{}, opts ...dig.ProvideOption) {
	err := container.Provide(constructor)
	if err != nil {
		log.Panic(err)
	}
}

package dig

import (
	"GIM/apps/user/internal/config"
	"GIM/apps/user/internal/server"
	"GIM/apps/user/internal/server/user"
	"GIM/apps/user/internal/service"
	"GIM/domain/cache"
	"GIM/domain/repo"
	"go.uber.org/dig"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(user.NewUserServer)
	Provide(service.NewUserService)
	Provide(repo.NewUserRepository)
	Provide(repo.NewAvatarRepository)
	Provide(repo.NewChatMemberRepository)
	Provide(cache.NewChatMemberCache)
	Provide(cache.NewUserCache)
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

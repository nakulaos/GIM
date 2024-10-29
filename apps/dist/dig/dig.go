package dig

import (
	"GIM/apps/dist/internal/config"
	"GIM/apps/dist/internal/server"
	"GIM/apps/dist/internal/server/dist"
	"GIM/apps/dist/internal/service"
	"GIM/domain/cache"
	"go.uber.org/dig"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(dist.NewDistServer)
	Provide(service.NewDistService)
	Provide(cache.NewServerMgrCache)
	Provide(cache.NewConvoCache)
	Provide(cache.NewChatMemberCache)
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

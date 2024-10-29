package dig

import (
	"GIM/apps/cache/internal/config"
	"GIM/apps/cache/internal/server"
	srv_cache "GIM/apps/cache/internal/server/cache"
	"GIM/apps/cache/internal/service"
	"GIM/domain/cache"
	"go.uber.org/dig"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(cache.NewChatMemberCache)
	Provide(service.NewCacheService)
	Provide(srv_cache.NewCacheServer)
	Provide(server.NewServer)
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

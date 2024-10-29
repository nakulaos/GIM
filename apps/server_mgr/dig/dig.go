package dig

import (
	"GIM/apps/server_mgr/internal/config"
	"GIM/apps/server_mgr/internal/server"
	"GIM/apps/server_mgr/internal/server/server_mgr"
	"GIM/apps/server_mgr/internal/service"
	"GIM/domain/cache"
	"go.uber.org/dig"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(server_mgr.NewServerMgrServer)
	Provide(service.NewServerMgrService)
	Provide(cache.NewServerMgrCache)
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

package dig

import (
	"GIM/apps/msg_gateway/internal/config"
	"GIM/apps/msg_gateway/internal/server"
	"GIM/apps/msg_gateway/internal/server/gateway"
	"GIM/apps/msg_gateway/internal/service"
	"GIM/domain/cache"
	"go.uber.org/dig"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(service.NewWsService)
	Provide(gateway.NewGatewayServer)
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

package dig

import (
	"GIM/apps/cloud_msg/internal/config"
	"GIM/apps/cloud_msg/internal/server"
	"GIM/apps/cloud_msg/internal/server/cloud_msg"
	"GIM/apps/cloud_msg/internal/service"
	"go.uber.org/dig"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(cloud_msg.NewCloudMessageServer)
	Provide(service.NewCloudMessageService)
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

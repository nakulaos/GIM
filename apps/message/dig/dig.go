package dig

import (
	"GIM/apps/message/internal/config"
	"GIM/apps/message/internal/server"
	"GIM/apps/message/internal/server/message"
	"GIM/apps/message/internal/service"
	"GIM/domain/cache"
	"go.uber.org/dig"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(message.NewMessageServer)
	Provide(service.NewMessageService)
	Provide(cache.NewChatMemberCache)
	Provide(cache.NewChatMessageCache)

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

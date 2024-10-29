package dig

import (
	"GIM/apps/msg_history/internal/config"
	"GIM/apps/msg_history/internal/server"
	"GIM/apps/msg_history/internal/server/msg_history"
	"GIM/apps/msg_history/internal/service"
	"GIM/domain/cache"
	"GIM/domain/repo"
	"go.uber.org/dig"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(msg_history.NewMessageHistoryServer)
	Provide(service.NewMessageHistoryService)
	Provide(repo.NewChatMessageRepository)
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

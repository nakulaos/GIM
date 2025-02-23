package dig

import (
	"GIM/apps/chat_msg/internal/config"
	"GIM/apps/chat_msg/internal/server"
	"GIM/apps/chat_msg/internal/server/chat_msg"
	"GIM/apps/chat_msg/internal/service"
	"GIM/domain/cache"
	"GIM/domain/mrepo"
	"GIM/domain/repo"
	"go.uber.org/dig"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(chat_msg.NewChatMessageServer)
	Provide(service.NewChatMessageService)
	Provide(repo.NewChatMessageRepository)
	Provide(repo.NewChatMemberRepository)
	Provide(mrepo.NewMessageHotRepository)
	Provide(cache.NewChatMessageCache)
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

package dig

import (
	"GIM/apps/chat/internal/config"
	"GIM/apps/chat/internal/server"
	"GIM/apps/chat/internal/server/chat"
	"GIM/apps/chat/internal/service"
	"GIM/domain/cache"
	"GIM/domain/repo"
	"go.uber.org/dig"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(chat.NewChatServer)
	Provide(service.NewChatService)
	Provide(repo.NewChatRepository)
	Provide(repo.NewChatMemberRepository)
	Provide(repo.NewChatInviteRepository)
	Provide(repo.NewUserRepository)
	Provide(repo.NewAvatarRepository)
	Provide(cache.NewChatCache)
	Provide(cache.NewChatMessageCache)
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

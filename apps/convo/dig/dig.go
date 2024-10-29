package dig

import (
	"GIM/apps/convo/internal/config"
	"GIM/apps/convo/internal/server"
	"GIM/apps/convo/internal/server/convo"
	"GIM/apps/convo/internal/service"
	"GIM/domain/cache"
	"GIM/domain/repo"
	"go.uber.org/dig"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(convo.NewConvoServer)
	Provide(service.NewConvoService)
	Provide(repo.NewChatMemberRepository)
	Provide(cache.NewConvoCache)
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

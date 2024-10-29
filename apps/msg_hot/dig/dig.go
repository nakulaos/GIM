package dig

import (
	"GIM/apps/msg_hot/internal/config"
	"GIM/apps/msg_hot/internal/server"
	"GIM/apps/msg_hot/internal/server/msg_hot"
	"GIM/apps/msg_hot/internal/service"
	"GIM/domain/mrepo"
	"go.uber.org/dig"
	"log/slog"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(mrepo.NewMessageHotRepository)
	Provide(msg_hot.NewMessageHotServer)
	Provide(service.NewMessageHotService)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}

func Provide(constructor interface{}, opts ...dig.ProvideOption) {
	err := container.Provide(constructor)
	if err != nil {
		slog.Warn(err.Error())
	}
}

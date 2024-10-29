package dig

import (
	"GIM/apps/lbs/internal/config"
	"GIM/apps/lbs/internal/server"
	"GIM/apps/lbs/internal/server/lbs"
	"GIM/apps/lbs/internal/service"
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
	Provide(lbs.NewLbsServer)
	Provide(service.NewLbsService)
	Provide(mrepo.NewLbsRepository)
	Provide(repo.NewUserRepository)
	Provide(repo.NewUserLocationRepository)
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

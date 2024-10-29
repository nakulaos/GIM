package main

import (
	"GIM/apps/dist/dig"
	"GIM/apps/dist/internal/config"
	"GIM/pkg/commands"
	"GIM/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}

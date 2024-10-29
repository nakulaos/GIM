package main

import (
	"GIM/apps/server_mgr/dig"
	"GIM/apps/server_mgr/internal/config"
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

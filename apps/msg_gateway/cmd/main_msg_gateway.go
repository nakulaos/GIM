package main

import (
	"GIM/apps/msg_gateway/dig"
	"GIM/apps/msg_gateway/internal/config"
	"GIM/pkg/commands"
	"GIM/pkg/common/xredis"
	"runtime"
)

func init() {
	conf := config.GetConfig()
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}

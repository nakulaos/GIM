package main

import (
	"GIM/apps/msg_hot/dig"
	"GIM/apps/msg_hot/internal/config"
	"GIM/pkg/commands"
	"GIM/pkg/common/xmongo"
	"GIM/pkg/common/xredis"
)

// 弃用
func init() {
	conf := config.GetConfig()
	xmongo.NewMongoClient(conf.Mongo)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}

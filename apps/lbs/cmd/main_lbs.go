package main

import (
	"GIM/apps/lbs/dig"
	"GIM/apps/lbs/internal/config"
	"GIM/pkg/commands"
	"GIM/pkg/common/xmongo"
	"GIM/pkg/common/xmysql"
	"GIM/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xmongo.NewMongoClient(conf.Mongo)
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}

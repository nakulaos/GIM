package main

import (
	"GIM/apps/auth/dig"
	"GIM/apps/auth/internal/config"
	"GIM/pkg/commands"
	"GIM/pkg/common/xmysql"
	"GIM/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}

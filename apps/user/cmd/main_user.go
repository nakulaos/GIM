package main

import (
	"GIM/apps/user/dig"
	"GIM/apps/user/internal/config"
	"GIM/pkg/commands"
	"GIM/pkg/common/xes"
	"GIM/pkg/common/xmysql"
	"GIM/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
	xes.NewElasticsearchClient(conf.Elasticsearch)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}

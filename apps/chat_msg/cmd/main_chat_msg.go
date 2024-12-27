package main

import (
	"GIM/apps/chat_msg/dig"
	"GIM/apps/chat_msg/internal/config"
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

package main

import (
	"GIM/apps/apis/upload/internal/config"
	"GIM/apps/apis/upload/internal/server"
	"GIM/pkg/commands"
	"GIM/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	commands.Run(server.NewServer())
}

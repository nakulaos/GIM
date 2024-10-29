package main

import (
	"GIM/apps/interfaces/internal/config"
	_ "GIM/apps/interfaces/internal/config"
	"GIM/apps/interfaces/internal/server"
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

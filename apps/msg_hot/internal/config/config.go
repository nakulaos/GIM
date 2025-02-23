package config

import (
	"GIM/pkg/common/xlog"
	"GIM/pkg/common/xsnowflake"
	"GIM/pkg/conf"
	"GIM/pkg/utils"
	"flag"
	"github.com/spf13/cast"
)

type Config struct {
	Name        string              `yaml:"name"`
	ServerID    int                 `yaml:"server_id"`
	Log         string              `yaml:"log"`
	Etcd        *conf.Etcd          `yaml:"etcd"`
	Mongo       *conf.Mongo         `yaml:"mongo"`
	Redis       *conf.Redis         `yaml:"redis"`
	MsgConsumer *conf.KafkaConsumer `yaml:"msg_consumer"`
}

var (
	config = new(Config)
)

var (
	confFile = flag.String("cfg", "./configs/msg_hot.yaml", "config file")
	serverId = flag.Int("sid", 1, "server id")
)

func init() {
	flag.Parse()
	utils.YamlToStruct(*confFile, config)

	config.ServerID = *serverId

	xsnowflake.NewSnowflake(config.ServerID)
	xlog.Shared(config.Log, config.Name+cast.ToString(config.ServerID))
}

func NewConfig() *Config {
	return config
}

func GetConfig() *Config {
	return config
}

package config

import (
	"flag"
	"github.com/spf13/cast"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name       string        `yaml:"name"`
	ServerID   int           `yaml:"server_id"`
	Log        string        `yaml:"log"`
	GrpcServer *conf.Grpc    `yaml:"grpc_server"`
	Etcd       *conf.Etcd    `yaml:"etcd"`
	Mysql      *conf.Mysql   `yaml:"mysql"`
	Redis      *conf.Redis   `yaml:"redis"`
	Monitor    *conf.Monitor `yaml:"monitor"`
	Mongo      *conf.Mongo   `yaml:"mongo"`
}

var (
	config = new(Config)
)

var (
	confFile = flag.String("cfg", "./configs/lbs.yaml", "config file")
	serverId = flag.Int("sid", 1, "server id")
	grpcPort = flag.Int("gp", 8000, "grpc server port")
)

func init() {
	flag.Parse()
	utils.YamlToStruct(*confFile, config)

	config.ServerID = *serverId
	config.GrpcServer.ServerID = config.ServerID
	config.GrpcServer.Port = *grpcPort

	xlog.Shared(config.Log, config.Name+cast.ToString(config.ServerID))
}

func NewConfig() *Config {
	return config
}

func GetConfig() *Config {
	return config
}

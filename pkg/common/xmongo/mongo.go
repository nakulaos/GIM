package xmongo

import (
	"GIM/pkg/common/xlog"
	"GIM/pkg/conf"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

var (
	cli *MongoClient
)

type MongoClient struct {
	cfg *conf.Mongo
	db  *mongo.Database
}

func NewMongoClient(cfg *conf.Mongo) *MongoClient {
	cli = &MongoClient{cfg: cfg}
	cli.db, _ = connectDB(cfg)
	return cli
}

func GetDB() *mongo.Database {
	if cli.db == nil {
		cli.db, _ = connectDB(cli.cfg)
	}
	return cli.db
}

func connectDB(cfg *conf.Mongo) (db *mongo.Database, err error) {
	var (
		uri           string
		clientOptions *options.ClientOptions
		client        *mongo.Client
	)

	uri = fmt.Sprintf("mongodb://%s/?maxPoolSize=%d", strings.Join(cfg.Hosts, ","), cfg.MaxPoolSize)
	// Set client options
	clientOptions = options.Client().ApplyURI(uri).SetAuth(
		options.Credential{
			AuthMechanism: "SCRAM-SHA-256", // [ 'SCRAM-SHA-1', 'SCRAM-SHA-256' ]
			Username:      cfg.Username,
			Password:      cfg.Password,
			AuthSource:    cfg.AuthSource,
		}).SetConnectTimeout(time.Duration(cfg.Timeout) * time.Millisecond)

	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	db = client.Database(cfg.Db)
	return
}

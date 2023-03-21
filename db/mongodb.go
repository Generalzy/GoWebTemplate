package db

import (
	"context"
	"github.com/Generalzy/GeneralSaaS/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(c *conf.MongoDBConf) (*mongo.Client, error) {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(c.Uri)

	// 连接到MongoDB
	// 与mongo.NewClient()等价
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

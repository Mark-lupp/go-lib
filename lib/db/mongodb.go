package db

import (
	"context"
	"time"

	"github.com/Mark-lupp/go-lib/lib/config"
	"github.com/Mark-lupp/go-lib/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoDb() {
	//链接mongo服务
	opt := options.Client().ApplyURI(config.GetMgoConfig().GetUrl())
	opt.SetLocalThreshold(3 * time.Second) //只使用与mongo操作耗时小于3秒的
	if config.GetMgoConfig().GetName() != "" && config.GetMgoConfig().GetPass() != "" {
		opt.SetAuth(options.Credential{
			Username: config.GetMgoConfig().GetName(),
			Password: config.GetMgoConfig().GetPass(),
		})
	}
	opt.SetMaxConnIdleTime(5 * time.Second) //指定连接可以保持空闲的最大毫秒数
	opt.SetMaxPoolSize(200)                 //使用最大的连接数
	if mgo, err = mongo.Connect(context.TODO(), opt); err != nil {
		panic(err)
	}
	log.NewLogger().Debug("mgo success : " + config.GetMgoConfig().GetUrl())

}

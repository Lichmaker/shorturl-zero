package svc

import (
	"time"

	"golang.org/x/sync/singleflight"

	"github.com/Shopify/sarama"
	"github.com/lichmaker/short-url-micro/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config               config.Config
	GormDB               *gorm.DB
	Redis                *redis.Redis
	ShortenSg            *singleflight.Group
	KafkaProducerMsgChan chan *sarama.ProducerMessage
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysqldriver.New(mysqldriver.Config{
		DSN: c.Mysql.DataSource,
	}))
	if err != nil {
		panic(err)
	}
	dbCore, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbCore.SetConnMaxLifetime(time.Duration(c.Mysql.MaxLifeTime))
	dbCore.SetMaxIdleConns(int(c.Mysql.MaxIdleConn))
	dbCore.SetMaxOpenConns(int(c.Mysql.MaxOpenConn))

	return &ServiceContext{
		Config: c,
		GormDB: db,
		Redis: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
		}),
		ShortenSg: &singleflight.Group{},
	}
}

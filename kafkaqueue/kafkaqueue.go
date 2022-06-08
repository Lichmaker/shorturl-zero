package main

import (
	"context"
	"flag"
	"time"

	"github.com/Shopify/sarama"
	"github.com/lichmaker/short-url-micro/kafkaqueue/internal/config"
	"github.com/lichmaker/short-url-micro/kafkaqueue/internal/server"
	"github.com/lichmaker/short-url-micro/pkg/gormlogger"
	"github.com/lichmaker/short-url-micro/pkg/kafkahelper"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var configFile = flag.String("f", "etc/config.yml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// log、prometheus、trace、metricsUrl.
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	consumerGroup, err := sarama.NewConsumerGroup(c.KafkaConfig.Addr, c.KafkaConfig.ConsumerGroup, kafkahelper.GetSaramaConfig(c.KafkaConfig))
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysqldriver.New(mysqldriver.Config{
		DSN: c.Mysql.DataSource,
	}), &gorm.Config{
		Logger: gormlogger.NewGormLogger(),
	})
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

	ctx, cancel := context.WithCancel(context.Background())
	mySer := &server.MyService{
		ConsumerG: consumerGroup,
		Ctx:       ctx,
		CtxCancel: cancel,
		GormDB:    db,
		Config:    c.KafkaConfig,
	}

	serviceGroup.Add(mySer)
	serviceGroup.Start()
}

package config

import (
	"github.com/lichmaker/short-url-micro/pkg/kafkahelper"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf

	KafkaConfig kafkahelper.Config
	Mysql       struct {
		DataSource  string
		MaxLifeTime int64
		MaxIdleConn int64
		MaxOpenConn int64
	}
}

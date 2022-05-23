package config

import (
	"github.com/lichmaker/short-url-micro/pkg/kafkahelper"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource  string
		MaxLifeTime int64
		MaxIdleConn int64
		MaxOpenConn int64
	}
	Jwt struct {
		Secret string
	}
	BloomRedisKey string
	KafkaConfig   kafkahelper.Config
}

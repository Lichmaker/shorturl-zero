package svc

import (
	"github.com/lichmaker/short-url-micro/api/internal/config"
	"github.com/lichmaker/short-url-micro/rpc/shorturl"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	ShortRpc shorturl.Shorturl
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		ShortRpc: shorturl.NewShorturl(zrpc.MustNewClient(c.ShortRpc)),
	}
}

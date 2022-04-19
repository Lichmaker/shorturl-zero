package main

import (
	"flag"
	"fmt"

	"github.com/lichmaker/short-url-micro/pkg/intercepter"
	"github.com/lichmaker/short-url-micro/rpc/internal/config"
	"github.com/lichmaker/short-url-micro/rpc/internal/server"
	"github.com/lichmaker/short-url-micro/rpc/internal/svc"
	short_url_micro "github.com/lichmaker/short-url-micro/rpc/type/short-url-micro"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/shorturl.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewShorturlServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		short_url_micro.RegisterShorturlServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// 加入自定义的拦截器
	s.AddUnaryInterceptors(intercepter.MyRpcIntercepter)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

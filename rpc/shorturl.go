package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/lichmaker/short-url-micro/pkg/intercepter"
	"github.com/lichmaker/short-url-micro/pkg/kafkahelper"
	"github.com/lichmaker/short-url-micro/pkg/kafkaproducer"
	"github.com/lichmaker/short-url-micro/rpc/internal/config"
	"github.com/lichmaker/short-url-micro/rpc/internal/server"
	"github.com/lichmaker/short-url-micro/rpc/internal/svc"
	short_url_micro "github.com/lichmaker/short-url-micro/rpc/type/short-url-micro"
	"go.uber.org/zap"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/logx/zapx"

	// "github.com/zeromicro/zero-contrib/logx/zapx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/shorturl.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	// 注入zap
	writer, err := zapx.NewZapWriter(
		zap.AddStacktrace(zap.ErrorLevel), // 调整error才显示trace
	)
	logx.Must(err)
	logx.SetWriter(writer)

	// kafka生产者
	kafkaCtx, cancelFn := context.WithCancel(context.Background())
	kafkaProducer := kafkaproducer.MustNew(kafkaCtx, cancelFn, c.KafkaConfig.Addr, kafkahelper.GetSaramaConfig(c.KafkaConfig))
	ctx.KafkaProducerMsgChan = kafkaProducer.Ch // 把生产者的消息channel注册到svcCtx中

	srv := server.NewShorturlServer(ctx)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		short_url_micro.RegisterShorturlServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// 加入自定义的拦截器
	s.AddUnaryInterceptors(intercepter.MyRpcIntercepter)

	// 建立一个 serviceGroup，把 rpc 服务注册进去
	serviceGroup := service.NewServiceGroup()
	serviceGroup.Add(s)

	// kafka生产者注册到group，等待启动
	serviceGroup.Add(kafkaProducer)

	defer serviceGroup.Stop()
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	serviceGroup.Start()

}

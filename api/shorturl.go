package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/lichmaker/short-url-micro/api/internal/config"
	"github.com/lichmaker/short-url-micro/api/internal/handler"
	"github.com/lichmaker/short-url-micro/api/internal/svc"
	"github.com/lichmaker/short-url-micro/pkg/apiresponse"
	"github.com/lichmaker/short-url-micro/pkg/errx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/shorturl.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiresponse.Do(r.Context(), w, nil, errx.NewWithCode(errx.CODE_NOT_FOUND_HANDLER))
	})), rest.WithNotAllowedHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiresponse.Do(r.Context(), w, nil, errx.New(errx.CODE_UNAUTHORIZED, "无权操作"))
	})), rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		apiresponse.Do(r.Context(), w, nil, errx.New(errx.CODE_UNAUTHORIZED, "认证错误"))
	}))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

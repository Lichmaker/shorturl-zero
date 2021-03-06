// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/lichmaker/short-url-micro/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/appKey/register",
				Handler: registerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/appKey/login",
				Handler: loginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:short",
				Handler: JumpHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/shorten",
				Handler: shortenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}

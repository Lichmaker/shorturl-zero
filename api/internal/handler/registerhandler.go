package handler

import (
	"net/http"

	"github.com/lichmaker/short-url-micro/api/internal/logic"
	"github.com/lichmaker/short-url-micro/api/internal/svc"
	"github.com/lichmaker/short-url-micro/api/internal/types"
	"github.com/lichmaker/short-url-micro/pkg/apiresponse"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func registerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(req)
		apiresponse.Do(r.Context(), w, resp, err)
	}
}

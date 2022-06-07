package handler

import (
	"net/http"

	"github.com/lichmaker/short-url-micro/api/internal/logic"
	"github.com/lichmaker/short-url-micro/api/internal/svc"
	"github.com/lichmaker/short-url-micro/api/internal/types"
	"github.com/lichmaker/short-url-micro/pkg/apiresponse"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			apiresponse.Do(r.Context(), w, nil, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(req)
		apiresponse.Do(r.Context(), w, resp, err)
	}
}

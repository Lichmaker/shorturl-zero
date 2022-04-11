package handler

import (
	"net/http"

	"github.com/lichmaker/short-url-micro/api/internal/logic"
	"github.com/lichmaker/short-url-micro/api/internal/svc"
	"github.com/lichmaker/short-url-micro/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func JumpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JumpRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewJumpLogic(r.Context(), svcCtx)
		longUrl, err := l.Jump(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			// fmt.Println(longUrl)
			w.Header().Set("Location", *longUrl)
			w.WriteHeader(http.StatusFound)
		}
	}
}

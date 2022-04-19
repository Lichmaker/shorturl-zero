package apiresponse

import (
	"context"
	"net/http"

	"github.com/lichmaker/short-url-micro/pkg/errx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

type Body struct {
	Code int64       `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Do(ctx context.Context, w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {

		s, ok := status.FromError(err)
		if ok {
			logx.WithContext(ctx).Errorf("RPC返回错误: code %d, msg %s", s.Code(), s.Message())
			body.Code = int64(s.Code())
			body.Msg = s.Message()
		} else {
			switch t := err.(type) {
			case *errx.CustomError:
				logx.WithContext(ctx).Errorf("业务返回错误: code %d, msg %s", t.Code, t.Msg)
				body.Code = t.Code
				body.Msg = t.Msg
			default:
				logx.WithContext(ctx).Errorf("错误捕捉：%s", err.Error())
				body.Code = errx.CODE_UNDEFINED
				body.Msg = err.Error()
			}
		}
	} else {
		// 正常输出内容
		body.Code = errx.CODE_SUCCESS
		body.Data = resp
		body.Msg = "success"
	}
	httpx.WriteJson(w, http.StatusOK, body)
}

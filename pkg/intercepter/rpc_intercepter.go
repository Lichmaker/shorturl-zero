package intercepter

import (
	"context"

	"github.com/lichmaker/short-url-micro/pkg/errx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MyRpcIntercepter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	res, err := handler(ctx, req)

	if err != nil {
		causeErr := errors.Cause(err)
		switch t := causeErr.(type) {
		case *errx.CustomError:
			// 记录日志，然后返回err
			// todo 日志记录 t.Msg

			err = status.Error(codes.Code(t.Code), t.Error())
		case error:
			// todo 日志记录
			logx.Errorf("内部错误 %s", t.Error())
			err = status.Error(codes.Code(errx.CODE_UNDEFINED), "服务器内部错误")
		default:
			err = status.Error(codes.Code(errx.CODE_UNDEFINED), "未知错误")
		}
	}

	return res, err
}

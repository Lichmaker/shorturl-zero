package logic

import (
	"context"

	"github.com/lichmaker/short-url-micro/api/internal/svc"
	"github.com/lichmaker/short-url-micro/api/internal/types"
	short_url_micro "github.com/lichmaker/short-url-micro/rpc/type/short-url-micro"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginRequest) (resp *types.LoginResponse, err error) {
	rpcResponse, err := l.svcCtx.ShortRpc.Verify(l.ctx, &short_url_micro.VerifyRequest{
		AppId:     req.AppId,
		AppSecret: req.AppSecret,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Token:     rpcResponse.Token,
		ExpiredAt: rpcResponse.ExpireAt,
	}, nil
}

package logic

import (
	"context"
	"errors"

	"github.com/lichmaker/short-url-micro/api/internal/svc"
	"github.com/lichmaker/short-url-micro/api/internal/types"
	short_url_micro "github.com/lichmaker/short-url-micro/rpc/type/short-url-micro"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	if req.AdminSecret != l.svcCtx.Config.AdminSecret {
		return nil, errors.New("密码错误")
	}
	rpcResponse, err := l.svcCtx.ShortRpc.Register(l.ctx, &short_url_micro.RegisterRquest{
		AppId: req.AppId,
		Name:  req.Name,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.RegisterResponse{}
	resp.AppId = rpcResponse.AppId
	resp.AppSecret = rpcResponse.AppSecret

	return
}

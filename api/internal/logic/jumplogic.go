package logic

import (
	"context"

	"github.com/lichmaker/short-url-micro/api/internal/svc"
	"github.com/lichmaker/short-url-micro/api/internal/types"
	short_url_micro "github.com/lichmaker/short-url-micro/rpc/type/short-url-micro"

	"github.com/zeromicro/go-zero/core/logx"
)

type JumpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJumpLogic(ctx context.Context, svcCtx *svc.ServiceContext) JumpLogic {
	return JumpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JumpLogic) Jump(req types.JumpRequest) (*string, error) {
	rpcResponse, err := l.svcCtx.ShortRpc.Get(l.ctx, &short_url_micro.GetRequest{
		Short: req.Short,
	})
	if err != nil {
		return nil, err
	}

	return &rpcResponse.Long, nil
}

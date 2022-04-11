package logic

import (
	"context"

	"github.com/lichmaker/short-url-micro/model/shorts"
	"github.com/lichmaker/short-url-micro/rpc/internal/svc"
	short_url_micro "github.com/lichmaker/short-url-micro/rpc/type/short-url-micro"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *short_url_micro.GetRequest) (*short_url_micro.GetResponse, error) {
	model, err := shorts.GetByShort(l.svcCtx.GormDB, in.Short)
	if err != nil {
		return nil, err
	}
	if model.Id == 0 {
		return nil, status.Error(codes.NotFound, "不存在的数据")
	}

	return &short_url_micro.GetResponse{
		Long: model.Long,
	}, nil
}

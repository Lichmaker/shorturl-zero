package logic

import (
	"context"

	"github.com/lichmaker/short-url-micro/pkg/errx"
	"github.com/lichmaker/short-url-micro/pkg/shorten"
	"github.com/lichmaker/short-url-micro/rpc/internal/svc"
	short_url_micro "github.com/lichmaker/short-url-micro/rpc/type/short-url-micro"
	"github.com/pkg/errors"

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
	shortenInstance := &shorten.Shorten{
		Ctx:      l.ctx,
		GormDB:   l.svcCtx.GormDB,
		Redis:    l.svcCtx.Redis,
		BloomKey: l.svcCtx.Config.BloomRedisKey,
		Sg:       l.svcCtx.ShortenSg,
	}

	model, err := shortenInstance.Get(in.Short)
	if err != nil {
		return nil, err
	}
	if model.Id == 0 {
		return nil, errors.Wrapf(errx.NewWithCode(errx.CODE_DATA_NOT_FOUND), "查询short model不存在。 short:%s", in.Short)
	}

	return &short_url_micro.GetResponse{
		Long: model.Long,
	}, nil
}

package logic

import (
	"context"

	"github.com/lichmaker/short-url-micro/pkg/shorten"
	"github.com/lichmaker/short-url-micro/rpc/internal/svc"
	short_url_micro "github.com/lichmaker/short-url-micro/rpc/type/short-url-micro"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *short_url_micro.ShortenRequest) (*short_url_micro.ShortenResponse, error) {

	shortenInstance := &shorten.Shorten{
		Ctx:      l.ctx,
		GormDB:   l.svcCtx.GormDB,
		Redis:    l.svcCtx.Redis,
		BloomKey: l.svcCtx.Config.BloomRedisKey,
		Sg:       l.svcCtx.ShortenSg,
	}

	model, err := shortenInstance.Make(in.Long)
	if err != nil {
		return nil, err
	}

	// host := helpers.FillHttpScheme(l.svcCtx.Config.AppHost)
	// host = strings.TrimRight(host, "/")
	// shortUrl := host + "/" + model.Short

	return &short_url_micro.ShortenResponse{
		Short: model.Short,
	}, nil
}

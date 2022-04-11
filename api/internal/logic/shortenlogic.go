package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/lichmaker/short-url-micro/api/internal/svc"
	"github.com/lichmaker/short-url-micro/api/internal/types"
	"github.com/lichmaker/short-url-micro/pkg/helpers"
	short_url_micro "github.com/lichmaker/short-url-micro/rpc/type/short-url-micro"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShortenLogic {
	return ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req types.ShortenRequest) (resp *types.ShortenResponse, err error) {
	rpcResponse, err := l.svcCtx.ShortRpc.Shorten(l.ctx, &short_url_micro.ShortenRequest{
		Long: req.Long,
	})
	if err != nil {
		return nil, err
	}

	host := fmt.Sprintf("%s:%d", l.svcCtx.Config.AppHost, l.svcCtx.Config.Port)
	host = helpers.FillHttpScheme(host)
	host = strings.TrimRight(host, "/")
	shortUrl := host + "/" + rpcResponse.Short

	return &types.ShortenResponse{
		Short: shortUrl,
	}, nil
}

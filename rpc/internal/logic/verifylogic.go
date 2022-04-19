package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/lichmaker/short-url-micro/model/apps"
	"github.com/lichmaker/short-url-micro/pkg/errx"
	"github.com/lichmaker/short-url-micro/pkg/jwt"
	"github.com/lichmaker/short-url-micro/rpc/internal/svc"
	short_url_micro "github.com/lichmaker/short-url-micro/rpc/type/short-url-micro"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyLogic) Verify(in *short_url_micro.VerifyRequest) (*short_url_micro.VerifyResponse, error) {
	appModel, err := apps.FindByAppId(l.svcCtx.GormDB, in.AppId)
	if err != nil {
		return nil, err
	}
	if appModel.Id == 0 {
		return nil, errors.Wrapf(errx.NewWithCode(errx.CODE_DATA_NOT_FOUND), "不存在该APPID:%s", in.AppId)
	}
	if strings.Compare(appModel.AppSecret, in.AppSecret) != 0 {
		return nil, errors.Wrapf(errx.New(errx.CODE_UNAUTHORIZED, "密码错误"), "密码错误! appid:%s, appsecret:%s", in.AppId, in.AppSecret)
	}

	myJwt := &jwt.MyJwt{Secret: l.svcCtx.Config.Jwt.Secret}
	genJwt, err := myJwt.Generate(in.AppId)
	if err != nil {
		return nil, err
	}

	return &short_url_micro.VerifyResponse{
		Token:    genJwt,
		ExpireAt: fmt.Sprint(myJwt.ExpiresAt),
	}, nil
}

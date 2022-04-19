package logic

import (
	"context"

	"github.com/lichmaker/short-url-micro/model/apps"
	"github.com/lichmaker/short-url-micro/pkg/errx"
	"github.com/lichmaker/short-url-micro/pkg/helpers"
	"github.com/lichmaker/short-url-micro/rpc/internal/svc"
	short_url_micro "github.com/lichmaker/short-url-micro/rpc/type/short-url-micro"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *short_url_micro.RegisterRquest) (*short_url_micro.RegisterResponse, error) {
	appIdModel, err := apps.FindByAppId(l.svcCtx.GormDB, in.AppId)
	if err != nil {
		return nil, err
	}
	if appIdModel.Id > 0 {
		return nil, errors.Wrapf(errx.NewWithCode(errx.CODE_DATA_EXISTS), "注册失败，appid已存在:%s", in.AppId)
	}

	model := &apps.Apps{
		AppId:     in.AppId,
		Name:      in.Name,
		AppSecret: helpers.RandomStr(32),
	}
	l.svcCtx.GormDB.Create(&model)

	return &short_url_micro.RegisterResponse{
		AppId:     model.AppId,
		AppSecret: model.AppSecret,
	}, nil
}

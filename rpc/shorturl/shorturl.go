// Code generated by goctl. DO NOT EDIT!
// Source: shorturl.proto

package shorturl

import (
	"context"

	"github.com/lichmaker/short-url-micro/rpc/type/short-url-micro"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetRequest       = short_url_micro.GetRequest
	GetResponse      = short_url_micro.GetResponse
	RegisterResponse = short_url_micro.RegisterResponse
	RegisterRquest   = short_url_micro.RegisterRquest
	ShortenRequest   = short_url_micro.ShortenRequest
	ShortenResponse  = short_url_micro.ShortenResponse
	VerifyRequest    = short_url_micro.VerifyRequest
	VerifyResponse   = short_url_micro.VerifyResponse

	Shorturl interface {
		Register(ctx context.Context, in *RegisterRquest, opts ...grpc.CallOption) (*RegisterResponse, error)
		Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
		Shorten(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShortenResponse, error)
		Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	}

	defaultShorturl struct {
		cli zrpc.Client
	}
)

func NewShorturl(cli zrpc.Client) Shorturl {
	return &defaultShorturl{
		cli: cli,
	}
}

func (m *defaultShorturl) Register(ctx context.Context, in *RegisterRquest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := short_url_micro.NewShorturlClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultShorturl) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	client := short_url_micro.NewShorturlClient(m.cli.Conn())
	return client.Verify(ctx, in, opts...)
}

func (m *defaultShorturl) Shorten(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShortenResponse, error) {
	client := short_url_micro.NewShorturlClient(m.cli.Conn())
	return client.Shorten(ctx, in, opts...)
}

func (m *defaultShorturl) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	client := short_url_micro.NewShorturlClient(m.cli.Conn())
	return client.Get(ctx, in, opts...)
}

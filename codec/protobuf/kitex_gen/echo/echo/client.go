// Code generated by Kitex v0.2.1. DO NOT EDIT.

package echo

import (
	"context"
	"github.com/bbbearxyz/kitex-benchmark/codec/protobuf/kitex_gen/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/cloudwego/kitex/transport"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Send(ctx context.Context, Req *echo.Request, callOptions ...callopt.Option) (r *echo.Response, err error)
	StreamTest(ctx context.Context, callOptions ...callopt.Option) (stream Echo_StreamTestClient, err error)
	TCPCostTest(ctx context.Context, callOptions ...callopt.Option) (stream Echo_TCPCostTestClient, err error)
}

type Echo_StreamTestClient interface {
	streaming.Stream
	Send(*echo.Request) error
	Recv() (*echo.Response, error)
}

type Echo_TCPCostTestClient interface {
	streaming.Stream
	Send(*echo.Request) error
	Recv() (*echo.Response, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, client.WithTransportProtocol(transport.GRPC))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kEchoClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kEchoClient struct {
	*kClient
}

func (p *kEchoClient) Send(ctx context.Context, Req *echo.Request, callOptions ...callopt.Option) (r *echo.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Send(ctx, Req)
}

func (p *kEchoClient) StreamTest(ctx context.Context, callOptions ...callopt.Option) (stream Echo_StreamTestClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StreamTest(ctx)
}

func (p *kEchoClient) TCPCostTest(ctx context.Context, callOptions ...callopt.Option) (stream Echo_TCPCostTestClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TCPCostTest(ctx)
}

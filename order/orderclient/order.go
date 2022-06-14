// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package orderclient

import (
	"context"

	"saga/order/order"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateRequest  = order.CreateRequest
	CreateResponse = order.CreateResponse

	Order interface {
		Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
		CreateRevert(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultOrder) CreateRevert(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.CreateRevert(ctx, in, opts...)
}
package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"saga/api/internal/config"
	"saga/order/orderclient"
	"saga/product/productclient"
)

type ServiceContext struct {
	Config     config.Config
	OrderRpc   orderclient.Order
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}

package logic

import (
	"context"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"saga/api/internal/svc"
	"saga/api/internal/types"
	"saga/order/order"
	"saga/product/product"
)

type OrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderLogic {
	return OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderLogic) Order(req types.CreateOrderRequest) (resp *types.CreateOrderResponse, err error) {

	orderRpcBusiServer, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}

	productRpcBusiServer, err := l.svcCtx.Config.ProductRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}

	// dtm 服务的 etcd 注册地址
	var dtmServer = "etcd://localhost:2379/dtmservice"

	// 创建一个gid
	gid := dtmgrpc.MustGenGid(dtmServer)

	// 创建一个saga协议的事务
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(
			orderRpcBusiServer+"/orderclient.Order/Create",
			orderRpcBusiServer+"/orderclient.Order/CreateRevert",
			&order.CreateRequest{
				Uid:    req.Uid,
				Pid:    req.Pid,
				Amount: req.Amount,
				Status: 0,
			}).
		Add(
			productRpcBusiServer+"/productclient.Product/DecrStock",
			productRpcBusiServer+"/productclient.Product/DecrStockRevert",
			&product.DecrStockRequest{
				Id:  req.Pid,
				Num: 1,
			})

	// 事务提交
	err = saga.Submit()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateOrderResponse{}, nil
}

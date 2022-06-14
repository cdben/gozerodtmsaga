package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"saga/model"
	"saga/order/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(mysql, c.CacheRedis),
	}
}

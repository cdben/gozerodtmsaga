package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"saga/model"
	"saga/product/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(mysql, c.CacheRedis),
	}
}

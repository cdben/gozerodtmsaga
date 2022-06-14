GOPATH:=$(shell go env GOPATH)

.PHONY: rpc
rpc:
	@goctl rpc proto -src ./order/order.proto -dir ./order && goctl rpc proto -src ./product/product.proto -dir ./product

.PHONY: api
api:
	@goctl api go -api ./api/order.api -dir ./api


.PHONY: model
model:
	@goctl model mysql ddl -src ./scripts/sql/product.sql -dir ./model -c && \
	goctl model mysql ddl -src ./scripts/sql/order.sql -dir ./model -c


.PHONY: order_api
order_api:
	@go run api/order.go

.PHONY: product
product:
	@go run product/product.go

.PHONY: order
order:
	@go run order/order.go
package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/qtask/examples/flowserver/services/order"
	"github.com/micro-plat/qtask/qtask"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	app.Initializing(func(c component.IContainer) error {
		return nil
	})

	app.Micro("/order/request", order.NewRequestHandler)
	app.Flow("/order/pay", order.NewPayHandler)

	qtask.Bind(app, 10, 3)
}

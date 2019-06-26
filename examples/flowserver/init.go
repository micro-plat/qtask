package main

import (
	"sync"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/qtask"
	"github.com/micro-plat/qtask/examples/flowserver/services/order"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *flowserver) init() {
	var one sync.Once
	r.Initializing(func(c component.IContainer) error {
		one.Do(func() {
			if err := qtask.CreateDB(c); err != nil {
				c.GetLogger().Error(err)
			}
		})
		return nil
	})

	r.Micro("/order/request", order.NewRequestHandler, "*")
	r.Flow("/order/pay", order.NewPayHandler, "*")
	qtask.Bind(r.MicroApp, 10, 3)
}

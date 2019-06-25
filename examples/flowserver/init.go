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

		//appconf.func#//
		//#appconf.func//

		//db.init#//
		//#db.init//

		//cache.init#//
		//#cache.init//

		//queue.init#//
		//#queue.init//

		//login.router#//
		//#login.router//

		//service.router#//
		r.Micro("/order/request", order.NewRequestHandler, "*")
		r.Flow("/order/pay", order.NewPayHandler, "*")

		r.CRON("/task/scan", qtask.Scan)      //定时扫描任务
		r.CRON("/task/clear", qtask.Clear(7)) //定时清理任务，删除７天的任务数据

		//#service.router//

		return nil
	})
}

package qtask

import (
	"fmt"
	"sync"

	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/hydra"
	"github.com/micro-plat/qtask/services"
)

//Bind 绑定服务
//注册 /task/scan 为cron
//注册 /task/clear 为cron
var doOnce sync.Once

func Bind(app *hydra.MicroApp, scanSecond int, dayBefore int) {
	if scanSecond >= 60 {
		panic(fmt.Sprintf("qtask.bind　扫描时间取值为0-59,当前值:%d", scanSecond))
	}
	if scanSecond <= 0 {
		scanSecond = 0
	}
	ch := app.GetDynamicCron()

	doOnce.Do(func() {
		go func() {
			ch <- &conf.Task{Cron: fmt.Sprintf("@every %ds", scanSecond), Service: "/task/scan"}
			ch <- &conf.Task{Cron: "@hourly", Service: "/task/clear"}
		}()
	})
	app.CRON("/task/scan", services.Scan)              //定时扫描任务
	app.CRON("/task/clear", services.Clear(dayBefore)) //定时清理任务，删除７天的任务数
}

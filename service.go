package qtask

import (
	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/hydra"
)

//Bind 绑定服务
//注册 /task/scan 为cron
//注册 /task/clear 为cron
func Bind(app *hydra.MicroApp, dayBefore int) {
	app.CRON("/task/scan", Scan)              //定时扫描任务
	app.CRON("/task/clear", Clear(dayBefore)) //定时清理任务，删除７天的任务数
	ch := app.GetDynamicCron()
	ch <- &conf.Task{Cron: "@every 1s", Service: "/task/scan"}
	ch <- &conf.Task{Cron: "@hourly", Service: "/task/clear"}
}

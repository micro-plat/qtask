package qtask

import (
	"github.com/micro-plat/hydra/hydra"
)

//Bind 绑定服务
//注册 /task/scan 为cron
//注册 /task/clear 为cron
func Bind(app *hydra.MicroApp, day int) {
	app.CRON("/task/scan", Scan)        //定时扫描任务
	app.CRON("/task/clear", Clear(day)) //定时清理任务，删除７天的任务数据
}

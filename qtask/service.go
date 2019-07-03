package qtask

import (
	"fmt"

	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra"
)

//Bind 绑定服务
//注册 /task/scan 为cron
//注册 /task/clear 为cron
func Bind(app *hydra.MicroApp, scanSecond int, dayBefore int) {
	if scanSecond >= 60 {
		panic(fmt.Sprintf("qtask.bind　扫描时间取值为0-59,当前值:%d", scanSecond))
	}
	if scanSecond <= 0 {
		scanSecond = 0
	}
	app.CRON("/task/scan", Scan)              //定时扫描任务
	app.CRON("/task/clear", Clear(dayBefore)) //定时清理任务，删除７天的任务数
	ch := app.GetDynamicCron()
	ch <- &conf.Task{Cron: fmt.Sprintf("@every %ds", scanSecond), Service: "/task/scan"}
	ch <- &conf.Task{Cron: "@hourly", Service: "/task/clear"}
}

//Scan 扫描任务，定时从ＤＢ中扫描待处理任务并放入消息队列
func Scan(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("---------------qtask:任务扫描----------------")
	db, err := ctx.GetContainer().GetDB(dbName)
	if err != nil {
		return err
	}
	rows, err := queryTasks(db)
	if err != nil {
		return err
	}
	ctx.Log.Info("发送任务到消息队列")
	queue, err := ctx.GetContainer().GetQueue(queueName)
	if err != nil {
		return err
	}
	for _, row := range rows {
		qName := row.GetString("queue_name")
		content := row.GetString("content")
		if err := queue.Push(qName, content); err != nil {
			return fmt.Errorf("发送消息(%s)到消息队列(%s)失败:%v", content, qName, err)
		}
	}
	ctx.Log.Infof("处理消息:%d条", rows.Len())
	return "success"
}

//Clear 清理任务，删除指定时间以前的任务
func Clear(day int) func(ctx *context.Context) (r interface{}) {
	return func(ctx *context.Context) (r interface{}) {
		ctx.Log.Infof("---------------qtask:清理%d天前的任务----------------", day)
		db, err := ctx.GetContainer().GetDB(dbName)
		if err != nil {
			return err
		}
		err = clearTask(db, day)
		if err != nil {
			return err
		}
		return "success"
	}
}

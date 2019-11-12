package services

import (
	"fmt"

	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/qtask/modules/db"
	"github.com/micro-plat/qtask/qtask"
)

//Scan 扫描任务，定时从ＤＢ中扫描待处理任务并放入消息队列
func Scan(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("---------------qtask:任务扫描----------------")
	xdb, err := ctx.GetContainer().GetDB(qtask.DBName)
	if err != nil {
		return err
	}
	rows, err := db.QueryTasks(xdb)
	if err != nil {
		return err
	}
	ctx.Log.Info("发送任务到消息队列")
	queue, err := ctx.GetContainer().GetQueue(qtask.QueueName)
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
		xdb, err := ctx.GetContainer().GetDB(qtask.DBName)
		if err != nil {
			return err
		}
		err = db.ClearTask(xdb, day)
		if err != nil {
			return err
		}
		return "success"
	}
}

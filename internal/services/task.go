package services

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/qtask/internal/modules/const/conf"
	"github.com/micro-plat/qtask/internal/modules/db"
)

//Scan 扫描任务，定时从ＤＢ中扫描待处理任务并放入消息队列
func Scan(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("---------------qtask:任务扫描----------------")
	xdb, err := hydra.C.DB().GetDB(conf.DBName)
	if err != nil {
		return err
	}
	rows, err := db.QueryTasks(xdb)
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		return "empty"
	}

	ctx.Log().Debug("发送任务到消息队列")
	queue, err := hydra.C.Queue().GetQueue(conf.QueueName)
	if err != nil {
		return err
	}
	for _, row := range rows {
		qName := row.GetString("queue_name")
		content := row.GetString("content")
		if err := queue.Send(qName, content); err != nil {
			return fmt.Errorf("发送消息(%s)到消息队列(%s)失败:%v", content, qName, err)
		}
	}
	ctx.Log().Debugf("处理消息:%d条", rows.Len())
	return "success"
}

//Clear 清理任务，删除指定时间以前的任务
func Clear() func(ctx hydra.IContext) (r interface{}) {
	return func(ctx hydra.IContext) (r interface{}) {
		ctx.Log().Debugf("---------------qtask:清理任务----------------")
		xdb, err := hydra.C.DB().GetDB(conf.DBName)
		if err != nil {
			return err
		}
		ctx.Log().Debug("1.开始清除任务")
		if err = db.ClearTask(xdb); err != nil {
			return err
		}
		ctx.Log().Debug("2.完成清除任务")
		return "success"
	}
}

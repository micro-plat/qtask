package qtask

import (
	"fmt"

	"github.com/micro-plat/hydra/context"
)

//Scan 扫描任务，定时从ＤＢ中扫描待处理任务并放入消息队列
func Scan(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("---------------qtask:任务扫描----------------")
	rows, err := queryTasks(ctx.GetContainer().GetRegularDB(dbName))
	if err != nil {
		return err
	}
	ctx.Log.Info("发送任务到消息队列")
	queue := ctx.GetContainer().GetRegularQueue(queueName)
	for _, row := range rows {
		qName := row.GetString("name")
		content := row.GetString("content")
		if err := queue.Push(qName, content); err != nil {
			return fmt.Errorf("发送消息(%s)到消息队列(%s)失败:%v", content, qName, err)
		}
	}
	ctx.Log.Infof("处理消息:%d条", rows.Len())
	return "success"
}

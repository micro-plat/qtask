package qtask

import (
	"strings"

	"github.com/micro-plat/hydra/conf"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/hydra"
)

//ConsumeCallBack ConsumeCallBack
type ConsumeCallBack = func(ctx *context.Context) error

//RegConsume 注册消费程序业务
func RegConsume(app *hydra.MicroApp, queueName string, callback ConsumeCallBack, tags ...string) {
	path := strings.Replace(queueName, ":", "_", -1)
	queryChan := app.GetDynamicQueue()
	queryChan <- &conf.Queue{
		Queue:   queueName,
		Service: path,
	}

	app.MQC(path, func(ctx *context.Context) interface{} {
		taskID := ctx.Request.GetInt64("task_id")
		err := Processing(ctx, taskID)
		if err != nil {
			return err
		}
		err = callback(ctx)
		if err != nil {
			return err
		}
		err = Finish(ctx, taskID)
		if err != nil {
			return err
		}
		return nil
	}, tags...)
}

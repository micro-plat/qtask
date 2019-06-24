package qtask

import (
	"github.com/micro-plat/hydra/context"
)

//Clear 清理任务，删除指定时间以前的任务
func Clear(day int) func(ctx *context.Context) (r interface{}) {
	return func(ctx *context.Context) (r interface{}) {
		ctx.Log.Infof("---------------qtask:清理%d天前的任务----------------", day)
		err := clearTask(ctx.GetContainer().GetRegularDB(dbName), day)
		if err != nil {
			return err
		}
		return "success"
	}
}

package order

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/qtask/qtask"
)

type PayHandler struct {
}

func NewPayHandler() (u *PayHandler) {
	return &PayHandler{}
}

//Handle .
func (u *PayHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------------------订单支付流程-------------------")

	if err := qtask.Processing(ctx, ctx.Request().GetInt64("task_id")); err != nil {
		return err
	}
	if ctx.Request().GetInt64("task_id")%2 == 0 {
		return fmt.Errorf("订单支付未完成")
	}

	if err := qtask.Finish(ctx, ctx.Request().GetInt64("task_id")); err != nil {
		return err
	}
	return "success"
}

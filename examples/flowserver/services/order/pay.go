package order

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/qtask"
)

type PayHandler struct {
	container component.IContainer
}

func NewPayHandler(container component.IContainer) (u *PayHandler) {
	return &PayHandler{container: container}
}

//Handle .
func (u *PayHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------------------订单支付流程-------------------")
	qtask.Processing(ctx, ctx.Request.GetInt64("task_id"))
	qtask.Finish(ctx, ctx.Request.GetInt64("task_id"))
	return "success"
}

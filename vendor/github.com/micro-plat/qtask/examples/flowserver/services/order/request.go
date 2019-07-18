package order

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/qtask/qtask"
)

type RequestHandler struct {
	container component.IContainer
}

func NewRequestHandler(container component.IContainer) (u *RequestHandler) {
	return &RequestHandler{container: container}
}

//Handle .
func (u *RequestHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("----------------订单创建---------------------")

	_, err := qtask.Create(ctx, "订单支付任务－立即", map[string]interface{}{
		"order_no": "87698990232",
	}, 300, "QTASK:TEST:ORDER-PAY", qtask.WithFirstTry(60), qtask.WithDeadline(3600))
	if err != nil {
		return err
	}

	return "success"
}
func (u *RequestHandler) DelayHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("----------------订单创建---------------------")

	_, err := qtask.Delay(ctx, "订单支付任务－延迟", map[string]interface{}{
		"order_no": "87698990232",
	}, 300, "QTASK:TEST:ORDER-PAY", qtask.WithFirstTry(60), qtask.WithDeadline(86400))
	if err != nil {
		return err
	}

	return "success"
}

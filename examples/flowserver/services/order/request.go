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
	queueName := ctx.Request.GetString("queuename")
	if queueName == "" {
		queueName = "QTASK:TEST:ORDER-PAY"
	}
	db, err := u.container.GetRegularDB().Begin()
	if err != nil {
		return err
	}
	_, callbacks, err := qtask.Create(db, "task_id", map[string]interface{}{
		"order_no": "87698990232",
	}, 300, queueName, qtask.WithDeadline(1000), qtask.WithDeleteDeadline(1000))
	if err != nil {
		db.Rollback()
		return err
	}
	// _, callback, err := qtask.Create(db, "task_id", map[string]interface{}{
	// 	"order_no": "87698990232",
	// }, 300, queueName, qtask.WithDeadline(1000), qtask.WithDeleteDeadline(1000))
	// if err != nil {
	// 	db.Rollback()
	// 	return err
	// }
	db.Commit()

	if err = callbacks(ctx); err != nil {
		return err
	}
	return "success"
}

// DelayHandle 延迟函数
func (u *RequestHandler) DelayHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("----------------订单创建---------------------")

	_, err := qtask.Delay(ctx, "订单支付任务－延迟", map[string]interface{}{
		"order_no": "87698990232",
	}, 300, 300, "QTASK:TEST:ORDER-PAY", qtask.WithDeadline(86400), qtask.WithDeleteDeadline(1000))
	if err != nil {
		return err
	}

	return "success"
}

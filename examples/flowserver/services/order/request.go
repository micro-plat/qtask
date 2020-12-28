package order

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/qtask/modules/const/conf"
	"github.com/micro-plat/qtask/qtask"
)

type RequestHandler struct {
}

func NewRequestHandler() (u *RequestHandler) {
	return &RequestHandler{}
}

//Handle .
func (u *RequestHandler) Handle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("----------------订单创建---------------------")
	queueName := ctx.Request().GetString("queuename")
	if queueName == "" {
		queueName = "QTASK:TEST:ORDER-PAY"
	}
	db, err := hydra.C.DB().GetRegularDB(conf.DBName).Begin()
	if err != nil {
		return err
	}
	_, callbacks, err := qtask.Create(db, "task_id", map[string]interface{}{
		"order_no": "87698990232",
	}, 300, queueName, qtask.WithDeadline(1000), qtask.WithDeleteDeadline(1000), qtask.WithMaxCount(1), qtask.WithOrderNO("111"))
	if err != nil {
		db.Rollback()
		return err
	}

	db.Commit()

	if err = callbacks(ctx); err != nil {
		return err
	}
	return "success"
}

// DelayHandle 延迟函数
func (u *RequestHandler) DelayHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("----------------订单创建---------------------")

	_, err := qtask.Delay(ctx, "订单支付任务－延迟", map[string]interface{}{
		"order_no": "87698990232",
	}, 300, 300, "QTASK:TEST:ORDER-PAY", qtask.WithDeadline(86400), qtask.WithDeleteDeadline(1000), qtask.WithOrderNO("111"))
	if err != nil {
		return err
	}

	return "success"
}

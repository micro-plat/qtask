package order

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask"
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
	queueName = types.GetString(queueName, "ORDER-PAY")
	db, err := hydra.C.DB().GetRegularDB(qtask.GetDBName()).Begin()
	if err != nil {
		return err
	}
	_, callbacks, err := qtask.Create(db, "task_id", map[string]interface{}{
		"order_no": "87698990232",
	}, 60, queueName, qtask.WithDeadline(1000), qtask.WithDeleteDeadline(1000), qtask.WithMaxCount(100), qtask.WithOrderNO("111"))
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
	}, 300, 300, global.MQConf.GetQueueName("ORDER-PAY"), qtask.WithDeadline(86400), qtask.WithDeleteDeadline(1000), qtask.WithOrderNO("111"))
	if err != nil {
		return err
	}

	return "success"
}

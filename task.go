package qtask

import (
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask/internal/modules/db"
)

//Create 创建实时任务，将任务信息保存到数据库并发送消息队列
//c:hydra.IContext,db.IDBExecuter,db.IDBTrans
//taskName:任务名称
//input:任务参数，序列化为json后存储，一般只允许传入关键参数。系统会在此输入参数中增加一个参数"task_id",业务流程需使用"task_id"修改任务为处理中或完结任务
//intervalTimeout:任务单次执行的超时时长，即下次执行时间距离上次执行时间的秒数，任务被放入消息队列、被标记为正在处理等都会自动调整下次执行时间
//mq: 消息队列名称
func Create(c interface{}, taskName string, input map[string]interface{}, intervalTimeout int, mq string, opts ...Option) (taskID int64, fcallback func(c interface{}) error, err error) {

	m, xdb, err := getTrans(c)
	if err != nil {
		return 0, nil, err
	}
	taskID, fcallback, err = create(xdb, c, taskName, input, intervalTimeout, mq, opts...)
	if !m {
		return taskID, fcallback, err
	}
	if err != nil {
		xdb.Rollback()
		return
	}
	xdb.Commit()
	return

}

//Delay 创建延迟任务，将任务信息保存到数据库，超时后将任务取出放到消息队列
//c:hydra.IContext,db.IDBExecuter,db.IDBTrans
//taskName:任务名称
//input:任务关健参数，序列化为json后存储，一般只允许传入关键参数。系统会在此输入参数中增加一个参数"task_id",业务流程需使用"task_id"修改任务为处理中或完结任务
// 首次执行时间
//intervalTimeout:任务单次执行的超时时长，即下次执行时间距离上次执行时间的秒数，任务被放入消息队列、被标记为正在处理等都会自动调整下次执行时间
//mq: 消息队列名称
//调用此函数将延后下次被放入消息队列的时间
func Delay(c interface{}, taskName string, input map[string]interface{}, firstTime int, intervalTimeout int, mq string, opts ...Option) (taskID int64, err error) {
	m, xdb, err := getTrans(c)
	if err != nil {
		return 0, err
	}
	taskID, err = delay(xdb, c, taskName, input, firstTime, intervalTimeout, mq, opts...)
	if !m {
		return taskID, err
	}
	if err != nil {
		xdb.Rollback()
		return
	}
	xdb.Commit()
	return

}

//ProcessingByInput 将任务修改为处理中。可以不调用，直接调用Finish完结任务，
//调用后系统自动延时，并累加任务处理次数
//任务被正式处理前调用此函数
//调用后当下次执行时间小于当前时间后会重新放入消息队列进行处理
func ProcessingByInput(c interface{}, input types.IXMap) error {
	return Processing(c, input.GetInt64("task_id"))
}

//Processing 将任务修改为处理中。可以不调用，直接调用Finish完结任务，
//调用后系统自动延时，并累加任务处理次数
//任务被正式处理前调用此函数
//调用后当下次执行时间小于当前时间后会重新放入消息队列进行处理
func Processing(c interface{}, taskID int64) error {
	m, xdb, err := getTrans(c)
	if err != nil {
		return err
	}
	err = db.Processing(xdb, taskID)
	if !m {
		return err
	}
	if err != nil {
		xdb.Rollback()
		return err
	}
	xdb.Commit()
	return nil
}

//FinishByInput 任务完成
//任务终结，不再放入消息队列
func FinishByInput(c interface{}, input types.IXMap) error {
	return Finish(c, input.GetInt64("task_id"))
}

//Finish 任务完成
//任务终结，不再放入消息队列
func Finish(c interface{}, taskID int64) error {
	m, xdb, err := getTrans(c)
	if err != nil {
		return err
	}
	err = db.Finish(xdb, taskID)
	if !m {
		return err
	}
	if err != nil {
		xdb.Rollback()
		return err
	}
	xdb.Commit()
	return nil
}

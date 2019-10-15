package qtask

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	ldb "github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/queue"
	"github.com/micro-plat/qtask/qtask/db"
)

//Create 创建实时任务，将任务信息保存到数据库并发送消息队列
//c:*context.Context,component.IContainer,db.IDBExecuter
//name:任务名称
//input:任务关健参数，序列化为json后存储，一般只允许传入关键参数。系统会在此输入参数中增加一个参数"task_id",业务流程需使用"task_id"修改任务为处理中或完结任务
//intervalTimeout:任务单次执行的超时时长，即下次执行时间距离上次执行时间的秒数，任务被放入消息队列、被标记为正在处理等都会自动调整下次执行时间
//mq: 消息队列名称
func Create(c interface{}, name string, input map[string]interface{}, intervalTimeout int, mq string, opts ...Option) (taskID int64, err error) {

	m, xdb, err := getTrans(c)
	if err != nil {
		return 0, err
	}
	taskID, err = create(xdb, c, name, input, intervalTimeout, mq, opts...)
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

//Delay 创建延迟任务，将任务信息保存到数据库，超时后将任务取出放到消息队列
//调用此函数将延后下次被放入消息队列的时间
func Delay(c interface{}, name string, input map[string]interface{}, intervalTimeout int, mq string, opts ...Option) (taskID int64, err error) {
	m, xdb, err := getTrans(c)
	if err != nil {
		return 0, err
	}
	taskID, err = delay(xdb, c, name, input, intervalTimeout, mq, opts...)
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

//Processing 将任务修改为处理中。系统自动延时，并累加任务处理次数
//任务被正式处理前调用此函数
//调用后当下次执行时间小于当前时间后会重新放入消息队列进行处理
func Processing(c interface{}, taskID int64) error {
	m, xdb, err := getTrans(c)
	if err != nil {
		return err
	}
	err = db.ProcessingTask(xdb, taskID)
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

//Finish 任务完成
//任务终结，不再放入消息队列
func Finish(c interface{}, taskID int64) error {
	m, xdb, err := getTrans(c)
	if err != nil {
		return err
	}
	err = db.FinishTask(xdb, taskID)
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

func getTrans(c interface{}) (bool, ldb.IDBTrans, error) {
	b, e, err := getDB(c)
	if err != nil {
		return false, nil, err
	}
	if b {
		return false, e.(ldb.IDBTrans), nil
	}
	t, err := e.(ldb.IDB).Begin()
	if err != nil {
		return false, nil, err
	}
	return true, t, nil
}

//---------------------------------内部函数-----------------------------------
func getDB(c interface{}) (bool, ldb.IDBExecuter, error) {
	switch v := c.(type) {
	case *context.Context:
		db, err := v.GetContainer().GetDB(dbName)
		return false, db, err
	case component.IContainer:
		db, err := v.GetDB(dbName)
		return false, db, err
	case ldb.IDB:
		return false, v, nil
	case ldb.IDBTrans:
		return true, v, nil
	default:
		return false, nil, fmt.Errorf("不支持的参数类型")
	}
}
func getQueue(c interface{}) (db queue.IQueue, err error) {
	switch v := c.(type) {
	case *context.Context:
		return v.GetContainer().GetQueue(queueName)
	case component.IContainer:
		return v.GetQueue(queueName)
	case queue.IQueue:
		return v, nil
	default:
		return nil, fmt.Errorf("不支持的参数类型")
	}
}

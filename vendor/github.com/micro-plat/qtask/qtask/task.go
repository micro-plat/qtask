package qtask

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	ldb "github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/jsons"
	"github.com/micro-plat/lib4go/queue"
	"github.com/micro-plat/qtask/qtask/db"
)

//Create 创建实时任务，将任务信息保存到数据库并发送消息队列
//c:*context.Context,component.IContainer,db.IDB
//name:任务名称
//input:任务关健参数，序列化为json后存储，一般只允许传入关键参数。系统会在此输入参数中增加一个参数"task_id",业务流程需使用"task_id"修改任务为处理中或完结任务
//intervalTimeout:任务单次执行的超时时长，即下次执行时间距离上次执行时间的秒数，任务被放入消息队列、被标记为正在处理等都会自动调整下次执行时间
//mq: 消息队列名称
func Create(c interface{}, name string, input map[string]interface{}, intervalTimeout int, mq string, opts ...Option) (taskID int64, err error) {

	xdb, err := getDB(c)
	if err != nil {
		return 0, err
	}

	args := make(map[string]interface{})
	for _, opt := range opts {
		opt(args)
	}

	//保存任务
	taskID, err = db.SaveTask(xdb, name, input, intervalTimeout, mq, args)
	if err != nil {
		return 0, err
	}

	//发送到消息队列
	input["task_id"] = taskID
	buff, err := jsons.Marshal(input)
	if err != nil {
		return 0, fmt.Errorf("任务输入参数转换为json失败:%v(%+v)", err, input)
	}
	queue, err := getQueue(c)
	if err != nil {
		return 0, err
	}
	return taskID, queue.Push(mq, string(buff))
}

//Delay 创建延迟任务，将任务信息保存到数据库，超时后将任务取出放到消息队列
//调用此函数将延后下次被放入消息队列的时间
func Delay(c interface{}, name string, input map[string]interface{}, intervalTimeout int, mq string, opts ...Option) (taskID int64, err error) {
	xdb, err := getDB(c)
	if err != nil {
		return 0, err
	}
	args := make(map[string]interface{})
	for _, opt := range opts {
		opt(args)
	}

	return db.SaveTask(xdb, name, input, intervalTimeout, mq, args)
}

//Processing 将任务修改为处理中。系统自动延时，并累加任务处理次数
//任务被正式处理前调用此函数
//调用后当下次执行时间小于当前时间后会重新放入消息队列进行处理
func Processing(c interface{}, taskID int64) error {
	xdb, err := getDB(c)
	if err != nil {
		return err
	}
	return db.ProcessingTask(xdb, taskID)

}

//Finish 任务完成
//任务终结，不再放入消息队列
func Finish(c interface{}, taskID int64) error {
	xdb, err := getDB(c)
	if err != nil {
		return err
	}
	return db.FinishTask(xdb, taskID)
}

//---------------------------------内部函数-----------------------------------
func getDB(c interface{}) (ldb.IDB, error) {
	switch v := c.(type) {
	case *context.Context:
		return v.GetContainer().GetDB(dbName)
	case component.IContainer:
		return v.GetDB(dbName)
	case ldb.IDB:
		return v, nil
	default:
		return nil, fmt.Errorf("不支持的参数类型")
	}
}
func getQueue(c interface{}) (db queue.IQueue, err error) {
	switch v := c.(type) {
	case *context.Context:
		return v.GetContainer().GetQueue(queueName)
	case component.IContainer:
		return v.GetQueue(dbName)
	case queue.IQueue:
		return v, nil
	default:
		return nil, fmt.Errorf("不支持的参数类型")
	}
}

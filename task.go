package qtask

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

//Create 创建任务,返回任务编号和错误信息
//c:*context.Context,component.IContainer,db.IDB
//name:任务名称
//input:任务关健参数，序列化为json后存储，一般只允许传入关键参数。系统会在此输入参数中增加一个参数"task_id",业务流程需使用"task_id"修改任务为处理中或完结任务
//timeout:任务单次执行的超时时长，即下次执行时间距离上次执行时间的秒数，任务被放入消息队列、被标记为正在处理等都会自动调整下次执行时间
//mq: 消息队列名称
func Create(c interface{}, name string, input map[string]interface{}, timeout int, mq string) (taskID int64, err error) {

	db,err:=getDB(c)
	if err!=nil{
		return 0,err
	}
	return createTask(db, name, input, timeout, mq)	

}

//Delay 修改任务的下次执行时间，任务未完成时可调用此函数。
//调用此函数将延后下次被放入消息队列的时间
func Delay(c interface{},taskID int64) error {
	db,err:=getDB(c)
	if err!=nil{
		return err
	}
	return delayTask(db, taskID)
}

//Processing 将任务修改为处理中。系统自动延时，并累加任务处理次数
//任务被正式处理前调用此函数
//调用后当下次执行时间小于当前时间后会重新放入消息队列进行处理
func Processing(c interface{}, taskID int64) error {
	db,err:=getDB(c)
	if err!=nil{
		return err
	}
	return processingTask(db, taskID)
}

//Finish 任务完成
//任务终结，不再放入消息队列
func Finish(c interface{},taskID int64) error {
	db,err:=getDB(c)
	if err!=nil{
		return err
	}
	return finishTask(db, taskID)
}

func getDB(c interface{})(db db.IDB,error){
	switch v := c.(type) {
	case *context.Context:
		return v.GetContainer().GetRegularDB(dbName),nil
	case component.IContainer:
		return v.GetRegularDB(dbName),nil
	case db.IDB:
		return v,nil
	default:
		return nil,fmt.Errorf("不支持的参数类型")
	}

}
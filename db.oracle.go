// +build oci

package qtask

import (
	"fmt"

	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"

	"github.com/micro-plat/lib4go/jsons"
)

//自定义安装程序
func CreateDB(c interface{}) error {
	xdb, err := getDB(c)
	if err != nil {
		return err
	}
	return db.CreateDB(xdb, "src/github.com/micro-plat/qtask/sql/oracle")
}

func saveTask(db db.IDB, name string, input map[string]interface{}, timeout int, mq string) (int64, error) {
	imap := map[string]interface{}{
		"name": name,
		"seq":  taskSeqName,
	}

	//获取任务编号
	taskID, _, _, err := db.Scalar(sqlGetSEQ, imap)
	if err != nil {
		return 0, fmt.Errorf("获取任务(%s)编号失败 %v", name, err)
	}

	//处理任务参数
	input["task_id"] = taskID
	buff, err := jsons.Marshal(input)
	if err != nil {
		return 0, fmt.Errorf("任务输入参数转换为json失败:%v(%+v)", err, input)
	}
	imap["content"] = string(buff)
	imap["task_id"] = taskID
	imap["next_interval"] = timeout
	imap["queue_name"] = mq

	//保存任务信息
	row, s, p, err := db.Execute(sqlCreateTaskID, imap)
	if err != nil || row != 1 {
		return 0, fmt.Errorf("创建任务(%s)失败 %v %s,%v", name, err, s, p)
	}
	return types.GetInt64(taskID), nil
}

func queryTasks(db db.IDB) (rows db.QueryRows, err error) {
	imap := map[string]interface{}{
		"name": "获取任务列表",
		"seq":  taskSeqName,
	}

	//获取任务编号
	batchID, s, _, err := db.Scalar(sqlGetBatch, imap)
	if err != nil {
		return nil, fmt.Errorf("获取批次编号失败 %v %s", err, s)
	}

	imap["batch_id"] = batchID

	row, _, _, err := db.Execute(sqlUpdateTask, imap)
	if err != nil {
		return nil, fmt.Errorf("修改任务批次失败 %v", err)
	}
	if row == 0 {
		return nil, context.NewError(204, "未查询到待处理任务")
	}
	rows, _, _, err = db.Query(sqlQueryWaitProcess, imap)
	if err != nil {
		return nil, fmt.Errorf("根据批次查询任务失败 %v", err)
	}

	return rows, nil
}

func clearTask(db db.IDB, day int) error {
	input := map[string]interface{}{
		"day": day,
	}
	rows, _, _, err := db.Execute(sqlClearTask, input)
	if err != nil {
		return fmt.Errorf("清理%d天前的任务失败 %v", day, err)
	}
	if rows == 0 {
		return context.NewError(204, "无需清理")
	}
	return nil
}

//-----------------------SQL-----------------------------------------
const sqlGetSEQ = `select seq_qtask_system_task_id.nextval from dual`
const sqlGetBatch = `select seq_qtask_system_task_batch_id.nextval from dual`

const sqlCreateTaskID = `insert into tsk_system_task
  (task_id,
   name,
   next_execute_time,
   max_execute_time,
   next_interval,
   status,
   queue_name,
   msg_content)
values
  (@task_id,
   @name,
   sysdate + #next_interval / 24 / 60 / 60,
   sysdate+3,
   @next_interval,
   20,
   @queue_name,
   @content)
`

const sqlProcessingTask = `update tsk_system_task t set t.next_execute_time=sysdate+t.next_interval/24/60/60,
t.status=30,t.count=t.count + 1,t.last_execute_time=sysdate
where t.status in(20,30) and t.task_id=@task_id`

const sqlFinishTask = `update tsk_system_task t set t.next_execute_time= to_date('2099-12-31', '%yyyy-%mm-%dd'),
t.status=0
where t.status in(20,30) and t.task_id=@task_id`

const sqlUpdateTask = `update tsk_system_task t set t.batch_id=@batch_id,t.next_execute_time= sysdate+t.next_interval/24/60/60
where t.status in(20,30) and t.next_execute_time <= sysdate and t.max_execute_time > sysdate 
and rownum<=200`

const sqlQueryWaitProcess = `select queue_name,msg_content content from tsk_system_task t where t.batch_id=@batch_id
and t.next_execute_time > sysdate`

const sqlClearTask = `delete from tsk_system_task t 
where t.create_time < sysdate - #day`

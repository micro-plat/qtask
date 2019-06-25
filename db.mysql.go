// +build !oci

package qtask

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"

	"github.com/micro-plat/lib4go/jsons"
)

//自定义安装程序
func CreateDB(c interface{}) error {
	db, err := getDB(c)
	if err != nil {
		return err
	}
	path, err := getSQLPath("mysql")
	if err != nil {
		return err
	}
	sqls, err := getSQL(path)
	if err != nil {
		return err
	}
	for _, sql := range sqls {
		if sql != "" {
			if _, q, _, err := db.Execute(sql, map[string]interface{}{}); err != nil {
				if !strings.Contains(err.Error(), "1050") && !strings.Contains(err.Error(), "1061") && !strings.Contains(err.Error(), "1091") {
					return fmt.Errorf("执行SQL失败： %v %s", err, q)
				}
			}
		}
	}
	return nil
}

func saveTask(db db.IDB, name string, input map[string]interface{}, timeout int, mq string) (taskID int64, err error) {
	imap := map[string]interface{}{
		"name": name,
	}

	//获取任务编号
	taskID, row, _, _, err := db.Executes(sqlGetSEQ, imap)
	if err != nil || row != 1 {
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
	imap["interval"] = timeout
	imap["queue_name"] = mq

	//保存任务信息
	row, _, _, err = db.Execute(sqlCreateTaskID, imap)
	if err != nil || row != 1 {
		return 0, fmt.Errorf("创建任务(%s)失败 %v", name, err)
	}
	return taskID, nil
}

func queryTasks(db db.IDB) (rows db.QueryRows, err error) {
	imap := map[string]interface{}{
		"name": "获取任务列表",
	}

	//获取任务编号
	batchID, row, _, _, err := db.Executes(sqlGetSEQ, imap)
	if err != nil || row != 1 {
		return nil, fmt.Errorf("获取批次编号失败 %v", err)
	}

	imap["batch_id"] = batchID

	row, _, _, err = db.Execute(sqlUpdateTask, imap)
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
	_, _, _, err = db.Execute(sqlClearSEQ, input)
	if err != nil {
		return fmt.Errorf("清理%d天前的序列数据失败 %v", day, err)
	}
	return nil
}

//-----------------------SQL-----------------------------------------
const sqlGetSEQ = `insert into tsk_system_seq (name) values (@name)`
const sqlCreateTaskID = `insert into tsk_system_task(task_id,name,next_execute_time,max_execute_time,
	interval,status,queue_name,msg_content)values(
	@task_id,@name,date_add(now(),interval @interval second),date_add(now(),interval 86400 second),
	@interval,20,@queue_name,@content)`

const sqlDelayTask = `update tsk_system_task t set t.next_execute_time= date_add(now(),interval t.interval second)
where t.status in(20,30) and t.task_id=@task_id`

const sqlProcessingTask = `update tsk_system_task t set t.next_execute_time=date_add(now(),interval t.interval second),
t.status=30,t.count=t.count + 1,t.last_execute_time=now()
where t.status in(20,30) and t.task_id=@task_id`

const sqlFinishTask = `update tsk_system_task t set t.next_execute_time= STR_TO_DATE('2099-12-31', '%Y-%m-%d'),
t.status=0
where t.status in(20,30) and t.task_id=@task_id`

const sqlUpdateTask = `update tsk_system_task t set t.batch_id=@batch_id,t.next_execute_time= date_add(now(),interval t.interval second)
where t.status in(20,30) and t.max_execute_time>now()`

const sqlQueryWaitProcess = `select queue_name　name,msg_content content from tsk_system_seq t where t.batch_id=@batch_id
and t.next_execute_time>now()`

const sqlClearTask = `delete from tsk_system_task t 
where t.create_time < date_add(now(),interval -#day day)`

const sqlClearSEQ = `delete from tsk_system_seq t 
where t.create_time < date_add(now(),interval -#day day)`

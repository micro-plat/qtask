// +build oci

package qtask

import (
	"fmt"
	"strings"

	"github.com/micro-plat/lib4go/db"

	"github.com/micro-plat/lib4go/jsons"
)

//自定义安装程序
func CreateDB(c interface{}) error {
	db, err := getDB(c)
	if err != nil {
		return 0, err
	}
	path, err := getSQLPath("oracle")
	if err != nil {
		return err
	}
	sqls, err := s.Conf.GetSQL(path)
	if err != nil {
		return err
	}
	db, err := c.GetDB()
	if err != nil {
		return err
	}
	for _, sql := range sqls {
		if sql != "" {
			if _, q, _, err := db.Execute(sql, map[string]interface{}{}); err != nil {
				if !strings.Contains(err.Error(), "ORA-00942") {
					s.Conf.Log.Errorf("执行SQL失败： %v %s\n", err, q)
				}
			}
		}
	}
	return nil
}

func createTask(db db.IDB, name string, input map[string]interface{}, timeout int, mq string) (taskID int64, err error) {
	imap := map[string]interface{}{
		"name": name,
		"seq":  taskSeqName,
	}

	//获取任务编号
	taskID, _, _, err := db.Scalar(sqlGetSEQ, imap)
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

//-----------------------SQL-----------------------------------------
const sqlGetSEQ = `select #seq.nextval from dual`

const sqlCreateTaskID = `insert into tsk_system_task(task_id,name,next_execute_time,max_execute_time,
	interval,status,queue_name,msg_content)values(
	@task_id,@name,sysdate+@interval/24/60/60,
	@interval,20,@queue_name,@content)`

const sqlDelayTask = `update tsk_system_task t set t.next_execute_time= sysdate+t.interval/24/60/60
where t.status in(20,30) and t.task_id=@task_id`

const sqlProcessingTask = `update tsk_system_task t set t.next_execute_time=sysdate+t.interval/24/60/60,
t.status=30,t.count=t.count + 1,t.last_execute_time=sysdate
where t.status in(20,30) and t.task_id=@task_id`

const sqlFinishTask = `update tsk_system_task t set t.next_execute_time= STR_TO_DATE('2099-12-31', '%Y-%m-%d'),
t.status=0
where t.status in(20,30) and t.task_id=@task_id`

const sqlUpdateTask = `update tsk_system_task t set t.batch_id=@batch_id,t.next_execute_time= sysdate+t.interval/24/60/60
where t.status in(20,30) and t.max_execute_time > sysdate`

const sqlQueryWaitProcess = `select queue_name　name,msg_content content from tsk_system_seq t where t.batch_id=@batch_id
and t.next_execute_time > sysdate`

const sqlClearTask = `delete from tsk_system_seq t 
where t.create_time < sysdate - #day`

const sqlClearSEQ = `delete from tsk_system_seq t 
where t.create_time < sysdate - #day`

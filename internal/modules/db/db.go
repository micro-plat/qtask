package db

import (
	"fmt"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/jsons"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask/internal/modules/const/conf"
	"github.com/micro-plat/qtask/internal/modules/const/sql"
)

// Processing 开始处理任务
func Processing(db db.IDBExecuter, taskID int64) error {
	imap := map[string]interface{}{
		"task_id": taskID,
	}
	row, err := db.Execute(sql.SQLProcessingTask, imap)
	if err != nil || row != 1 {
		return fmt.Errorf("修改任务为处理中(%d)失败 %v", taskID, err)
	}
	return nil
}

// Finish 结束任务
func Finish(db db.IDBExecuter, taskID int64) error {
	imap := map[string]interface{}{
		"task_id": taskID,
	}
	_, err := db.Execute(sql.SQLFinishTask, imap)
	if err != nil {
		return fmt.Errorf("关闭任务(%d)失败 %v", taskID, err)
	}
	return nil
}

// SaveTask 创建任务
func SaveTask(db db.IDBExecuter, name string, input map[string]interface{}, timeout int, mq string, args map[string]interface{}) (taskID int64, err error) {

	imap := map[string]interface{}{
		"name": name,
	}
	for k, v := range args {
		imap[k] = v
	}
	//获取任务编号
	taskID, err = getNewID(db, sql.SQLGetSEQ, imap)
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
	imap["first_timeout"] = types.DecodeInt(imap["first_timeout"], nil, timeout, imap["first_timeout"])
	imap["max_timeout"] = types.DecodeInt(imap["max_timeout"], nil, 259200, imap["max_timeout"])
	imap["delete_interval"] = types.DecodeInt(imap["delete_interval"], nil, 0, imap["delete_interval"])
	imap["max_count"] = types.DecodeInt(imap["max_count"], nil, 100, imap["max_count"])
	imap["queue_name"] = mq
	imap["plat_name"] = conf.GetPlatName()
	//保存任务信息
	row, err := db.Execute(sql.SQLCreateTask, imap)
	if err != nil || row != 1 {
		return 0, fmt.Errorf("创建任务(%s)失败 %v", name, err)
	}
	return types.GetInt64(taskID), nil
}

// ClearTask 清除任务
func ClearTask(db db.IDBExecuter) error {

	rows, err := db.Execute(sql.SQLClearTask, nil)
	if err != nil {
		return fmt.Errorf("清理任务失败 %v", err)
	}
	if rows == 0 {
		return errs.NewError(204, "无需清理")
	}
	return nil
}

// query 查询执行任务
func query(db db.IDBExecuter, SQLGetBatch string, SQLUpdateTask string, SQLQueryWaitProcess string) (batchID int64, rows db.QueryRows, err error) {
	imap := map[string]interface{}{
		"name": "获取任务列表",
	}

	//获取任务编号
	batchID, err = getNewID(db, SQLGetBatch, imap)
	if err != nil {
		return 0, nil, err
	}

	imap["batch_id"] = batchID
	imap["plat_name"] = conf.GetPlatName()

	row, err := db.Execute(SQLUpdateTask, imap)
	if err != nil {
		return 0, nil, fmt.Errorf("修改任务批次失败 %v", err)
	}
	if row == 0 {
		return 0, nil, errs.NewError(204, "未查询到待处理任务")
	}
	rows, err = db.Query(SQLQueryWaitProcess, imap)
	if err != nil {
		return 0, nil, fmt.Errorf("根据批次查询任务失败 %v", err)
	}
	return batchID, rows, nil
}

// failedTasks 失败任务处理
func failedTasks(db db.IDBExecuter, SQLFailedTask string) error {
	_, err := db.Execute(SQLFailedTask, nil)
	if err != nil {
		return fmt.Errorf("修改失败任务批次发生异常,err:%v", err)
	}
	return nil
}

// QueryTasks 查询任务
func QueryTasks(db db.IDBExecuter) (rows db.QueryRows, err error) {

	// 失败任务处理
	if err := failedTasks(db, sql.SQLFailedTask); err != nil {
		return nil, err
	}
	// 查询正在执行任务
	_, rows, err = query(db, sql.SQLGetSEQ, sql.SQLUpdateTask, sql.SQLQueryWaitProcess)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

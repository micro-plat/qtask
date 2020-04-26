// +build !oracle

package db

import (
	"fmt"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/qtask/modules/const/sql"
)

// SaveTask 保存任务
func SaveTask(db db.IDBExecuter, name string, input map[string]interface{}, timeout int, mq string, args map[string]interface{}) (taskID int64, err error) {

	taskID, err = create(db, name, input, timeout, mq, args, sql.SQLGetSEQ, sql.SQLCreateTaskID)
	if err != nil {
		return
	}
	input["seq_id"] = taskID
	_, _, _, err = db.Execute(sql.SQLClearSEQ, input)
	if err != nil {
		return 0, fmt.Errorf("删除序列数据失败 %v", err)
	}
	return
}

// QueryTasks 查询任务
func QueryTasks(db db.IDBExecuter) (rows db.QueryRows, err error) {

	// 失败任务处理
	if err := failedTasks(db, sql.SQLFailedTask); err != nil {
		return nil, err
	}
	// 查询正在执行任务
	batchID, rows, err := query(db, sql.SQLGetSEQ, sql.SQLUpdateTask, sql.SQLQueryWaitProcess)
	if err != nil {
		return nil, err
	}
	return rows, clearSEQ(db, map[string]interface{}{"seq_id": batchID})
}

// ClearTask 清除任务
func ClearTask(db db.IDBExecuter) error {

	if err := clear(db, sql.SQLClearTask); err != nil {
		return err
	}

	return clearSEQ(db, nil)
}

// getNewID 获取新ID
func getNewID(db db.IDBExecuter, SQLGetSEQ string, imap map[string]interface{}) (taskID int64, err error) {
	id, row, _, _, err := db.Executes(SQLGetSEQ, imap)
	if err != nil || row != 1 {
		return 0, fmt.Errorf("获取批次编号失败 %v", err)
	}
	return id, nil
}

// clearSEQ 清理序列
func clearSEQ(db db.IDBExecuter, imap map[string]interface{}) error {
	_, _, _, err := db.Execute(sql.SQLClearSEQ, imap)
	if err != nil {
		return fmt.Errorf("清理序列数据失败%v", err)
	}
	return nil
}

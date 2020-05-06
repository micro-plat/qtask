// +build !oracle

package db

import (
	"fmt"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/qtask/modules/const/sql"
)

// SaveTask 保存任务
func SaveTask(db db.IDBExecuter, name string, input map[string]interface{}, timeout int, mq string, args map[string]interface{}) (taskID int64, err error) {

	return create(db, name, input, timeout, mq, args, sql.SQLGetSEQ, sql.SQLCreateTaskID)
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

// ClearTask 清除任务
func ClearTask(db db.IDBExecuter) error {

	return clear(db, sql.SQLClearTask)
}

// getNewID 获取新ID
func getNewID(db db.IDBExecuter, SQLGetSEQ string, imap map[string]interface{}) (taskID int64, err error) {
	id, row, _, _, err := db.Executes(SQLGetSEQ, imap)
	if err != nil || row != 1 {
		return 0, fmt.Errorf("获取批次编号失败 %v", err)
	}
	return id, nil
}

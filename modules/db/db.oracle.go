// +build oracle

package db

import (
	"fmt"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask/modules/const/sql"
)

func SaveTask(db db.IDBExecuter, name string, input map[string]interface{}, timeout int, mq string, args map[string]interface{}) (int64, error) {
	return create(db, name, input, timeout, mq, args, sql.SQLGetSEQ, sql.SQLCreateTaskID)
}

func QueryTasks(db db.IDBExecuter) (rows db.QueryRows, err error) {
	_, rows, err = query(db, sql.SQLGetBatch, sql.SQLUpdateTask, sql.SQLQueryWaitProcess)
	return rows, err
}

// ClearTask 清除任务
func ClearTask(db db.IDBExecuter, day int) error {
	return clear(db, day, sql.SQLClearTask)
}

// getNewID 获取新ID
func getNewID(db db.IDBExecuter, SQLGetSEQ string, imap map[string]interface{}) (taskID int64, err error) {
	id, _, _, err := db.Scalar(SQLGetSEQ, imap)
	if err != nil {
		return 0, fmt.Errorf("获取任务(%-s)编号失败 %v", imap, err)
	}
	return types.GetInt64(id), nil
}

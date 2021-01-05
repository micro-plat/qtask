// +build !oracle

package db

import (
	"fmt"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/qtask/internal/modules/const/sql"
)

// getNewID 获取新ID
func getNewID(db db.IDBExecuter, SQLGetSEQ string, imap map[string]interface{}) (taskID int64, err error) {
	id, row, err := db.Executes(SQLGetSEQ, imap)
	if err != nil || row < 1 {
		return 0, fmt.Errorf("获取批次编号失败 %v", err)
	}
	_, err = db.Execute(sql.SQLClearSEQ, map[string]interface{}{"seq_id": id})
	if err != nil || row < 1 {
		return 0, fmt.Errorf("清理序列数据失败%v", err)
	}
	return id, nil
}

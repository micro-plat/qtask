// +build oracle

package db

import (
	"fmt"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

// getNewID 获取新ID
func getNewID(db db.IDBExecuter, SQLGetSEQ string, imap map[string]interface{}) (taskID int64, err error) {
	id, err := db.Scalar(SQLGetSEQ, imap)
	if err != nil {
		return 0, fmt.Errorf("获取任务(%-s)编号失败 %v", imap, err)
	}
	return types.GetInt64(id), nil
}

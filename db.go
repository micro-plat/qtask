package qtask

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/hydra/registry"
	"github.com/micro-plat/lib4go/db"
)

func delayTask(db db.IDB, taskID int64) error {
	imap := map[string]interface{}{
		"task_id": taskID,
	}
	row, _, _, err := db.Execute(sqlDelayTask, imap)
	if err != nil || row != 1 {
		return fmt.Errorf("延长下次执行时间(%d)失败 %v", taskID, err)
	}
	return nil
}
func processingTask(db db.IDB, taskID int64) error {
	imap := map[string]interface{}{
		"task_id": taskID,
	}
	row, _, _, err := db.Execute(sqlProcessingTask, imap)
	if err != nil || row != 1 {
		return fmt.Errorf("修改任务为处理中(%d)失败 %v", taskID, err)
	}
	return nil
}

func finishTask(db db.IDB, taskID int64) error {
	imap := map[string]interface{}{
		"task_id": taskID,
	}
	row, _, _, err := db.Execute(sqlFinishTask, imap)
	if err != nil || row != 1 {
		return fmt.Errorf("关闭任务(%d)失败 %v", taskID, err)
	}
	return nil
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

//getSQLPath 获取getSQLPath
func getSQLPath(tp string) (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", fmt.Errorf("未配置环境变量GOPATH")
	}
	path := strings.Split(gopath, ";")
	if len(path) == 0 {
		return "", fmt.Errorf("环境变量GOPATH配置的路径为空")
	}
	return filepath.Join(path[0], "src/github.com/micro-plat/qtask/sql/"+tp), nil
}
func getSQL(dir string) ([]string, error) {
	files, err := filepath.Glob(registry.Join(dir, "*.sql"))
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBufferString("")
	for _, f := range files {
		buf, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}
		_, err = buff.Write(buf)
		if err != nil {
			return nil, err
		}
		buff.WriteString(";")
	}
	tables := make([]string, 0, 8)
	tbs := strings.Split(buff.String(), ";")
	for _, t := range tbs {
		if tb := strings.TrimSpace(t); len(tb) > 0 {
			tables = append(tables, tb)
		}
	}
	return tables, nil
}

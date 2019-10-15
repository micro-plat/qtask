package qtask

import (
	"fmt"

	"github.com/micro-plat/lib4go/jsons"

	ldb "github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/qtask/qtask/db"
)

func create(xdb ldb.IDBExecuter, c interface{}, name string, input map[string]interface{}, intervalTimeout int, mq string, opts ...Option) (taskID int64, err error) {

	args := make(map[string]interface{})
	for _, opt := range opts {
		opt(args)
	}

	//保存任务
	taskID, err = db.SaveTask(xdb, name, input, intervalTimeout, mq, args)
	if err != nil {
		return 0, err
	}

	//发送到消息队列
	input["task_id"] = taskID
	buff, err := jsons.Marshal(input)
	if err != nil {
		return 0, fmt.Errorf("任务输入参数转换为json失败:%v(%+v)", err, input)
	}
	container, ok := args["container"]
	if !ok {
		container = c
	}
	queue, err := getQueue(container)
	if err != nil {
		return 0, err
	}
	return taskID, queue.Push(mq, string(buff))
}

func delay(xdb ldb.IDBExecuter, c interface{}, name string, input map[string]interface{}, intervalTimeout int, mq string, opts ...Option) (taskID int64, err error) {

	args := make(map[string]interface{})
	for _, opt := range opts {
		opt(args)
	}

	return db.SaveTask(xdb, name, input, intervalTimeout, mq, args)
}

package qtask

import (
	"fmt"

	ldb "github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/jsons"
	"github.com/micro-plat/qtask/qtask/db"
)

func create(xdb ldb.IDBExecuter, c interface{}, name string,
	input map[string]interface{}, intervalTimeout int, mq string, opts ...Option) (taskID int64, callback func(c interface{}) error, err error) {

	args := make(map[string]interface{})
	for _, opt := range opts {
		opt(args)
	}

	//保存任务
	taskID, err = db.SaveTask(xdb, name, input, intervalTimeout, mq, args)
	if err != nil {
		return 0, nil, err
	}

	//发送到消息队列
	input["task_id"] = taskID

	callback = func(c interface{}) error {
		buff, err := jsons.Marshal(input)
		if err != nil {
			return fmt.Errorf("任务输入参数转换为json失败:%v(%+v)", err, input)
		}
		queue, err := getQueue(c)
		if err != nil {
			return err
		}
		return queue.Push(mq, string(buff))
	}
	return taskID, callback, nil
}

func delay(xdb ldb.IDBExecuter, c interface{}, name string, input map[string]interface{}, intervalTimeout int, mq string, opts ...Option) (taskID int64, err error) {

	args := make(map[string]interface{})
	for _, opt := range opts {
		opt(args)
	}

	return db.SaveTask(xdb, name, input, intervalTimeout, mq, args)
}

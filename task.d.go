package qtask

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components/queues"
	ldb "github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/qtask/internal/modules/const/conf"
	"github.com/micro-plat/qtask/internal/modules/db"
)

func getTrans(c interface{}) (bool, ldb.IDBTrans, error) {
	b, e, err := getDB(c)
	if err != nil {
		return false, nil, err
	}
	if b {
		return false, e.(ldb.IDBTrans), nil
	}
	t, err := e.(ldb.IDB).Begin()
	if err != nil {
		return false, nil, err
	}
	return true, t, nil
}

//---------------------------------内部函数-----------------------------------
func getDB(c interface{}) (bool, ldb.IDBExecuter, error) {
	switch v := c.(type) {
	case ldb.IDB:
		return false, v, nil
	case ldb.IDBTrans:
		return true, v, nil
	default:
		db, err := hydra.C.DB().GetDB(conf.DBName)
		return false, db, err
	}
}
func getQueue(c interface{}) (db queues.IQueue, err error) {
	switch v := c.(type) {
	case queues.IQueue:
		return v, nil
	default:
		return hydra.C.Queue().GetQueue(conf.QueueName)
	}
}

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

	fcallback := func(input *map[string]interface{}) func(c interface{}) error {
		return func(c interface{}) error {
			queue, err := getQueue(c)
			if err != nil {
				return err
			}
			return queue.Send(mq, input)
		}
	}
	return taskID, fcallback(&input), nil
}

func delay(xdb ldb.IDBExecuter, c interface{}, name string, input map[string]interface{}, firstTime int, intervalTimeout int, mq string, opts ...Option) (taskID int64, err error) {

	args := make(map[string]interface{})
	for _, opt := range opts {
		opt(args)
	}
	args["first_timeout"] = firstTime
	return db.SaveTask(xdb, name, input, intervalTimeout, mq, args)
}

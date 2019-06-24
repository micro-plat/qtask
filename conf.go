package qtask

var dbName = "db"
var queueName = "queue"
var taskSeqName = "seq_qtask_system_task_id"

//Config　配置数据库，消息队列的配置名称
func Config(db string, mq string) {
	dbName = db
	queueName = mq
}

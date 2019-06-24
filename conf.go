package qtask

var dbName = "db"
var queueName = "queue"
var taskSeqName = ""

//Config　配置数据库，消息队列的配置名称
func Config(db string, mq string, orclTaskSeqName string) {
	dbName = db
	queueName = mq
	taskSeqName = orclTaskSeqName
}

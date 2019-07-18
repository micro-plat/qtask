package qtask

var dbName = "db"
var queueName = "queue"

//Config　配置数据库，消息队列的配置名称
func Config(db string, mq string) {
	dbName = db
	queueName = mq
}

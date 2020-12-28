package qtask

import (
	"fmt"

	"github.com/micro-plat/qtask/modules/const/conf"
)

// Config 配置数据库，消息队列的配置名称
func Config(opts ...ConfOption) {
	for _, opt := range opts {
		opt()
	}
}

//ConfOption 配置选项
type ConfOption func()

// WithScanInterval 秒数，后补时间间隔
func WithScanInterval(scanInterval int) ConfOption {
	return func() {
		if scanInterval >= 60 || scanInterval < 0 {
			panic(fmt.Sprintf("扫描时间取值为0-59,当前值:%d", scanInterval))
		}
		conf.ScanInterval = scanInterval
	}
}

// WithDBName 数据库节点名称
func WithDBName(dbName string) ConfOption {
	return func() {
		conf.DBName = dbName
	}
}

// WithQueueName 消息队列节点名称
func WithQueueName(queueName string) ConfOption {
	return func() {
		conf.QueueName = queueName
	}
}

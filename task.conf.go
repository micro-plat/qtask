package qtask

import (
	"fmt"

	"github.com/micro-plat/qtask/internal/modules/const/conf"
)

// Config 配置数据库，消息队列的配置名称,请通过hydra.OnReadyByInsert修改配置参数
func Config(opts ...ConfOption) {
	for _, opt := range opts {
		opt()
	}
}

//GetDBName 获取已配置的db节点名
func GetDBName() string {
	return conf.DBName
}

//GetQueueName 获取已配置的queue节点名
func GetQueueName() string {
	return conf.QueueName
}

//GetScanInterval 获取已配置的tScanInterval
func GetScanInterval() int {
	return conf.ScanInterval
}

//GetPlatName 获取队列对应的平台名称
func GetPlatName() string {
	return conf.GetPlatName()
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

// WithPlatName 设置消息队列对应平台的名称
func WithPlatName(platName string) ConfOption {
	return func() {
		conf.PlatName = platName
	}
}

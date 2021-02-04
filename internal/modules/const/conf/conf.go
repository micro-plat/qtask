package conf

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

// DBName 数据库名称
var DBName = "db"

// QueueName 队列名称
var QueueName = "queue"

//PlatName 平台名称
var PlatName = ""

//ScanInterval 扫描时长秒
var ScanInterval = 3

//GetPlatName 获取队列对应的平台名称
func GetPlatName() string {
	return types.GetString(PlatName, hydra.G.PlatName)
}

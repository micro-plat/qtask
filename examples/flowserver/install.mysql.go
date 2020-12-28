// +build !oracle

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/mqc"
	"github.com/micro-plat/hydra/conf/server/queue"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func init() {
	hydra.Conf.API("9090")
	hydra.Conf.Vars().DB().MySQLByConnStr("db", "oms_t:123456@tcp(192.168.0.36)/oms_t?charset=utf8")
	hydra.Conf.MQC(mqc.WithRedis("redis")).Queue(queue.NewQueue("QTASK:TEST:ORDER-PAY", "/order/pay"))
	hydra.Conf.Vars().Queue().Redis("redis", "192.168.0.111:6379", queueredis.WithAddrs("192.168.0.112:6379",
		"192.168.0.113:6379", "192.168.0.114:6379", "192.168.0.115:6379", "192.168.0.116:6379"))

}

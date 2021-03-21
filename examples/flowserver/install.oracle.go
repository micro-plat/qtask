// +build oracle

package main

import (
	_ "github.com/mattn/go-oci8"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/mqc"
	"github.com/micro-plat/hydra/conf/server/queue"
	"github.com/micro-plat/hydra/conf/vars/queue/queueredis"
)

func init() {
	hydra.OnReady(func() {
		if hydra.G.IsDebug() {
			hydra.Conf.API("9090")
			hydra.Conf.Vars().DB().MySQLByConnStr("db", "hydra/123456@orcl136")
			hydra.Conf.MQC(mqc.WithRedis("redis")).Queue(queue.NewQueue("ORDER-PAY", "/order/pay"))
			hydra.Conf.Vars().Queue().Redis("queue", "192.168.0.111:6379", queueredis.WithAddrs("192.168.0.112:6379",
				"192.168.0.113:6379", "192.168.0.114:6379", "192.168.0.115:6379", "192.168.0.116:6379"))
			return
		}
		hydra.Conf.API(hydra.ByInstall)
		hydra.Conf.Vars().DB().MySQLByConnStr("db", hydra.ByInstall)
		hydra.Conf.MQC(mqc.WithRedis("queue")).Queue(queue.NewQueue("ORDER-PAY", "/order/pay"))
		hydra.Conf.Vars().Queue().Redis("queue", hydra.ByInstall)
	})

}

// +build oracle

package main

import (
	_ "github.com/zkfy/go-oci8"
)

func init() {
	app.IsDebug = true

	app.Conf.API.SetMainConf(`{"address":":9090"}`)

	app.Conf.Plat.SetVarConf("db", "db", `{			
			"provider":"ora",
			"connString":"hydra/123456@orcl136",
			"maxOpen":20,
			"maxIdle":10,
			"lifeTime":600		
	}`)

	app.Conf.MQC.SetSubConf("queue", `{
		"queues":[
		  {
			  "queue":"QTASK:TEST:ORDER-PAY",
			  "service":"/order/pay",
			   "concurrency":1000
		  }]}`)
	app.Conf.MQC.SetSubConf("server", `
		{
			"proto":"redis",
			"addrs":[
					"192.168.0.111:6379",
					"192.168.0.112:6379",
					"192.168.0.113:6379",
					"192.168.0.114:6379",
					"192.168.0.115:6379",
					"192.168.0.116:6379"
			],
			"db":1,
			"dial_timeout":10,
			"read_timeout":10,
			"write_timeout":10,
			"pool_size":10
	}
	`)

	app.Conf.Plat.SetVarConf("queue", "queue", `
		{
			"proto":"redis",
			"addrs":[
					"192.168.0.111:6379",
					"192.168.0.112:6379",
					"192.168.0.113:6379",
					"192.168.0.114:6379",
					"192.168.0.115:6379",
					"192.168.0.116:6379"
			],
			"db":1,
			"dial_timeout":10,
			"read_timeout":10,
			"write_timeout":10,
			"pool_size":10
	}
	`)

}

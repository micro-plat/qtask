// +build !oracle

package main

import _ "github.com/go-sql-driver/mysql"

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func init() {
	app.IsDebug = true
	app.Conf.API.SetMainConf(`{"address":":9090"}`)
	app.Conf.Plat.SetVarConf("db", "db", `{			
			"provider":"mysql",
			"connString":"oms_t:123456@tcp(192.168.0.36)/oms_t?charset=utf8",
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

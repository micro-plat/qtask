// +build !prod

package main

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (s *flowserver) install() {
	//api.port#//
	s.Conf.API.SetMainConf(`{"address":":9090"}`)
	//#api.port//

	//api.appconf#//
	//#api.appconf//

	//api.cros#//
	//#api.cros//

	//api.jwt#//
	//#api.jwt//

	//api.metric#//
	//#api.metric//

	//db#//
	s.Conf.Plat.SetVarConf("db", "db", `{			
			"provider":"ora",
			"connString":"ecoupon/ecoupon@orcl136",
			"maxOpen":20,
			"maxIdle":10,
			"lifeTime":600		
	}`)
	//#db//

	//cache#//
	//#cache//

	//queue#//
	//#queue//

	//cron.appconf#//
	//#cron.appconf//

	//cron.task#//
	//#cron.task//

	//cron.metric#//
	//#cron.metric//

	//mqc.server#//
	//#mqc.server//

	//mqc.queue#//

	s.Conf.MQC.SetSubConf("queue", `{
		"queues":[
		  {
			  "queue":"QTASK:TEST:ORDER-PAY",
			  "service":"/sms/dispatch/request"
		  }]}`)

	s.Conf.Plat.SetVarConf("queue", "queue", `
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
	//#mqc.queue//

	//mqc.metric#//
	//#mqc.metric//

	//web.port#//
	//#web.port//

	//web.static#//
	//#web.static//

	//web.metric#//
	//#web.metric//

	//ws.appconf#//
	//#ws.appconf//

	//ws.jwt#//
	//#ws.jwt//

	//ws.metric#//
	//#ws.metric//

	//rpc.port#//
	//#rpc.port//

	//rpc.metric#//
	//#rpc.metric//
}

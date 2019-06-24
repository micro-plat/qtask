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
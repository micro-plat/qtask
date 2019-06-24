// +build prod

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
		
	//api.jwt.prod#//
	//#api.jwt.prod//
	
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
	
	//自定义安装程序
	s.Conf.API.Installer(func(c component.IContainer) error {
		if !s.Conf.Confirm("创建数据库表结构,添加基础数据?") {
			return nil
		}
		path, err := getSQLPath()
		if err != nil {
			return err
		}
		sqls, err := s.Conf.GetSQL(path)
		if err != nil {
			return err
		}
		db, err := c.GetDB()
		if err != nil {
			return err
		}
		for _, sql := range sqls {
			if sql != "" {
				if _, q, _, err := db.Execute(sql, map[string]interface{}{}); err != nil {
					if !strings.Contains(err.Error(), "ORA-00942") {
						s.Conf.Log.Errorf("执行SQL失败： %v %s\n", err, q)
					}
				}
			}
		}
		return nil
	})
}

//getSQLPath 获取getSQLPath
func getSQLPath() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", fmt.Errorf("未配置环境变量GOPATH")
	}
	path := strings.Split(gopath, ";")
	if len(path) == 0 {
		return "", fmt.Errorf("环境变量GOPATH配置的路径为空")
	}
	return filepath.Join(path[0], "src/github.com/micro-plat/qtask/examples/flowserver/modules/const/sql"), nil
}

package main
import (
	
	"github.com/micro-plat/hydra/component"
	

	)


//init 检查应用程序配置文件，并根据配置初始化服务
func (r *flowserver) init() {
	r.Initializing(func(c component.IContainer) error {
	//appconf.func#//
	//#appconf.func//

	//db.init#//
	//#db.init//

	//cache.init#//
	//#cache.init//

	//queue.init#//
	//#queue.init//

	
	//login.router#//
	//#login.router//

	//service.router#//
	//#service.router//

	return nil
	})
}
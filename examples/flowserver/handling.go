package main
import (
	"github.com/micro-plat/hydra/context"
	
)

//bind 检查应用程序配置文件，并根据配置初始化服务
func (r *flowserver) handling() {
	//每个请求执行前执行
	r.Handling(func(ctx *context.Context) (rt interface{}) {		
	//handling.jwt#//
	//#handling.jwt//

	
		return nil
	})
}

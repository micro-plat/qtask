package qtask

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/qtask/internal/modules/const/conf"
	"github.com/micro-plat/qtask/internal/services"
)

//BindFlow 绑定自动流程(调用后将自动扫描数据库后补队列数据，并定期清除过期数据)
func BindFlow() {
	hydra.S.CRON("/task/scan", services.Scan, fmt.Sprintf("@every %ds", conf.ScanInterval)) //定时扫描任务
	hydra.S.CRON("/task/clear", services.Clear(), "@daily")                                 //定时清理任务
}

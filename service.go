package qtask

import (
	"fmt"
	"sync"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/qtask/internal/modules/const/conf"
	"github.com/micro-plat/qtask/internal/services"
)

var once sync.Once

func init() {
	hydra.OnReady(func() {
		once.Do(func() {
			hydra.S.CRON("/task/scan", services.Scan)     //定时扫描任务
			hydra.S.CRON("/task/clear", services.Clear()) //定时清理任务

			hydra.CRON.Add(fmt.Sprintf("@every %ds", conf.ScanInterval), "/task/scan")
			hydra.CRON.Add("@daily", "/task/clear")
		})
	})

}

package main

import (
	"github.com/micro-plat/qtask"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
	"github.com/micro-plat/qtask/examples/flowserver/services/order"
)

var app = hydra.NewApp(
	hydra.WithPlatName("qtask"),
	hydra.WithSystemName("flowserver"),
	hydra.WithServerTypes(http.API, cron.CRON, mqc.MQC))

func main() {
	app.Micro("/order/request", order.NewRequestHandler)
	app.MQC("/order/pay", order.NewPayHandler)
	qtask.BindFlow()

	//检查配置是否正确
	app.OnStarting(func(conf hydra.IAPPConf) error {
		if _, err := hydra.C.DB().GetDB(); err != nil {
			return err
		}
		if _, err := hydra.C.Queue().GetQueue(); err != nil {
			return err
		}
		return nil
	})

	app.Start()
}

package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
)

var app = hydra.NewApp(
	hydra.WithPlatName("qtask"),
	hydra.WithSystemName("flowserver"),
	hydra.WithServerTypes(http.API, cron.CRON, mqc.MQC))

func main() {
	app.Start()
}

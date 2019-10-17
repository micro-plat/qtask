package main

import "github.com/micro-plat/hydra/hydra"

var app = hydra.NewApp(
	hydra.WithPlatName("qtask"),
	hydra.WithSystemName("flowserver"),
	hydra.WithServerTypes("api-cron-mqc"))

func main() {
	app.Start()
}

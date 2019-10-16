package main

import "github.com/micro-plat/hydra/hydra"

type flowserver struct {
	*hydra.MicroApp
}

var app = &flowserver{
	hydra.NewApp(
		hydra.WithPlatName("qtask"),
		hydra.WithSystemName("flowserver"),
		hydra.WithServerTypes("api-cron-mqc")),
}

func main() {
	app.Start()
}

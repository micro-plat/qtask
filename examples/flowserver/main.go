package main

import "github.com/micro-plat/hydra/hydra"

type flowserver struct {
	*hydra.MicroApp
}

func main() {
	app := &flowserver{
		hydra.NewApp(
			hydra.WithPlatName("github.com"),
			hydra.WithSystemName("flowserver"),
			hydra.WithServerTypes("api-cron-mqc"),
			hydra.WithDebug()),
	}

	app.init()
	app.install()
	app.handling()

	app.Start()
}

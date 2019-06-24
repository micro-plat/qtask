
package main

import "github.com/micro-plat/hydra/hydra"
import _ "github.com/go-sql-driver/mysql"

type flowserver struct {
	*hydra.MicroApp
}

func main() {
	app := &flowserver{
		hydra.NewApp(
			hydra.WithPlatName("github.com"),
			hydra.WithSystemName("flowserver"),
			hydra.WithServerTypes("api"),
			hydra.WithDebug()),
	}

	app.init()
	app.install()
	app.handling()

	app.Start()
}

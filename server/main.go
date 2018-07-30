package main

import (
	"github.com/liangdas/mqant"
	//	"github.com/liangdas/mqant/app"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/modules"
)

//var ShutdownApp = func(app module.App, code int) {
//	os.Exit(code)
//}
var OnLoaded = func(app module.App) {
}
var OnStartup = func(app module.App) {
}

func main() {
	app := mqant.CreateApp()
	app.OnConfigurationLoaded(OnLoaded)
	app.OnStartup(OnStartup)
	app.Run(true,
		modules.MasterModule(),
	)
}

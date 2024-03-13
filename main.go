package main

import (
	"openidea-shopyfyx/app"
	"openidea-shopyfyx/config"
)

func main() {
	config.EnvBinder()

	app.InitFiberApp()
}

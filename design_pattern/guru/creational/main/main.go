package main

import (
	"tucker_algorithm/design_pattern/guru/creational/app"
	"tucker_algorithm/design_pattern/guru/creational/factories"
)

func main() {
	macApp := configureApplication("mac")
	macApp.Paint()
	windowApp := configureApplication("window")
	windowApp.Paint()
}

func configureApplication(system string) *app.Application {
	switch system {
	case "mac":
		return app.NewApplication(&factories.MacOsFactory{})
	case "window":
		return app.NewApplication(&factories.WindowOsFactory{})
	default:
		panic("invalid system name")
	}
}

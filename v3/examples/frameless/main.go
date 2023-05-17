package main

import (
	"embed"
	_ "embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed assets
var assets embed.FS

func main() {

	app := application.New(application.Options{
		Name:        "Frameless Demo",
		Description: "A demo of frameless windows",
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
		Assets: application.AssetOptions{
			FS: assets,
		},
	})

	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Frameless: true,
	})

	err := app.Run()

	if err != nil {
		log.Fatal(err.Error())
	}
}

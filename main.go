package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"godesk/app/controller"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	config := controller.NewConfig()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "GoDesk",
		Width:  860,
		Height: 620,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        false,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
		},
		Bind: []interface{}{
			app,
			config,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

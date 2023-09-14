package main

import (
	"log/slog"
	"os"

	"gioui.org/app"
	"gioui.org/unit"

	"github.com/andrdru/giorun"
	"github.com/andrdru/giorun/examples/hello/controllers"
)

func main() {
	// init controller
	controller := controllers.NewController()

	// init giorun
	run := giorun.NewRun(
		controller.Widgets(),
		[]app.Option{
			app.Size(unit.Dp(800), unit.Dp(600)),
			app.Title("hello"),
		},
	)

	// set UI invalidate func
	controller.SetInvalidateUIFunc(run.InvalidateUI)

	// run event loop
	go func() {
		if err := run.Loop(); err != nil {
			slog.Default().Error(err.Error())
			os.Exit(1)
		}

		os.Exit(0)
	}()

	// start invalidate watcher and display app
	// this blocks main()
	run.Display()
}

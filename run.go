package giorun

import (
	"context"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/andrdru/giorun/watcher"
	runWidget "github.com/andrdru/giorun/widget"
)

type (
	//Run wrap gioui application
	Run struct {
		w  *app.Window
		th *material.Theme

		list    *widget.List
		widgets []runWidget.Item
		inset   layout.Inset

		invalidateWatcher invalidateWatcher
	}

	invalidateWatcher interface {
		Watch(ctx context.Context)
		Trigger(value bool)
	}
)

const (
	WatcherTimeoutDefault = 5 * time.Millisecond
)

var (
	WatcherFunc            func() invalidateWatcher
	LayoutUniformInsetFunc func() layout.Inset
	WidgetListFunc         func() layout.List
)

func NewRun(widgets []runWidget.Item, gioOpts []app.Option) *Run {
	if LayoutUniformInsetFunc == nil {
		LayoutUniformInsetFunc = func() layout.Inset { return layout.UniformInset(unit.Dp(16)) }
	}

	if WidgetListFunc == nil {
		WidgetListFunc = func() layout.List { return layout.List{Axis: layout.Vertical} }
	}

	ret := &Run{
		w:     app.NewWindow(gioOpts...),
		th:    material.NewTheme(),
		inset: LayoutUniformInsetFunc(),

		list: &widget.List{
			List: WidgetListFunc(),
		},

		widgets: widgets,
	}

	if WatcherFunc == nil {
		WatcherFunc = func() invalidateWatcher { return watcher.NewWatcher(ret.w.Invalidate, WatcherTimeoutDefault) }
	}

	ret.invalidateWatcher = WatcherFunc()

	return ret
}

func (ui *Run) Loop() error {
	var ops op.Ops

	for {
		select {
		case e := <-ui.w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err

			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)

				ui.Layout(gtx, ui.th)
				e.Frame(gtx.Ops)
			}
		}
	}
}

// Display show app window after init
// see app.Main for details
func (ui *Run) Display() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go ui.invalidateWatcher.Watch(ctx)

	app.Main()
}

// InvalidateUI re Run
func (ui *Run) InvalidateUI() {
	ui.invalidateWatcher.Trigger(true)
}

// Layout set gio ui layout
func (ui *Run) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	var widgetsList []layout.Widget

	for _, el := range ui.widgets {
		el.Handler()

		widgetsList = append(widgetsList, el.Widget(th))
	}

	return material.List(th, ui.list).Layout(gtx, len(widgetsList), func(gtx layout.Context, i int) layout.Dimensions {
		return ui.inset.Layout(gtx, widgetsList[i])
	})
}

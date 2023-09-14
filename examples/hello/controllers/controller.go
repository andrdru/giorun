package controllers

import (
	"github.com/andrdru/giorun/examples/hello/widgets"
	runWidget "github.com/andrdru/giorun/widget"
)

// Controller app controller, handle logic and views
type Controller struct {
	invalidateUIFunc func()

	// widgets are views
	widgets    []runWidget.Item
	widgetByID *runWidget.List

	greeting string
}

func NewController() *Controller {
	ret := &Controller{}

	// init widgets
	// handler passed to widget constructor
	ret.widgets = []runWidget.Item{
		widgets.NewHelloEditor(ret.handleHelloEditorEdit),
		widgets.NewPrintButton(ret.handleStartButtonClick),
		widgets.NewOutputEditor(),
	}

	ret.widgetByID = runWidget.NewList(ret.widgets)

	return ret
}

// SetInvalidateUIFunc invalidate func tells app to redraw layout
func (l *Controller) SetInvalidateUIFunc(invalidateUIFunc func()) {
	l.invalidateUIFunc = invalidateUIFunc
}

// Widgets widgets list for giorun.Run
func (l *Controller) Widgets() []runWidget.Item {
	return l.widgets
}

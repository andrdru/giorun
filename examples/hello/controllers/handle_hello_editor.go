package controllers

import (
	"github.com/andrdru/giorun/examples/hello/widgets"
)

func (l *Controller) handleHelloEditorEdit() {
	data := l.widgetByID.Get(widgets.HelloEditorID).(widgets.HelloEditorData)

	l.greeting = data.Value
}

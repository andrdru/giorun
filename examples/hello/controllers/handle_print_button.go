package controllers

import (
	"github.com/andrdru/giorun/examples/hello/widgets"
)

func (l *Controller) handleStartButtonClick() {
	l.widgetByID.Set(
		widgets.OutputEditorID,
		widgets.OutputEditorData{Value: l.greeting},
	)

	l.widgetByID.Set(
		widgets.HelloEditorID,
		widgets.HelloEditorData{Value: ""},
	)

	l.invalidateUIFunc()
}

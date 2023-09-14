package widgets

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	runWidget "github.com/andrdru/giorun/widget"
)

type (
	// PrintButton button view
	PrintButton struct {
		w *widget.Clickable

		handler func()
	}
)

const (
	PrintButtonID = "print_button"
)

var _ runWidget.Item = &PrintButton{}

func NewPrintButton(handler func()) *PrintButton {
	return &PrintButton{
		w:       &widget.Clickable{},
		handler: handler,
	}
}

func (l *PrintButton) ID() string {
	return PrintButtonID
}

func (l *PrintButton) Widget(th *material.Theme) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		b := material.Button(th, l.w, "Say")

		return b.Layout(gtx)
	}
}

func (l *PrintButton) Handler() {
	if l.w.Clicked() {
		l.handler()
	}
}

func (l *PrintButton) Set(data any) {
	// noop
}

func (l *PrintButton) Get() any {
	// noop
	return nil
}

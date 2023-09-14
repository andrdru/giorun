package widgets

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	runWidget "github.com/andrdru/giorun/widget"
)

type (
	// OutputEditor output readonly view
	OutputEditor struct {
		data OutputEditorData
	}

	OutputEditorData struct {
		Value string
	}
)

const (
	OutputEditorID = "output_editor"
)

var _ runWidget.Item = &HelloEditor{}

func NewOutputEditor() *OutputEditor {
	return &OutputEditor{}
}

func (l *OutputEditor) ID() string {
	return OutputEditorID
}

func (l *OutputEditor) Widget(th *material.Theme) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		// workaround: for some reason SetText causes panic when editor used for output
		// recreate widget.Editor on redraw
		w := &widget.Editor{
			ReadOnly: true,
		}

		w.SetText(l.data.Value)

		e := material.Editor(th, w, "")

		border := widget.Border{Color: color.NRGBA{A: 0xFF, R: 0xFF}, CornerRadius: unit.Dp(4), Width: unit.Dp(2)}

		return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(8)).Layout(gtx, e.Layout)
		})
	}
}

func (l *OutputEditor) Handler() {
	// noop
}

func (l *OutputEditor) Set(data any) {
	v, ok := data.(OutputEditorData)
	if ok {
		l.data = v
	}
}

func (l *OutputEditor) Get() any {
	return l.data
}

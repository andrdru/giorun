package widgets

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	runWidget "github.com/andrdru/giorun/widget"
)

type (
	// HelloEditor input view
	HelloEditor struct {
		w *widget.Editor

		handler func()

		data HelloEditorData
	}

	HelloEditorData struct {
		Value string
	}
)

const (
	HelloEditorID = "hello_editor"
)

var _ runWidget.Item = &HelloEditor{}

func NewHelloEditor(handler func()) *HelloEditor {
	return &HelloEditor{
		w: &widget.Editor{
			SingleLine: true,
		},
		handler: handler,
	}
}

func (h *HelloEditor) ID() string {
	return HelloEditorID
}

func (h *HelloEditor) Widget(th *material.Theme) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		e := material.Editor(th, h.w, "print anything")
		e.Font.Style = font.Italic

		border := widget.Border{Color: color.NRGBA{A: 0xFF}, CornerRadius: unit.Dp(4), Width: unit.Dp(2)}

		return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(8)).Layout(gtx, e.Layout)
		})
	}
}

func (h *HelloEditor) Handler() {
	for _, e := range h.w.Events() {
		switch e.(type) {
		case widget.ChangeEvent:
			h.handler()
		}
	}
}

func (h *HelloEditor) Set(data any) {
	v, ok := data.(HelloEditorData)
	if ok {
		h.data = v

		h.w.SetText(v.Value)
	}
}

func (h *HelloEditor) Get() any {
	h.data.Value = h.w.Text()

	return h.data
}

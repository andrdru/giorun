package widget

import (
	"sync"

	"gioui.org/widget"
)

type (
	Editor struct {
		mu *sync.Mutex

		editor *widget.Editor
	}
)

func NewEditor(e *widget.Editor) *Editor {
	return &Editor{
		mu:     &sync.Mutex{},
		editor: e,
	}
}

func (e *Editor) SetText(s string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.editor.SetText(s)
}

func (e *Editor) Editor() *widget.Editor {
	return e.editor
}

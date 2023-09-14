package widget

import (
	"sync"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type (
	//Item widget item interface
	Item interface {
		// ID should be unique
		ID() string
		// Widget format gio widget
		Widget(th *material.Theme) layout.Widget
		// Handler handle gio widget events
		Handler()
		// Set widget data
		Set(any)
		// Get widget data
		Get() any
	}

	//List of Items
	List struct {
		mu sync.Mutex

		data map[string]Item
	}
)

func NewList(list []Item) *List {
	return &List{
		mu: sync.Mutex{},

		data: sliceToMapString(list, func(v Item) string { return v.ID() }),
	}
}

// Set item data by id
// handle with Item.Set
func (l *List) Set(id string, data any) (ok bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	v, ok := l.data[id]
	if !ok {
		return false
	}

	v.Set(data)

	return true
}

// Get item data by id
// handle with Item.Get
func (l *List) Get(id string) (data any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	v, ok := l.data[id]
	if !ok {
		return nil
	}

	return v.Get()
}

func sliceToMapString[T any](v []T, keyFunc func(v T) string) (m map[string]T) {
	if len(v) == 0 {
		return nil
	}

	m = make(map[string]T, len(v))
	for _, el := range v {
		m[keyFunc(el)] = el
	}

	return m
}

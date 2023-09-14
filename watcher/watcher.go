package watcher

import (
	"context"
	"sync"
	"time"
)

type (
	Watcher struct {
		mu *sync.Mutex

		updateFunc func()
		frequency  time.Duration
		triggerred bool
	}
)

func NewWatcher(updateFunc func(), frequency time.Duration) *Watcher {
	return &Watcher{
		mu:         &sync.Mutex{},
		updateFunc: updateFunc,
		frequency:  frequency,
	}
}

func (w *Watcher) Watch(ctx context.Context) {
	ticker := time.NewTicker(w.frequency)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
		}

		if !w.isTriggered() {
			continue
		}

		w.updateFunc()
		w.Trigger(false)
	}
}

func (w *Watcher) Trigger(value bool) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.triggerred = value
}

func (w *Watcher) isTriggered() bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.triggerred
}

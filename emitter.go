package events

import "sync"

type Emitter interface {
	// Subscribe an event listener
	On(name string, listener Listener) Emitter

	// Unsubscribe an event by name
	Off(name string) Emitter

	// Fires an event
	Emit(event Event) error
}

func NewEmitter() Emitter {
	return &factoryEmitter{
		listeners: make(map[string][]Listener),
	}
}

type factoryEmitter struct {
	listeners map[string][]Listener
}

func (e *factoryEmitter) On(name string, listener Listener) Emitter {
	e.listeners[name] = append(e.listeners[name], listener)
	return e
}

func (e *factoryEmitter) Off(name string) Emitter {
	e.listeners[name] = make([]Listener, 0)
	return e
}

func (e *factoryEmitter) Emit(event Event) error {
	listeners, ok := e.listeners[event.Name()]
	if !ok {
		return nil
	}

	if event.IsAsync() {
		return e.runAsync(event, listeners)
	} else {
		return e.runSync(event, listeners)
	}
}

func (e *factoryEmitter) runAsync(event Event, listeners []Listener) error {
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	for _, listener := range listeners {
		wg.Add(1)
		go func(listener Listener) {
			defer wg.Done()
			mutex.Lock()
			_ = listener(event)
			mutex.Unlock()
		}(listener)
	}
	wg.Wait()
	return nil
}

func (e *factoryEmitter) runSync(event Event, listeners []Listener) error {
	for _, listener := range listeners {
		err := listener(event)
		if err != nil {
			return err
		}
	}

	return nil
}

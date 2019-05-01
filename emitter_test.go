package events

import (
	"fmt"
	"gitlab.com/gobeer/errors"
	"testing"
)

func TestNewEmitter(t *testing.T) {
	e := NewEmitter()
	if e == nil {
		t.Error("Emitter is nil")
	}
}

func TestFactoryEmitter_On(t *testing.T) {
	e := NewEmitter()
	e.On("some_name", func(event Event) errors.Error {
		return nil
	})
}

func TestFactoryEmitter_Off(t *testing.T) {
	e := NewEmitter()
	e.Off("some_name")
}

func TestFactoryEmitter_Emit_SyncEvent(t *testing.T) {
	n := "some_event"
	e := NewEmitter()
	s := "some_message"
	c := 0
	e.On(n, func(event Event) errors.Error {
		if event.Payload() != s {
			t.Errorf("Expects payload is %s", s)
		}
		c = 10
		return nil
	})
	e.On(n, func(event Event) errors.Error {
		c = 5
		return nil
	})
	e.Emit(NewEvent(n, s))
	if c != 5 {
		t.Errorf("Expects c = %d", 5)
	}
}

func TestFactoryEmitter_Emit_AsyncEvent(t *testing.T) {
	n := "some_event"
	e := NewEmitter()
	s := "some_message"
	c := 0
	e.On(n, func(event Event) errors.Error {
		c = 10
		return nil
	})
	e.On(n, func(event Event) errors.Error {
		c = 5
		return nil
	})
	e.Emit(NewAsyncEvent(n, s))
	if c != 5 && c != 10 {
		t.Errorf("Expects c = %d OR c = %d", 5, 10)
	} else {
		fmt.Printf("c = %d\n", c)
	}
}
package events

import (
	"errors"
	"fmt"
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
	e.On("some_name", func(event Event) error {
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
	e.On(n, func(event Event) error {
		if event.Payload() != s {
			t.Errorf("Expects payload is %s", s)
		}
		c = 10
		return nil
	})
	e.On(n, func(event Event) error {
		c = 5
		return nil
	})
	err := e.Emit(NewEvent(n, s))
	if err != nil {
		t.Errorf("Expect err is nil. Got %s", err)
	}
	if c != 5 {
		t.Errorf("Expects c = %d", 5)
	}
}

func TestFactoryEmitter_Emit_SyncEvent_ReturnError(t *testing.T) {
	n := "some_event"
	e := NewEmitter()
	s := "some_message"
	c := 0
	e.On(n, func(event Event) error {
		return errors.New("something goes wrong")
	})
	e.On(n, func(event Event) error {
		c = 5
		return nil
	})
	err := e.Emit(NewEvent(n, s))
	if err == nil {
		t.Error("Expect err is is not nil")
	}
	if c != 0 {
		t.Error("Expects c = 0")
	}
}

func TestFactoryEmitter_Emit_AsyncEvent(t *testing.T) {
	n := "some_event"
	e := NewEmitter()
	s := "some_message"
	c := 0
	e.On(n, func(event Event) error {
		c = 10
		return nil
	})
	e.On(n, func(event Event) error {
		c = 5
		return nil
	})
	err := e.Emit(NewAsyncEvent(n, s))
	if err != nil {
		t.Errorf("Expect err is nil. Got %s", err)
	}
	if c != 5 && c != 10 {
		t.Errorf("Expects c = %d OR c = %d", 5, 10)
	} else {
		fmt.Printf("c = %d\n", c)
	}
}

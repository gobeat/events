package events

import "testing"

func TestNewEvent(t *testing.T) {
	n := "some_name"
	p := 5
	e := NewEvent(n, p)
	if e.Name() != n || e.Payload() != p || e.IsAsync() {
		t.Error("Invalid event is created")
	}
}

func TestNewAsyncEvent(t *testing.T) {
	n := "some_name"
	p := 5
	e := NewAsyncEvent(n, p)
	if e.Name() != n || e.Payload() != p || !e.IsAsync() {
		t.Error("Invalid event is created")
	}
}

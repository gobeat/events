package events

type Listener func(event Event) error

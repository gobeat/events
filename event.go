package events

type Event interface {
	Name() string
	Payload() interface{}
	IsAsync() bool
}

func NewEvent(name string, payload interface{}) Event {
	return &factoryEvent{
		name:    name,
		payload: payload,
		isAsync: false,
	}
}

func NewAsyncEvent(name string, payload interface{}) Event {
	return &factoryEvent{
		name:    name,
		payload: payload,
		isAsync: true,
	}
}

type factoryEvent struct {
	name    string
	payload interface{}
	isAsync bool
}

func (e *factoryEvent) Name() string {
	return e.name
}

func (e *factoryEvent) Payload() interface{} {
	return e.payload
}

func (e *factoryEvent) IsAsync() bool {
	return e.isAsync
}

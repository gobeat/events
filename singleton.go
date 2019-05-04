package events

var (
	e Emitter
)

func EmitterInstance() Emitter {
	if e == nil {
		e = NewEmitter()
	}

	return e
}

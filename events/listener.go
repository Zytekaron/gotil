package events

type Listener struct {
	id        uint64
	event     string
	predicate interface{}
	handler   interface{}
}

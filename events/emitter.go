package events

import (
	"reflect"
	"sync"
	"sync/atomic"
)

type EventEmitter struct {
	listeners map[string]map[uint64]*Listener
	globals   map[uint64]*Listener

	listenerMutex sync.Mutex
	globalMutex   sync.Mutex
}

type Listener struct {
	id        uint64
	event     string
	predicate interface{}
	handler   interface{}
}

// All is a way to specify "all events"
const All = "*"

var idCounter uint64

// New creates a new EventEmitter
func New() *EventEmitter {
	return &EventEmitter{
		listeners: make(map[string]map[uint64]*Listener),
		globals:   make(map[uint64]*Listener),
	}
}

// On creates an event handler and returns
// a function that can be called to delete it
//
// Panics when the event or handler is omitted
//
// Use All to capture all events. Catch-all handlers must accept
// ...interface{} OR match the parameters of all other handlers
//
// When Emit is called, if the arguments do not match, it will panic
func (e *EventEmitter) On(event string, handler interface{}) func() {
	return e.on(event, nil, handler)
}

// OnConditional is equivalent to On, but with a predicate
func (e *EventEmitter) OnConditional(event string, predicate, handler interface{}) func() {
	return e.on(event, predicate, handler)
}

func (e *EventEmitter) on(event string, predicate, handler interface{}) func() {
	if event == "" {
		panic("event name is empty")
	}
	if handler == nil {
		panic("handler func is nil")
	}

	listener := &Listener{
		id:        atomic.AddUint64(&idCounter, 1),
		event:     event,
		predicate: predicate,
		handler:   handler,
	}

	if event == All {
		e.globalMutex.Lock()
		e.globals[listener.id] = listener
		e.globalMutex.Unlock()
		return func() {
			e.globalMutex.Lock()
			delete(e.globals, listener.id)
			e.globalMutex.Unlock()
		}
	}

	// the rest of the f1 uses the listener map
	e.listenerMutex.Lock()
	defer e.listenerMutex.Unlock()

	if _, ok := e.listeners[event]; !ok {
		e.listeners[event] = make(map[uint64]*Listener)
	}
	e.listeners[event][listener.id] = listener

	return func() {
		e.listenerMutex.Lock()
		defer e.listenerMutex.Unlock()

		delete(e.listeners[event], listener.id)

		// if there are no more listeners, remove the map
		if len(e.listeners[event]) == 0 {
			delete(e.listeners, event)
		}
	}
}

// Dispatch dispatches an event
//
// All non-nil return values will be sent to the
// returned channel, then it will be closed
func (e *EventEmitter) Dispatch(event string, args ...interface{}) <-chan []interface{} {
	ch := make(chan []interface{})

	callArgs := make([]reflect.Value, len(args))
	for i, a := range args {
		callArgs[i] = reflect.ValueOf(a)
	}

	var wg sync.WaitGroup

	e.listenerMutex.Lock()
	listeners, ok := e.listeners[event]
	if ok {
		for _, listener := range listeners {
			wg.Add(1)
			go func(l *Listener) {
				callHandler(l, callArgs, ch)
				wg.Done()
			}(listener)
		}
	}
	e.listenerMutex.Unlock()

	e.globalMutex.Lock()
	for _, listener := range e.globals {
		wg.Add(1)
		go func(l *Listener) {
			callHandler(l, callArgs, ch)
			wg.Done()
		}(listener)
	}
	e.globalMutex.Unlock()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// Emit emits an event and ignores the handler results
func (e *EventEmitter) Emit(event string, args ...interface{}) {
	callArgs := make([]reflect.Value, len(args))
	for i, a := range args {
		callArgs[i] = reflect.ValueOf(a)
	}

	e.listenerMutex.Lock()
	listeners, ok := e.listeners[event]
	if ok {
		for _, listener := range listeners {
			go callHandler(listener, callArgs, nil)
		}
	}
	e.listenerMutex.Unlock()

	e.globalMutex.Lock()
	for _, listener := range e.globals {
		go callHandler(listener, callArgs, nil)
	}
	e.globalMutex.Unlock()
}

func callHandler(listener *Listener, args []reflect.Value, ch chan<- []interface{}) {
	if listener.predicate != nil {
		res := reflect.ValueOf(listener.predicate).Call(args)
		if !res[0].Interface().(bool) {
			return
		}
	}

	values := reflect.ValueOf(listener.handler).Call(args)

	if ch != nil {
		res := make([]interface{}, 0)
		for _, value := range values {
			res = append(res, value.Interface())
		}
		if res != nil {
			ch <- res
		}
	}
}

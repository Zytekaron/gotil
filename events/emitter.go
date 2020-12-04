package events

import (
	"github.com/zytekaron/gotil/random"
	"reflect"
	"sync"
)

type EventEmitter struct {
	listeners map[string]map[uint64]Listener
	globals   map[uint64]Listener
}

type Listener struct {
	id        uint64
	event     string
	predicate interface{}
	handler   interface{}
}

// Another way to specify "all events"
const (
	All = "*"
)

// Create a new EventEmitter
func New() *EventEmitter {
	return &EventEmitter{
		listeners: make(map[string]map[uint64]Listener),
		globals:   make(map[uint64]Listener),
	}
}

// Create an event handler.
//
// Panics when the event or handler is omitted.
//
// Provide a func(?) bool before the handler to include a predicate.
//
// Cancel the event by calling the returned function.
//
// Use "*" to capture all events. Catch-all handlers must accept
// ...interface{} OR match the parameters of all other handlers.
// When Emit is called, if the arguments do not match, it will panic.
func (e *EventEmitter) On(event string, functions ...interface{}) func() {
	if event == "" {
		panic("no event name specified")
	}
	if len(functions) == 0 {
		panic("no handler provided")
	}

	var predicate, handler interface{}
	if len(functions) == 1 {
		predicate = nil
		handler = functions[0]
	}
	if len(functions) == 2 {
		predicate = functions[0]
		handler = functions[1]
	}

	listener := Listener{
		id:        random.Uint64(),
		event:     event,
		predicate: predicate,
		handler:   handler,
	}

	if event == "*" {
		e.globals[listener.id] = listener
		return func() {
			delete(e.globals, listener.id)
		}
	}

	if _, ok := e.listeners[event]; !ok {
		e.listeners[event] = make(map[uint64]Listener, 0)
	}

	e.listeners[event][listener.id] = listener
	return func() {
		delete(e.listeners[event], listener.id)

		// if there are no more listeners, remove the map
		if len(e.listeners[event]) == 0 {
			delete(e.listeners, event)
		}
	}
}

// Emit an event.
//
// All non-nil return values will be sent to the
// returned channel, then it will be closed.
func (e *EventEmitter) Emit(event string, args ...interface{}) <-chan []interface{} {
	ch := make(chan []interface{})

	callArgs := make([]reflect.Value, len(args))
	for i, a := range args {
		callArgs[i] = reflect.ValueOf(a)
	}

	var wg sync.WaitGroup

	listeners, ok := e.listeners[event]
	if ok {
		for _, listener := range listeners {
			wg.Add(1)
			go callHandler(&wg, &listener, callArgs, ch)
		}
	}

	for _, listener := range e.globals {
		wg.Add(1)
		go callHandler(&wg, &listener, callArgs, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func call(function interface{}, args []reflect.Value) []reflect.Value {
	handler := reflect.ValueOf(function)
	return handler.Call(args)
}

func callHandler(wg *sync.WaitGroup, listener *Listener, args []reflect.Value, ch chan<- []interface{}) {
	if listener.predicate != nil {
		res := call(listener.predicate, args)
		if !res[0].Interface().(bool) {
			return
		}
	}

	values := call(listener.handler, args)

	res := make([]interface{}, 0)
	for _, value := range values {
		res = append(res, value.Interface())
	}
	if res != nil {
		ch <- res
	}
	wg.Done()
}

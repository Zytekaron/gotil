package events

import (
	"github.com/zytekaron/gotil/random"
	"reflect"
	"sync"
)

type EventEmitter struct {
	listeners map[string]map[uint64]Listener
}

// Create a new EventEmitter
func New() *EventEmitter {
	return &EventEmitter{
		listeners: make(map[string]map[uint64]Listener),
	}
}

// Create an event handler and get a canceller function
//
// Please note that the return value of this method
// is subject to change in the near future (to Listener)
func (e *EventEmitter) On(event string, functions ...interface{}) func() {
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
	if _, ok := e.listeners[event]; !ok {
		e.listeners[event] = make(map[uint64]Listener, 0)
	}

	e.listeners[event][listener.id] = listener
	return func() {
		delete(e.listeners[event], listener.id) // remove the listener

		if len(e.listeners[event]) == 0 { // if there are no more listeners
			delete(e.listeners, event) // remove the map of listeners
		}
	}
}

// Emit an event
func (e *EventEmitter) Emit(event string, args ...interface{}) <-chan interface{} {
	listeners, ok := e.listeners[event]
	if !ok {
		return nil
	}

	ch := make(chan interface{})

	callArgs := make([]reflect.Value, len(args))
	for i, a := range args {
		callArgs[i] = reflect.ValueOf(a)
	}

	var wg sync.WaitGroup
	for _, listener := range listeners {
		wg.Add(1)

		go func(listener Listener) {
			if listener.predicate != nil {
				predicate := reflect.ValueOf(listener.predicate)
				result := predicate.Call(callArgs)
				if !result[0].Interface().(bool) {
					wg.Done()
					return
				}
			}

			handler := reflect.ValueOf(listener.handler)
			res := make([]interface{}, 0)
			for _, value := range handler.Call(callArgs) {
				res = append(res, value.Interface())
			}
			ch <- res
			wg.Done()
		}(listener)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

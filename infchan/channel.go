package infchan

import (
	"github.com/zyedidia/generic/list"
)

// New creates a new infinite channel with the given type.
//
// The write operation does not block. and the read operation
// will only block when the internal buffer is empty.
//
// The internal buffer uses a linked list to store elements, so
// the time complexity of insertion and deletion is always O(1).
//
// Closing the input channel will effectively wait until the last
// element in the buffer has been read, then close the output channel.
func New[T any]() (in chan<- T, out <-chan T) {
	input := make(chan T)
	return input, Wrap(input)
}

// Wrap wraps an existing channel into a new infinitely buffered channel.
//
// The write operation does not block. and the read operation
// will only block when the internal buffer is empty.
//
// The internal buffer uses a linked list to store elements, so
// the time complexity of insertion and deletion is always O(1).
//
// Closing the input channel will effectively wait until the last
// element in the buffer has been read, then close the output channel.
func Wrap[T any](in <-chan T) (out <-chan T) {
	output := make(chan T)

	go func() {
		queue := list.New[T]()

		// read until in channel is closed and queue is empty
	loop:
		for in != nil || queue.Front != nil {
			// if the queue is empty, only read from the input channel
			if queue.Front == nil {
				val, ok := <-in
				if !ok {
					// this will not block due to nil channel
					// since the channel must be non-nil or
					// the queue must be non-empty to loop
					in = nil
					break
				}
				queue.PushBack(val)
			}

			// the queue is guaranteed to not be empty
			select {
			case val, ok := <-in:
				if !ok {
					// prevent further reads; further <-in reads
					// will block, allowing only the next case
					// to run. this will deplete the queue, then
					// lead to the channel being closed below.
					in = nil
					continue loop
				}
				queue.PushBack(val)
			case output <- queue.Front.Value:
				queue.Remove(queue.Front)
			}
		}

		close(output)
	}()

	return output
}

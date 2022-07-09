package infchan

import "testing"

func TestInfiniteChanCount(t *testing.T) {
	in, out := New[int]()

	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in) // depletes buffer, then closes out

	count := 0
	for range out {
		count++
	}
	if count != 10 {
		t.Error("expected 10 elements in the output channel, got", count)
	}
}

func TestInfiniteChanValues(t *testing.T) {
	in, out := New[int]()

	in <- 1
	if val := <-out; val != 1 {
		t.Error("expected result 1 from the output channel, got", val)
	}

	in <- 2
	in <- 3
	if val := <-out; val != 2 {
		t.Error("expected result 2 from the output channel, got", val)
	}
	if val := <-out; val != 3 {
		t.Error("expected result 3 from the output channel, got", val)
	}
}

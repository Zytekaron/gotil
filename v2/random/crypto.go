package random

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

type cryptoRand struct{}

var SecureRng = rand.New(&cryptoRand{})

// Seed panics when called
func (s *cryptoRand) Seed(int64) {
	panic("cannot seed crypto random source")
}

func (s *cryptoRand) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s *cryptoRand) Uint64() (i uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &i)
	if err != nil {
		panic(err)
	}
	return i
}

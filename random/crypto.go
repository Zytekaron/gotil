package random

import (
	cr "crypto/rand"
	"encoding/binary"
	"math/rand"
)

type cryptoRand struct{}

var (
	cryptoSource = cryptoRand{}
	SecureRng    = rand.New(&cryptoSource)
)

// Cannot seed crypto random source: Panics when called
func (s *cryptoRand) Seed(int64) {
	panic("cannot seed crypto random source")
}

func (s *cryptoRand) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s *cryptoRand) Uint64() (uint uint64) {
	err := binary.Read(cr.Reader, binary.BigEndian, &uint)
	if err != nil {
		panic(err)
	}
	return uint
}

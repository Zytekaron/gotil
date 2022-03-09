package random

import (
	crand "crypto/rand"
	"encoding/binary"
	"math"
)

type ChoiceValue interface {
	~rune | ~byte
}

// SecureSlice generates a random slice of a given length and value choice set
func SecureSlice[T ChoiceValue](length int, chars []T) ([]T, error) {
	// maximum valid value for modular division
	// to maintain a perfectly even distribution
	maxValid := (math.MaxUint16/len(chars))*len(chars) - 1

	// the resulting runes
	result := make([]T, length)
	index := 0
	// until the string is full
	for index < length {
		// the amount of entropy to generate for each pass
		entropyLength := int(math.Ceil(float64(length-index)*1.1) * 2)

		// generate entropy
		entropy := make([]byte, entropyLength)
		_, err := crand.Read(entropy)
		if err != nil {
			return nil, err
		}

		i := 0
		for i < entropyLength && index < length {
			// read 2 bytes
			bytes := entropy[i : i+2]
			value := int(binary.BigEndian.Uint16(bytes))
			i += 2
			// ignore values that would create an uneven distribution
			if value <= maxValid {
				// safe distribution for modular division
				result[index] = T(chars[value%len(chars)])
				index++
			}
		}
	}

	return result, nil
}

// MustSecureSlice generates a random slice a given length and
// value choice set and ignore errors caused by the random source
func MustSecureSlice[T ChoiceValue](length int, chars []T) []T {
	res, err := SecureSlice(length, chars)
	if err != nil {
		panic(err)
	}
	return res
}

// SecureString generates a random string of a given length and value choice set
func SecureString[T ~string](length int, chars T) (string, error) {
	res, err := SecureSlice(length, []rune(chars))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// MustSecureString generates a random string of a given length and
// value choice set and ignore errors caused by the random source
func MustSecureString[T ~string](length int, chars T) string {
	str, err := SecureString(length, chars)
	if err != nil {
		panic(err)
	}
	return str
}

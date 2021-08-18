package random

import (
	crand "crypto/rand"
	"encoding/binary"
	"math"
)

// SecureRunes generates a random rune slice of a given length and character set
func SecureRunes(length int, chars []rune) ([]rune, error) {
	// maximum valid value for modular division
	// to maintain a perfectly even distribution
	maxValid := (math.MaxUint16/len(chars))*len(chars) - 1

	// the resulting runes
	result := make([]rune, 0)
	// until the string is full
	for len(result) < length {
		// the amount of entropy to generate for each pass
		entropyLength := int(math.Ceil(float64(length-len(result))*1.1) * 2)

		// generate entropy
		entropy := make([]byte, entropyLength)
		_, err := crand.Read(entropy)
		if err != nil {
			return nil, err
		}

		// the entropy index
		i := 0
		for i < entropyLength && len(result) < length {
			// read 2 bytes
			bytes := entropy[i : i+2]
			value := int(binary.BigEndian.Uint16(bytes))
			i += 2
			// ignore values that would create an uneven distribution
			if value <= maxValid {
				// safe distribution for modular division
				char := chars[value%len(chars)]
				result = append(result, char)
			}
		}
	}

	return result, nil
}

// MustSecureRunes generates a random rune slice a given length
// and character set and ignore errors caused by the random source
func MustSecureRunes(length int, chars []rune) []rune {
	res, err := SecureRunes(length, chars)
	if err != nil {
		panic(err)
	}
	return res
}

// SecureString generates a random string of a given length and character set
func SecureString(length int, chars []rune) (string, error) {
	res, err := SecureRunes(length, chars)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// MustSecureString generates a random string of a given length
// and character set and ignore errors caused by the random source
func MustSecureString(length int, chars []rune) string {
	str, err := SecureString(length, chars)
	if err != nil {
		panic(err)
	}
	return str
}

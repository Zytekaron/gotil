package security

import (
	"github.com/zytekaron/gotil/random"
)

// A One Time Pad implementation in Go
type OneTimePad struct{}

const byteMax = 256

// Create a new OneTimePad
func NewOneTimePad() *OneTimePad {
	return &OneTimePad{}
}

// Encode a message using a single-use key.
// If a key is not provided, one will be generated and returned.
// You may append to the message and it will not affect the result.
//
// Security Tips:
// Do not use the same key twice, ever!
// Do not save the key once you have enciphered a message!
func (o *OneTimePad) Encode(message, key []byte) []byte {
	output := make([]byte, len(message))

	for i := 0; i < len(message); i++ {
		keyIndex := int(key[i])
		msgIndex := int(message[i])
		output[i] = byte((keyIndex + msgIndex) % byteMax)
	}

	return output
}

// Encode a message using a single-use key.
// If a key is not provided, one will be generated and returned.
// You may append to the message and it will not affect the result./
//
// Security Tips:
// Do not use the same key twice, ever!
// Do not save the key once you have enciphered a message!
func (o *OneTimePad) EncodeString(message, key string) string {
	return string(o.Encode([]byte(message), []byte(key)))
}

// Decode a message using a single-use key.
// The message may be longer than the key.
//
// Security Tips:
// Do not use the same key twice, ever!
// Do not save the key once you have enciphered a message!
func (o *OneTimePad) Decode(message, key []byte) []byte {
	output := make([]byte, len(message))

	for i := 0; i < len(message); i++ {
		letterIndex := int(message[i])
		keyIndex := int(key[i])

		output[i] = byte((letterIndex - keyIndex + byteMax) % byteMax)
	}

	return output
}

// Decode a message using a single-use key.
// The message may be longer than the key.
//
// Security Tips:
// Do not use the same key twice, ever!
// Do not save the key once you have enciphered a message!
func (o *OneTimePad) DecodeString(message, key string) string {
	return string(o.Decode([]byte(message), []byte(key)))
}

// Generate a key with a given length
//
// Security Tips:
// Do not use the same key twice, ever!
// Do not save the key once you have enciphered a message!
func (o *OneTimePad) GenerateKey(length int) []byte {
	bytes := make([]byte, length)
	random.SecureRng.Read(bytes)
	return bytes
}

package security

import (
	"github.com/zytekaron/gotil/v2/random"
)

// OneTimePad is an implementation of the
// One-Time Pad information security algorithm.
//
// Important security considerations: A key should
// never be used more than once, and it should be
// securely disposed of once it has been used.
type OneTimePad struct{}

const byteMax = 256

// NewOneTimePad creates a new OneTimePad.
func NewOneTimePad() *OneTimePad {
	return &OneTimePad{}
}

// Encode encodes a message using a single-use key.
//
// The length of the key must be greater than or equal to the length of the message.
func (o *OneTimePad) Encode(message, key []byte) []byte {
	output := make([]byte, len(message))

	for i := 0; i < len(message); i++ {
		msgIndex := int(message[i])
		keyIndex := int(key[i])
		output[i] = byte((msgIndex + keyIndex) % byteMax)
	}

	return output
}

// EncodeString calls Encode with the message and key bytes.
func (o *OneTimePad) EncodeString(message, key string) string {
	return string(o.Encode([]byte(message), []byte(key)))
}

// Decode decodes a message using a single-use key.
//
// The length of the key must be greater than or equal to the length of the message.
func (o *OneTimePad) Decode(message, key []byte) []byte {
	output := make([]byte, len(message))

	for i := 0; i < len(message); i++ {
		letterIndex := int(message[i])
		keyIndex := int(key[i])

		output[i] = byte((letterIndex - keyIndex + byteMax) % byteMax)
	}

	return output
}

// DecodeString calls Decode with the message and key bytes.
func (o *OneTimePad) DecodeString(message, key string) string {
	return string(o.Decode([]byte(message), []byte(key)))
}

// GenerateKey generates a key with a given length.
func (o *OneTimePad) GenerateKey(length int) []byte {
	bytes := make([]byte, length)
	random.SecureRng.Read(bytes)
	return bytes
}

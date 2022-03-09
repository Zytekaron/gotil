package security

import "github.com/zytekaron/gotil/random"

// OneTimeRunes is an implementation of the
// One-Time Pad information security algorithm
// that uses a custom alphabet instead of bytes
//
// Important security considerations: A key should
// never be used more than once, and it should be
// securely disposed of once it has been used.
type OneTimeRunes struct {
	Chars []rune
}

// NewOneTimeRunes creates a new OneTimeRunes with the specified characters
func NewOneTimeRunes(chars []rune) *OneTimeRunes {
	return &OneTimeRunes{
		Chars: chars,
	}
}

// NewDefaultRunes creates a new OneTimeRunes with the default characters
//
// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789:;,.!?
func NewDefaultRunes() *OneTimeRunes {
	return &OneTimeRunes{Chars: []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789:;,.!? ")}
}

// Encode encodes a message using a single-use key.
//
// The length of the key must be greater than or equal to the length of the message
func (o *OneTimeRunes) Encode(message, key []rune) []rune {
	output := make([]rune, len(message))

	for i := 0; i < len(message); i++ {
		letterIndex := index(o.Chars, message[i])
		keyIndex := index(o.Chars, key[i])

		out := (letterIndex + keyIndex) % len(o.Chars)
		output[i] = o.Chars[out]
	}

	return output
}

// EncodeString calls Encode with the message and key characters
func (o *OneTimeRunes) EncodeString(message, key string) string {
	return string(o.Encode([]rune(message), []rune(key)))
}

// Decode decodes a message using a single-use key.
//
// The length of the key must be greater than or equal to the length of the message
func (o *OneTimeRunes) Decode(message, key []rune) []rune {
	output := make([]rune, len(message))

	for i := 0; i < len(message); i++ {
		letterIndex := index(o.Chars, message[i])
		keyIndex := index(o.Chars, key[i])

		length := len(o.Chars)
		out := (letterIndex - keyIndex + length) % length
		output[i] = o.Chars[out]
	}

	return output
}

// DecodeString calls Decode with the message and key characters
func (o *OneTimeRunes) DecodeString(message, key string) string {
	return string(o.Decode([]rune(message), []rune(key)))
}

func index(charset []rune, char rune) int {
	for i, c := range charset {
		if c == char {
			return i
		}
	}
	return -1
}

// GenerateKey generates a key with a given length
func (o *OneTimeRunes) GenerateKey(length int) []rune {
	return random.MustSecureSlice(length, o.Chars)
}

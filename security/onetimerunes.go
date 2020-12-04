package security

import "github.com/zytekaron/gotil/random"

type OneTimeRunes struct {
	Chars []rune
}

// Create a new OneTimeRunes with the specified characters
func NewRunes(chars []rune) *OneTimeRunes {
	return &OneTimeRunes{Chars: chars}
}

// Create a new OneTimeRunes with the default characters
// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz:;,.!? 0123456789
func NewDefaultRunes() *OneTimeRunes {
	return &OneTimeRunes{Chars: []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz:;,.!? 0123456789")}
}

// Encode a message using a single-use key.
// If a key is not provided, one will be generated and returned.
// You may append to the message and it will not affect the result.
//
// Security Tips:
// Do not use the same key twice, ever!
// Do not save the key once you have enciphered a message!
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

// Encode a message using a single-use key.
// If a key is not provided, one will be generated and returned.
// You may append to the message and it will not affect the result./
//
// Security Tips:
// Do not use the same key twice, ever!
// Do not save the key once you have enciphered a message!
func (o *OneTimeRunes) EncodeString(message, key string) string {
	return string(o.Encode([]rune(message), []rune(key)))
}

// Decode a message using a single-use key.
// The message may be longer than the key.
//
// Security Tips:
// Do not use the same key twice, ever!
// Do not save the key once you have enciphered a message!
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

// Decode a message using a single-use key.
// The message may be longer than the key.
//
// Security Tips:
// Do not use the same key twice, ever!
// Do not save the key once you have enciphered a message!
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

// Generate a key with a given length
//
// Security Tips:
// Do not use the same key twice, ever!
// Do not save the key once you have enciphered a message!
func (o *OneTimeRunes) GenerateKey(length int) []rune {
	return random.MustSecureRunes(length, o.Chars)
}

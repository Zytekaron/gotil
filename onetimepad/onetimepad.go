package onetimepad

import "github.com/zytekaron/gotil/random"

type OneTimePad struct {
	Chars []rune
}

func New(chars []rune) *OneTimePad {
	return &OneTimePad{Chars: chars}
}

func NewDefault() *OneTimePad {
	return &OneTimePad{Chars: []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz;,.!? 0123456789")}
}

// Encode a message using a single-use key.
// If a key is not provided, one will be generated and returned.
// You may append to the message and it will not affect the result.
//
// Security Tips:
// Do not use the same key twice, ever!
// Do not save the key once you have enciphered a message!
func (o *OneTimePad) Encode(message, key []rune) []rune {
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
func (o *OneTimePad) EncodeString(message, key string) string {
	return string(o.Encode([]rune(message), []rune(key)))
}

// Decode a message using a single-use key.
// The message may be longer than the key.
//
// Security Tips:
// Do not use the same key twice, ever!
// Do not save the key once you have enciphered a message!
func (o *OneTimePad) Decode(message, key []rune) []rune {
	output := make([]rune, len(message))

	for i := 0; i < len(message); i++ {
		letterIndex := index(o.Chars, message[i])
		keyIndex := index(o.Chars, key[i])

		length := len(o.Chars)
		out := (length + letterIndex - keyIndex) % length
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
func (o *OneTimePad) DecodeString(message, key string) string {
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
func (o *OneTimePad) GenerateKey(length int) []rune {
	return random.MustSecureRunes(length, o.Chars)
}
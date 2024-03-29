package security

import (
	"testing"
)

func TestKeyGeneration(t *testing.T) {
	otp := NewOneTimePad()
	key := otp.GenerateKey(32)

	if len(key) != 32 {
		t.Error("expected key length of 32, got", len(key))
	}
}

func TestMessageIntegrity(t *testing.T) {
	otp := NewOneTimePad()

	message := "Hello there! This is a secret message that must be sent secretly."
	key := otp.GenerateKey(len(message))

	encoded := otp.Encode([]byte(message), key)
	decoded := otp.Decode(encoded, key)

	if string(decoded) != message {
		t.Error("encoded and decoded text are not the same\n\tmessage >", message, "\n\tdecoded >", string(decoded))
	}
}

package security

import (
	. "github.com/zytekaron/gotil/v2/security"
	"testing"
)

var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789,.-!;?$\"'\\ ")

func TestRunesKeyGeneration(t *testing.T) {
	otp := NewDefaultRunes()
	key := otp.GenerateKey(32)

	if len(key) != 32 {
		t.Error("expected key length of 32, got", len(key))
	}
}

func TestRunesMessageIntegrity(t *testing.T) {
	otp := NewOneTimeRunes(chars)

	message := "Hello there! This is a secret message that must be sent secretly."
	key := otp.GenerateKey(len(message))

	encoded := otp.Encode([]rune(message), key)
	decoded := otp.Decode(encoded, key)

	if string(decoded) != message {
		t.Error("encoded and decoded text are not the same\n\tmessage >", message, "\n\tdecoded >", string(decoded))
	}
}

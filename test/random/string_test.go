package random

import (
	. "github.com/zytekaron/gotil/random"
	"testing"
)

func TestSecureRunes(t *testing.T) {
	runes, err := SecureRunes(32, []rune("0123456789abcdef"))

	if err != nil {
		t.Error(err)
	}

	if len(runes) != 32 {
		t.Error("Expected []rune of length 32, got", len(runes))
	}
}

func TestSecureString(t *testing.T) {
	str, err := SecureString(32, []rune("0123456789abcdef"))

	if err != nil {
		t.Error(err)
	}

	if len(str) != 32 {
		t.Error("Expected []rune of length 32, got", len(str))
	}
}

func TestMustSecureRunes(t *testing.T) {
	MustSecureRunes(32, []rune("0123456789abcdef"))
}

func TestMustSecureString(t *testing.T) {
	MustSecureString(32, []rune("0123456789abcdef"))
}

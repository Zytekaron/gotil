package random

import "testing"

func TestSecureRunes(t *testing.T) {
	runes, err := SecureRunes(32, []rune("0123456789abcdef"))

	if err != nil {
		t.Error(err)
	}

	if len(runes) != 32 {
		t.Error("Expected []rune length of 32, got", len(runes))
	}
}

func TestSecureString(t *testing.T) {
	str, err := SecureString(32, []rune("0123456789abcdef"))

	if err != nil {
		t.Error(err)
	}

	if len(str) != 32 {
		t.Error("Expected []rune length of 32, got", len(str))
	}
}
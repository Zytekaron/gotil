package random

import (
	"testing"
)

func TestSecureRunes(t *testing.T) {
	runes, err := SecureSlice(32, []rune("0123456789abcdef"))
	if err != nil {
		t.Error(err)
	}
	if len(runes) != 32 {
		t.Error("expected result length to be 32, got", len(runes))
	}
}

func TestSecureString(t *testing.T) {
	str, err := SecureString(32, "0123456789abcdef")
	if err != nil {
		t.Error(err)
	}
	if len(str) != 32 {
		t.Error("expected result length to be 32, got", len(str))
	}
}

func TestMustSecureRunes(*testing.T) {
	MustSecureSlice(32, []rune("0123456789abcdef"))
}

func TestMustSecureString(*testing.T) {
	MustSecureString(32, "0123456789abcdef")
}

// testing for any conflicts in the distribution behavior
// that don't appear when using strings and rune slices
func TestByteSlice(t *testing.T) {
	bytes := []byte{0x00, 0x01, 0x1d, 0x22, 0x45, 0x6c, 0x77, 0x82, 0x95, 0x9f, 0xa4, 0xc7, 0xd9, 0xe7, 0xfe, 0xff}
	set := map[byte]bool{}
	for _, b := range bytes {
		set[b] = true
	}

	res := MustSecureSlice(256, bytes)
	for _, b := range res {
		if !set[b] {
			t.Errorf("found unexpected byte 0x%.2x\n", b)
		}
	}
}

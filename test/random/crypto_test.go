package random

import (
	. "github.com/zytekaron/gotil/random"
	"testing"
)

func TestCrypto(t *testing.T) {
	SecureRng.Float64()
}

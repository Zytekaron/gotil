package random

import (
	. "github.com/zytekaron/gotil/random"
	"testing"
)

func TestCrypto(*testing.T) {
	SecureRng.Float64()
}

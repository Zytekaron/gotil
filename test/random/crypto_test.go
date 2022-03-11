package random

import (
	. "github.com/zytekaron/gotil/v2/random"
	"testing"
)

func TestCrypto(*testing.T) {
	SecureRng.Float64()
}

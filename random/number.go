package random

// Generates a cryptographically secure random number
func Uint64() uint64 {
	return cryptoSource.Uint64()
}

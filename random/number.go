package random

// Uint64 generates a cryptographically secure random number
func Uint64() uint64 {
	return SecureRng.Uint64()
}

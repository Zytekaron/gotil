# gotil

Version 2.2.1

## Release Notes

### Version 2.2

Added new features and fixed/updated most documentation.
Some minor bug fixes and performance improvements.

- new features:
  - created the `infchan` package with functions allowing the creation of infinitely buffered channels
  - added synchronous rate limiting types to the `rl` package, with internal mutexing.
  - added `random.NewRandomizerRNG`, allowing for a custom RNG (`*rand.Rand`) to be passed
- bug fixes:
  - fixed an issue where `rl.Bucket#DrawMax` could draw a negative amount, effectively removing uses / increasing remaining buckets, if `ForceDraw` previously overdrew
- other changes:
  - tests now create files in the OS temp directory (`os.TempDir()`)
  - modified some internal implementations
    - improved speed of `security.OneTimeRunes` by using a map instead of indexing
    - minimally improved performance of some `rl.Bucket` operations
  - complete documentation refactoring
    - added documentation where it was missing
    - formatted existing documentation to central styles
    - fixed errors in a couple places
    - added a deprecation comment for `optional.Optional[T]#GetOrZero` in favor of `OrElseZero`

### Version 2.1

- new features:
  - added json marshalling and unmarshalling for type `optional.Optional[T]`
  - added `GetOrZero() T` method to type `optional.Optional[T]` for naming consistency

### Version 2.0

Updated library to Go 1.18, migrated to and created new generic types and functions.

- new features:
  - java-style optional type `optional.Optional[T]`
  - generic slice fuzzy function `fn.FuzzySlice[T comparable](needle, haystack []T) bool`
- converted functions and types to 1.18 generics:
  - `random.Randomizer[T]{}`
  - `slices.Sum[T Integer | Float](s []T) T`
  - `slices.Average[T Float](s []T) T` -- still uses moving average to help prevent overflow
  - `random.SecureRunes` -> `random.SecureSlice[T ~rune | ~byte](length int, set []T) ([]T, error)`

## License

**gotil** is licensed under the [MIT License](./LICENSE)

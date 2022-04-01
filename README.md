# gotil

Version 2.0.0

## Release Notes


### Version 2.1.0

- new features:
  - added json marshalling and unmarshalling for type `optional.Optional[T]`
  - added `GetOrZero() T` method to type `optional.Optional[T]` for naming consistency

### Version 2.0.0

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

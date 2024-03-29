package random

import (
	"errors"
	"github.com/zytekaron/gotil/v2/maths"
	"math/rand"
	"time"
)

var defaultRand = rand.NewSource(time.Now().UnixNano())

// Randomizer is an implementation of the alias method, designed
// to efficiently sample from a discrete probability distribution.
//
// Results are added along with their weight (likelihood of being
// drawn among other elements), and when prepared, they may be drawn.
//
// Time complexity of preparation: O(nlogn)
//
// Time complexity of sampling: O(1)
type Randomizer[T any] struct {
	rng       *rand.Rand
	weights   []float64
	results   []T
	sampleInt func() int
	ready     bool
}

// RandomizerElement is an element of a Randomizer.
type RandomizerElement[T any] struct {
	Weight float64
	Result T
}

// NewRandomizerRNG creates a Randomizer instance that uses a custom random number generator.
func NewRandomizerRNG[T any](rng *rand.Rand) *Randomizer[T] {
	return &Randomizer[T]{
		rng:     rng,
		weights: make([]float64, 0),
		results: make([]T, 0),
	}
}

// NewRandomizer creates a Randomizer instance using a simple pseudo-random number generator.
func NewRandomizer[T any]() *Randomizer[T] {
	return NewRandomizerRNG[T](rand.New(defaultRand))
}

// NewSecureRandomizer creates a Randomizer instance that uses a cryptographically secure random number generator.
func NewSecureRandomizer[T any]() *Randomizer[T] {
	return NewRandomizerRNG[T](SecureRng)
}

// Add adds elements to the Randomizer.
func (r *Randomizer[T]) Add(weight float64, result T) error {
	if r.ready {
		return errors.New("randomizer has already been prepared and is now immutable")
	}
	r.weights = append(r.weights, weight)
	r.results = append(r.results, result)
	return nil
}

// AddElement adds elements to the Randomizer.
func (r *Randomizer[T]) AddElement(element *RandomizerElement[T]) error {
	return r.Add(element.Weight, element.Result)
}

// AddMany adds multiple elements to the Randomizer from an array of elements.
func (r *Randomizer[T]) AddMany(elements []RandomizerElement[T]) error {
	if r.ready {
		return errors.New("randomizer has already been prepared and is now immutable")
	}
	for _, e := range elements {
		r.weights = append(r.weights, e.Weight)
		r.results = append(r.results, e.Result)
	}
	return nil
}

// Prepare the randomizer for sampling
func (r *Randomizer[T]) Prepare() {
	r.sampleInt = aliasMethod(r.weights, r.rng)
	r.ready = true
}

// Sample samples the alias method to get a random value.
func (r *Randomizer[T]) Sample() (T, error) {
	if !r.ready {
		var t T
		return t, errors.New("randomizer 'prepare' method must be called before sampling may begin")
	}
	return r.sample(), nil
}

// MustSample samples the alias method to get a random value.
func (r *Randomizer[T]) MustSample() T {
	res, err := r.Sample()
	if err != nil {
		panic(err)
	}
	return res
}

// SampleMany samples the randomizer to get multiple random values.
func (r *Randomizer[T]) SampleMany(count int) ([]T, error) {
	if !r.ready {
		return nil, errors.New("randomizer 'prepare' method must be called before sampling may begin")
	}
	results := make([]T, count)
	for i := range results {
		results[i] = r.sample()
	}
	return results, nil
}

// MustSampleMany samples the randomizer to get multiple random values.
func (r *Randomizer[T]) MustSampleMany(count int) []T {
	res, err := r.SampleMany(count)
	if err != nil {
		panic(err)
	}
	return res
}

func (r *Randomizer[T]) sample() T {
	return r.results[r.sampleInt()]
}

func aliasMethod(probabilities []float64, rng *rand.Rand) func() int {
	sum := maths.Sum(probabilities)

	probMultiplier := float64(len(probabilities)) / sum
	for i := range probabilities {
		probabilities[i] *= probMultiplier
	}

	overFull := make([]int, 0)
	underFull := make([]int, 0)
	for i, p := range probabilities {
		if p > 1 {
			overFull = append(overFull, i)
		} else if p < 1 {
			underFull = append(underFull, i)
		}
	}

	aliases := make(map[int]int)
	for len(overFull)+len(underFull) > 0 {
		if len(overFull) > 0 && len(underFull) > 0 {
			firstOver := overFull[0]
			firstUnder := underFull[0]

			probabilities[firstOver] += probabilities[firstUnder] - 1
			aliases[firstUnder] = firstOver

			underFull = underFull[1:]
			if probabilities[firstOver] > 1 {
				e := overFull[0]
				overFull = append(overFull, e)
			} else if probabilities[firstOver] < 1 {
				e := overFull[0]
				underFull = append(underFull, e)
			}
			overFull = overFull[1:]
		} else {
			if len(overFull) > 0 {
				for i := range overFull {
					probabilities[i] = 1
				}
				overFull = make([]int, 0)
			} else {
				for i := range underFull {
					probabilities[i] = 1
				}
				underFull = make([]int, 0)
			}
		}
	}

	return func() int {
		index := rng.Intn(len(probabilities))
		if rng.Float64() < probabilities[index] {
			return index
		} else {
			return aliases[index]
		}
	}
}

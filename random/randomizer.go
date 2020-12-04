package random

import (
	"errors"
	"math/rand"
	"time"
)

var defaultRand = rand.NewSource(time.Now().UnixNano())

// A randomizer implementation using the alias method
type Randomizer struct {
	secure    bool
	weights   []float64
	results   []interface{}
	sampleInt func() int
	ready     bool
}

type RandomizerElement struct {
	Weight float64
	Result interface{}
}

// Create a Randomizer instance
func NewRandomizer() *Randomizer {
	return &Randomizer{
		secure:  false,
		weights: make([]float64, 0),
		results: make([]interface{}, 0),
	}
}

// Create a Randomizer instance that uses
// a cryptographically secure source for
// random number and sample generation
func NewSecureRandomizer() *Randomizer {
	return &Randomizer{
		secure:  true,
		weights: make([]float64, 0),
		results: make([]interface{}, 0),
	}
}

// Add elements to the Randomizer
func (r *Randomizer) Add(weight float64, result interface{}) error {
	if r.ready {
		return errors.New("randomizer has already been prepared and is now immutable")
	}
	r.weights = append(r.weights, weight)
	r.results = append(r.results, result)
	return nil
}

// Add elements to the Randomizer
func (r *Randomizer) AddElement(element *RandomizerElement) error {
	return r.Add(element.Weight, element.Result)
}

// Add elements to the Randomizer from an array of arrays
func (r *Randomizer) AddMany(elements []RandomizerElement) error {
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
func (r *Randomizer) Prepare() {
	var rng *rand.Rand
	if r.secure {
		rng = SecureRng
	} else {
		rng = rand.New(defaultRand)
	}
	r.sampleInt = aliasMethod(r.weights, rng)
	r.ready = true
}

// Sample the alias method to get a random value respecting the weights of each element
func (r *Randomizer) Sample() (interface{}, error) {
	if !r.ready {
		return nil, errors.New("randomizer 'prepare' method must be called before sampling may begin")
	}
	return r.sample(), nil
}

// Sample the alias method to get a random value respecting the weights of each element
func (r *Randomizer) MustSample() interface{} {
	res, err := r.Sample()
	if err != nil {
		panic(err)
	}
	return res
}

// Sample the alias method to get random values respecting the weights of each element
func (r *Randomizer) SampleMany(count int) ([]interface{}, error) {
	if !r.ready {
		return nil, errors.New("randomizer 'prepare' method must be called before sampling may begin")
	}
	results := make([]interface{}, count)
	for i := range results {
		results[i] = r.sample()
	}
	return results, nil
}

// Sample the alias method to get random values respecting the weights of each element
func (r *Randomizer) MustSampleMany(count int) []interface{} {
	res, err := r.SampleMany(count)
	if err != nil {
		panic(err)
	}
	return res
}

func (r *Randomizer) sample() interface{} {
	return r.results[r.sampleInt()]
}

func aliasMethod(probabilities []float64, rng *rand.Rand) func() int {
	sum := float64sum(probabilities)

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

func float64sum(slice []float64) float64 {
	total := 0.0
	for _, e := range slice {
		total += e
	}
	return total
}

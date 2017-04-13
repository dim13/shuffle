// Package shuffle implements Fisher-Yates shuffle algorithm
package shuffle

import "math/rand"

type Interface interface {
	Len() int
	Swap(i, j int)
}

var rnd = rand.New(cryptoSrc{})

func Shuffle(v Interface) {
	for i := v.Len() - 1; i > 0; i-- {
		v.Swap(i, rnd.Intn(i+1))
	}
}

// IntSlice implements shuffle Interface for slice of ints
type IntSlice []int

func (v IntSlice) Len() int      { return len(v) }
func (v IntSlice) Swap(i, j int) { v[i], v[j] = v[j], v[i] }

// StringSlice implements shuffle Interface for slice of strings
type StringSlice []string

func (v StringSlice) Len() int      { return len(v) }
func (v StringSlice) Swap(i, j int) { v[i], v[j] = v[j], v[i] }

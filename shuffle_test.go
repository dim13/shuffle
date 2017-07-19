package shuffle

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

type fixedSrc struct{}

func (fixedSrc) Int63() int64 { return 4 } // chosen by fair dice roll
func (fixedSrc) Seed(_ int64) {}

func init() {
	rnd = rand.New(fixedSrc{}) // override default rand source
}

func TestShuffleInt(t *testing.T) {
	v := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	Shuffle(IntSlice(v))
	if !reflect.DeepEqual(v, r) {
		t.Errorf("got %v, want %v", v, r)
	}
}

func TestShuffleString(t *testing.T) {
	v := []string{"A", "B", "C"}
	r := []string{"B", "C", "A"}
	Shuffle(StringSlice(v))
	if !reflect.DeepEqual(v, r) {
		t.Errorf("got %v, want %v", v, r)
	}
}

func TestShuffleQuick(t *testing.T) {
	f := func(v IntSlice) bool {
		r := IntSlice{}
		if len(v) > 0 {
			r = append(v[1:], v[0])
		}
		Shuffle(v)
		return reflect.DeepEqual(v, r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

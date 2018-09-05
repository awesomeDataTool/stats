package stats

import (
	"math"
	"math/rand"
	"testing"
)

// makeFloatSlice makes a slice of float64s
func makeFloatSlice(c int) []float64 {
	lf := make([]float64, 0, c)
	for i := 0; i < c; i++ {
		f := float64(i * 100)
		lf = append(lf, f)
	}
	return lf
}

func makeRandFloatSlice(c int) []float64 {
	lf := make([]float64, 0, c)
	rand.Seed(unixnano())
	for i := 0; i < c; i++ {
		f := float64(i * 100)
		lf = append(lf, f)
	}
	return lf
}

func TestFloat64ToInt(t *testing.T) {
	m := float64ToInt(234.0234)
	if m != 234 {
		t.Errorf("%x != %x", m, 234)
	}
	m = float64ToInt(-234.0234)
	if m != -234 {
		t.Errorf("%x != %x", m, -234)
	}
	m = float64ToInt(1)
	if m != 1 {
		t.Errorf("%x != %x", m, 1)
	}
}

func TestCounts(t *testing.T) {
	s := []float64{1, 1, 2, 3, 4, 4, 4, 4, math.Inf(1)}
	m := counts(s)
	if m[1] != 2 {
		t.Errorf("%x != %x", m[1], 2)
	}
	if m[2] != 1 {
		t.Errorf("%x != %x", m[2], 1)
	}
	if m[4] != 4 {
		t.Errorf("%x != %x", m[4], 4)
	}
	if m[math.Inf(1)] != 1 {
		t.Errorf("%x != %x", m[1], 1)
	}
}

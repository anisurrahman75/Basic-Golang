package math

import (
	"fmt"
	"testing"
)

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func Average(slc []float64) float64 {
	var avg float64
	for _, i := range slc {
		avg += i
	}
	avg = (avg) / float64(len(slc))
	return avg
}

// for testing test file name as: "fileName_test" ="test" keyword must
func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			fmt.Println("sunny")
			t.Error(
				"\n\n",
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

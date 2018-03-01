package main

import "testing"

func TestSumHalfWay(t *testing.T) {
	tests := []struct {
		Get  string
		Want int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	for _, test := range tests {
		sum := sumHalfWay(test.Get)
		if sum != test.Want {
			t.Errorf("%d did not equal what was expected: %d", sum, test.Want)
		}
	}
}

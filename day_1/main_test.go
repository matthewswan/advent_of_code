package main

import (
	"testing"
)

func TestConvertStringToInt(t *testing.T) {
	tests := []struct {
		Get  string
		Want int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
		{"0", 0},
	}

	for _, test := range tests {
		sum := sumNext(test.Get)
		if sum != test.Want {
			t.Errorf("%d did not equal what was expected: %d", sum, test.Want)
		}
	}
}

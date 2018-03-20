package main

import "testing"

func TestGetCheckSum(t *testing.T) {
	tests := []struct {
		Get  []int
		Want int
	}{
		{[]int{5, 1, 9, 5}, 8},
		{[]int{7, 5, 3}, 4},
		{[]int{2, 4, 6, 8}, 6},
	}

	for _, test := range tests {
		checkSum := getCheckSum(test.Get)
		if checkSum != test.Want {
			t.Errorf("%d did not equal what was expected: %d", checkSum, test.Want)
		}
	}
}

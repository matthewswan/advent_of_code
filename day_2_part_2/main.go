package main

import "fmt"

func main() {
	var sum int
	input := [][]int{
		{5, 9, 2, 8},
		{9, 4, 7, 3},
		{3, 8, 6, 5},
	}
	for _, s := range input {
		sum += findDivisibleValuesSum(s)
	}
	fmt.Printf("Sum: %d\n", sum)
}

func findDivisibleValuesSum(s []int) int {
	var rowSum int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] != s[j] && s[i]%s[j] == 0 {
				rowSum += s[i] / s[j]
			}
		}
	}
	return rowSum
}

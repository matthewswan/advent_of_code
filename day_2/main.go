package main

import "fmt"

func main() {
	var sum int
	input := [][]int{
		{5, 1, 9, 5},
		{7, 5, 3},
		{2, 4, 6, 8},
	}
	for _, s := range input {
		sum += getCheckSum(s)
	}
	fmt.Printf("Sum: %d\n", sum)
}

func getCheckSum(s []int) int {
	minInt, maxInt := s[0], s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < minInt {
			minInt = s[i]
		}
		if s[i] > maxInt {
			maxInt = s[i]
		}
	}
	return maxInt - minInt
}

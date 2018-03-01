package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	numString := "91212129"
	total := sumNext(numString)
	fmt.Printf("Sum: %d\n", total)
}

func sumNext(str string) int {
	var sum int
	lastInt := convertStringToInt(string(str[(len(str) - 1)]))
	firstInt := convertStringToInt(string(str[0]))
	if firstInt == lastInt {
		sum = firstInt
	}
	for i := 0; i < len(str)-1; i++ {
		currentInt := convertStringToInt(string(str[i]))
		nextInt := convertStringToInt(string(str[i+1]))
		if currentInt == nextInt {
			sum += currentInt
		}
	}
	return sum
}

func convertStringToInt(str string) int {
	returnInt, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("strconv returned an error")
	}
	return returnInt
}

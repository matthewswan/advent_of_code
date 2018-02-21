package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	numString := "91212129"
	var sum int
	lastInt := convertStringToInt(string(numString[(len(numString) - 1)]))
	for i := 0; i < len(numString)-1; i++ {
		currentInt := convertStringToInt(string(numString[i]))
		nextInt := convertStringToInt(string(numString[i+1]))
		if i == 0 && currentInt == lastInt {
			sum = currentInt
		} else if currentInt == nextInt {
			sum += currentInt
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func convertStringToInt(str string) int {
	returnInt, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("strconv returned an error")
	}
	return returnInt
}

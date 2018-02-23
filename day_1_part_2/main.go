package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var sum int
	numString := "123123"
	subString1, subString2 := numString[:(len(numString)/2)], numString[(len(numString)/2):]
	for i, _ := range subString1 {
		if num1, num2 := convertStringToInt(string(subString1[i])), convertStringToInt(string(subString2[i])); num1 == num2 {
			sum += num1 + num2
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

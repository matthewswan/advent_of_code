package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	numString := "123123"
	total := sumHalfWay(numString)
	fmt.Printf("Sum: %d\n", total)
}

func sumHalfWay(str string) int {
	var sum int
	subString1, subString2 := str[:(len(str)/2)], str[(len(str)/2):]
	for i, _ := range subString1 {
		if num1, num2 := convertStringToInt(string(subString1[i])), convertStringToInt(string(subString2[i])); num1 == num2 {
			sum += num1 + num2
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

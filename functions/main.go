package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	tranformationFn := getTansformFunction(&numbers)
	doubleNumbers := transformNumbers(&numbers, tranformationFn)
	fmt.Println(doubleNumbers)
	tripledNumbers := transformNumbers(&numbers, tripleNumber)
	fmt.Println(tripledNumbers)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	transformedNumbers := []int{}
	for _, value := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(value))
	}
	return transformedNumbers
}

func doubleNumber(number int) int {
	return number * 2
}

func tripleNumber(number int) int {
	return number * 3
}

// function that returns a function
func getTansformFunction(numbers *[]int) func(int) int {
	if (*numbers)[0] == 1 {
		return doubleNumber
	} else {
		return tripleNumber
	}
}

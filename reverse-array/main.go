package main

import "fmt"

func main() {
	var input = []string{"a", "b", "c", "d"}

	newArray := reverseArrayWithNewFixSizeArray(input)

	fmt.Printf("reverseArrayWithFixArray: %v \n", newArray)

	newArray = reverseArrayWithNewDynamicSizeArray(input)

	fmt.Printf("reverseArrayWithDynamicArray: %v \n", newArray)

	fmt.Printf("reverseArray FROM: %v \n", input)
	reverseArray(input)
	fmt.Printf("reverseArray TO: %v \n", input)
}

func reverseArrayWithNewFixSizeArray(input []string) []string {

	newArray := make([]string, len(input))
	for i := 0; i < len(input); i++ {
		newArray[i] = input[len(input)-1-i]
	}

	return newArray
}

func reverseArrayWithNewDynamicSizeArray(input []string) []string {

	newArray := []string{}
	for i := len(input) - 1; i >= 0; i-- {
		newArray = append(newArray, input[i])
	}

	return newArray
}

func reverseArray(input []string) {
	for i := 0; i < len(input)/2; i++ {
		input[i], input[len(input)-1-i] = input[len(input)-1-i], input[i]
	}
}

package main

import "fmt"

var m map[int]int

func main() {
	m = map[int]int{}

	value := 42
	fmt.Println(recursiveFibonacci(value))
	fmt.Println(iterativeFibonacci(value))
	fmt.Println(recursiveFibonacciMemoization(value))
}

func recursiveFibonacciMemoization(num int) int {
	if num == 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	if _, found := m[num]; found {
		return m[num]
	}

	m[num] = recursiveFibonacciMemoization(num-1) + recursiveFibonacciMemoization(num-2)
	return m[num]
}

func recursiveFibonacci(num int) int {
	if num == 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	return recursiveFibonacci(num-1) + recursiveFibonacci(num-2)
}

func iterativeFibonacci(num int) int {

	if num == 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	total := 0
	curr := 1
	for i := 0; i < num; i++ {
		last := total
		total = total + curr
		curr = last
	}

	return total
}

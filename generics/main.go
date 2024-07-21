package main

import "fmt"

func main() {
	a, b := 1, 2
	result := add(a, b)
	fmt.Printf("Add %v and %v is equal to: %v\n", a, b, result)

	result = mult(a, b)
	fmt.Printf("Multiply %v and %v is equal to: %v\n", a, b, result)

}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func mult[T int | float64](a, b T) T {
	return a * b
}

package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	var float32Slice = []float32{1.0, 2.0, 3.0}
	fmt.Println(sumAllTypes(intSlice))
	fmt.Println(sumAllTypes(float32Slice)) // Multiple types work
	fmt.Println(isEmpty(intSlice))
	fmt.Println(isEmpty(float32Slice))
}

// T is a generic, which lets us define multiple types as parameter
func sumAllTypes[T int | float32](slice []T) T {
	var sum T = 0
	for _, num := range slice {
		sum += num
	}
	return sum
}

// Can also use any but be careful of what the function is doing when using any
// Ex: can't do += if the type is bool, it will throw an error
func isEmpty[T any](slice []T) bool{
	return len(slice) == 0
}
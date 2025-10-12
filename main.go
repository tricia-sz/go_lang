package main

import "fmt"

func main() {
	slice1 := []int{5, 1, 2, 3}
	slice2 := []string{"A", "B", "C", "D"}
	newInts := reverse((slice1))
	newString := reverse((slice2))
	fmt.Println(newInts)
	fmt.Println(newString)
}

func reverse[T int | string](slice []T) []T {
	newInts := make([]T, len(slice))

	newItensLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newInts[newItensLen] = slice[i]
		newItensLen--
	}

	return newInts

}

package main

import (
	"fmt"
)

func main() {
	// Example slice for testing
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// Test FindMax
	max := FindMax(numbers)
	fmt.Printf("Maximum value: %d\n", max)

	// Test RemoveDuplicates
	unique := RemoveDuplicates(numbers)
	fmt.Printf("After removing duplicates: %v\n", unique)

	// Test ReverseSlice
	reversed := ReverseSlice(numbers)
	fmt.Printf("Reversed: %v\n", reversed)

	// Test FilterEven
	evenOnly := FilterEven(numbers)
	fmt.Printf("Even numbers only: %v\n", evenOnly)
}

// FindMax returns the maximum value in a slice of integers.
// If the slice is empty, it returns 0.
func FindMax(numbers []int) int {
	
	if len(numbers) == 0 {
    	return 0
	}
	
	max := numbers[0]
	
	for i := 1; i < len(numbers); i++ {
	    if numbers[i] > max {
	        max = numbers[i]
	    }
	}
	return max
}

// RemoveDuplicates returns a new slice with duplicate values removed,
// preserving the original order of elements.
func RemoveDuplicates(numbers []int) []int {
	result := make([]int, 0, len(numbers))
	hashMap := make(map[int]struct{})
	for _, number := range numbers {
	    if _, ok := hashMap[number]; !ok {
	        result = append(result, number)
	        hashMap[number] = struct{}{}
	    }
	}
	return result
}

// ReverseSlice returns a new slice with elements in reverse order.
func ReverseSlice(slice []int) []int {
	result := make([]int, len(slice))
	
	for i := 0; i < len(slice); i++ {
	    result[i] = slice[len(slice) - i - 1]
	}
	return result
}

// FilterEven returns a new slice containing only the even numbers
// from the original slice.
func FilterEven(numbers []int) []int {
	result := make([]int, 0, len(numbers))
	for _, number := range numbers {
	    if number % 2 == 0 {
	        result = append(result, number)
	    }
	}
	return result
}

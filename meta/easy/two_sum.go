package main

import (
	"fmt"
)

// twoSum finds two numbers in the array that add up to the target
// Returns the indices of the two numbers such that they add up to target
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, exists := seen[complement]; exists {
			return []int{j, i}
		}

		seen[num] = i
	}

	return []int{}
}

// Add a main function to demonstrate the twoSum function
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("Input: nums = %v, target = %d\n", nums, target)
	fmt.Printf("Output: %v\n", result)

	// Add more examples if desired
	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	fmt.Printf("\nInput: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("Output: %v\n", result2)
}

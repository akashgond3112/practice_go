package easy

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

func main() {
	// Function to run examples
	runExample := func(nums []int, target int) {
		fmt.Printf("Input: nums = %v, target = %d\n", nums, target)

		result := twoSum(nums, target)

		if len(result) == 2 {
			fmt.Printf("Output: %v\n", result)
			fmt.Printf("Explanation: nums[%d] + nums[%d] = %d + %d = %d\n",
				result[0], result[1], nums[result[0]], nums[result[1]], target)
		} else {
			fmt.Println("No solution found.")
		}
		fmt.Println()
	}

	// Example 1
	runExample([]int{2, 7, 11, 15}, 9)

	// Example 2
	runExample([]int{3, 2, 4}, 6)

	// Example 3
	runExample([]int{3, 3}, 6)

	// Example 4
	runExample([]int{1, 2, 3, 4, 5}, 9)

	// Custom example - change these values to test your own inputs
	customNums := []int{8, 1, 5, 2, 4}
	customTarget := 6
	runExample(customNums, customTarget)

}

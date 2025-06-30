package main

import "fmt"

// customSort reorders characters in string s according to the specified order
// Characters in 'order' appear first (maintaining their frequency from s)
// Remaining characters appear after in their original frequency
func customSort(s string, order string) string {
	// Count frequency of each character in input string
	count := make(map[rune]int)
	for _, char := range s {
		count[char]++
	}

	// Pre-allocate result slice with capacity equal to input length
	result := make([]rune, 0, len(s))

	// Add characters according to specified order, respecting their frequency
	for _, char := range order {
		if freq, exists := count[char]; exists {
			for i := 0; i < freq; i++ {
				result = append(result, char)
			}
			delete(count, char) // Remove processed character
		}
	}

	// Add remaining characters that weren't in the order specification
	for char, freq := range count {
		for i := 0; i < freq; i++ {
			result = append(result, char)
		}
	}

	return string(result)
}

// Example usage and test cases
func main() {
	// Test case 1: Basic reordering
	s1 := "cba"
	order1 := "abc"
	fmt.Printf("Input: %s, Order: %s, Result: %s\n", s1, order1, customSort(s1, order1))
	// Expected: "abc"

	// Test case 2: With repeated characters
	s2 := "cbaabc"
	order2 := "abc"
	fmt.Printf("Input: %s, Order: %s, Result: %s\n", s2, order2, customSort(s2, order2))
	// Expected: "aaabbc"

	// Test case 3: Order doesn't contain all characters
	s3 := "helloworld"
	order3 := "dlrow"
	fmt.Printf("Input: %s, Order: %s, Result: %s\n", s3, order3, customSort(s3, order3))
	// Expected: "dllrrwoohee"

	// Test case 4: Empty strings
	s4 := ""
	order4 := "abc"
	fmt.Printf("Input: %s, Order: %s, Result: %s\n", s4, order4, customSort(s4, order4))
	// Expected: ""

	// Test case 5: Unicode characters
	s5 := "世界Hello"
	order5 := "He"
	fmt.Printf("Input: %s, Order: %s, Result: %s\n", s5, order5, customSort(s5, order5))
	// Expected: "Hello世界"
}

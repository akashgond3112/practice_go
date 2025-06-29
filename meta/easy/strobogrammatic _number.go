package main

// isStrobogrammatic checks if a number is strobogrammatic
func isStrobogrammatic(num string) bool {
	strobogrammaticPairs := map[byte]byte{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
	}

	left, right := 0, len(num)-1

	for left <= right {
		leftChar := num[left]
		rightChar := num[right]

		// Check if left character is valid
		rotated, exists := strobogrammaticPairs[leftChar]
		if !exists {
			return false
		}

		// Check if the rotated left character matches the right character
		if rotated != rightChar {
			return false
		}

		left++
		right--
	}

	return true
}

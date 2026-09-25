package twopointes

func validPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	// Compare characters from both ends.
	// If they match, move both pointers inward.
	for left < right {
		if s[left] != s[right] {

			// We are allowed to delete at most ONE character.
			//
			// At the first mismatch, only two possibilities exist:
			// 1. Delete s[left]
			// 2. Delete s[right]
			//
			// Example: "abca"
			//         ^  ^
			//         b != c
			//
			// Delete 'b' -> "aca"
			// Delete 'c' -> "aba"
			//
			// If either becomes a palindrome, the answer is true.
			return isPalindrome(s, left+1, right) ||
				isPalindrome(s, left, right-1)
		}

		left++
		right--
	}

	// No mismatch means the original string is already a palindrome.
	return true
}

func isPalindrome(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
package twopointes

func mergeAlternately(word1 string, word2 string) string {
	var result []byte

	i := 0
	j := 0

	// Continue while at least one string still has characters.
	//
	// Why "OR" instead of "AND"?
	// Example:
	// word1 = "abcd"
	// word2 = "pq"
	//
	// After "apbq", word2 is finished (j == 2),
	// but word1 still has "cd".
	// We must continue and append those remaining characters.
	for i < len(word1) || j < len(word2) {

		// Add from word1 if it still has characters.
		//
		// The condition prevents accessing word1[i]
		// when word1 has already been completely consumed.
		if i < len(word1) {
			result = append(result, word1[i])
			i++
		}

		// Add from word2 if it still has characters.
		//
		// This handles unequal-length strings.
		// Example: word1 = "abc", word2 = "pqrs"
		// After "apbqcr", only "s" remains,
		// so we append it without needing a special loop.
		if j < len(word2) {
			result = append(result, word2[j])
			j++
		}
	}

	return string(result)
}
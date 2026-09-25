package twopointes

import "sort"

func numRescueBoats(people []int, limit int) int {
	// WHY SORT?
	//
	// We need to know whether the lightest person can share
	// a boat with the heaviest person.
	//
	// After sorting:
	// left  -> lightest remaining person
	// right -> heaviest remaining person
	//
	// This lets us make a decision using two pointers.
	sort.Ints(people)

	left, right, res := 0, len(people)-1, 0

	for left <= right {

		// WHY CHECK THE LIGHTEST + HEAVIEST?
		//
		// The heaviest person MUST go into a boat.
		//
		// We want to know:
		//
		// "Can we put the lightest remaining person
		//  together with this heaviest person?"
		//
		// If even the lightest person cannot fit with the
		// heaviest person:
		//
		// people[left] + people[right] > limit
		//
		// then NOBODY else can fit with the heaviest person,
		// because everyone else is >= people[left].
		//
		// Example:
		// people = [2, 3, 5, 6]
		// limit = 7
		//
		// 2 + 6 = 8 > 7
		//
		// Since 2 is the lightest person, 3 + 6, 5 + 6
		// will also exceed 7.
		//
		// Therefore, 6 MUST go alone.
		if people[left]+people[right] <= limit {

			// The lightest and heaviest can share one boat.
			res++

			// WHY MOVE BOTH?
			//
			// Both people have now been placed into a boat,
			// so neither should be considered again.
			left++
			right--

		} else {

			// The heaviest person cannot fit even with the
			// lightest person.
			//
			// Therefore, the heaviest person has to take
			// a boat alone.
			res++

			// Only the heaviest person is removed.
			// The lightest person is still available.
			right--
		}
	}

	return res
}
package slidingwindow

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// WHY:
	// We need two positions:
	// "buy"  = the best (cheapest) day we have found to buy so far.
	// "sell" = the current day we are checking.
	//
	// We start sell at 1 because we need a day after buy.
	// We cannot buy and sell on the same day.
	maxProfit, buy, sell := 0, 0, 1

	for sell < len(prices) {

		// WHY:
		// If today's price is cheaper than our current buying price,
		// today's price is obviously a better day to buy.
		//
		// Example:
		// buy price  = 7
		// today      = 1
		//
		// There is no reason to keep 7 as our buying price because
		// buying at 1 gives us a better profit for any future selling day.
		if prices[sell] < prices[buy] {
			buy = sell
			sell++
			continue
		}

		// WHY:
		// Today's price is higher than our buying price,
		// so selling today gives us a possible profit.
		//
		// We calculate it using the cheapest buying price
		// we have found so far.
		profit := prices[sell] - prices[buy]

		// WHY:
		// We are checking every possible selling day.
		// So keep the largest profit we have seen.
		if profit > maxProfit {
			maxProfit = profit
		}

		// WHY:
		// We have finished checking today's selling price,
		// so move to the next day.
		sell++
	}

	return maxProfit
}
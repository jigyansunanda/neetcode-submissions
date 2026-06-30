func maxProfit(prices []int) int {
	minSoFar, profit := prices[0], 0

	for i:=1; i<len(prices); i++ {
		profit = max(profit, prices[i]-minSoFar)
		minSoFar = min(minSoFar, prices[i])
	}

	return profit
}

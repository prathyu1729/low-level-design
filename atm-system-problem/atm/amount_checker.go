package atm

import "sort"

type amountChecker struct {
	denominations map[float64]int
}

func (a *amountChecker) getDenominations(amount float64) (bool, map[float64]int) {

}

func canMakeAmount(amount int, denominations []int) (bool, map[int]int) {
	sort.Sort(sort.Reverse(sort.IntSlice(denominations))) // Sort denominations in descending order
	output := map[int]int{}
	dp := make([]bool, amount+1)
	dp[0] = true // Base case: 0 amount can always be made

	for _, coin := range denominations {
		for i := coin; i <= amount; i++ {
			dp[i] = dp[i] || dp[i-coin]
		}
	}

	return dp[amount], output
}

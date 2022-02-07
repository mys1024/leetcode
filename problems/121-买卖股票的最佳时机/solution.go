// 121. 买卖股票的最佳时机
// URL：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
// 难度：简单
// 关键词：动态规划、滚动数组
// 执行用时：104 ms, 在所有 Go 提交中击败了 95.29% 的用户
// 内存消耗：7.8 MB, 在所有 Go 提交中击败了 94.53% 的用户

package main

func maxProfit(prices []int) int {
	minPrice, maxProfit := prices[0], 0
	for _, curPrice := range prices {
		if curPrice < minPrice {
			minPrice = curPrice
		}
		if curPrice-minPrice > maxProfit {
			maxProfit = curPrice - minPrice
		}
	}
	return maxProfit
}

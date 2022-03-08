// 2055. 蜡烛之间的盘子
// URL：https://leetcode-cn.com/problems/plates-between-candles/
// 难度：中等
// 关键词：数组、前缀和
// 执行用时：260 ms, 在所有 Go 提交中击败了 73.08% 的用户
// 内存消耗：20 MB, 在所有 Go 提交中击败了 30.77% 的用户

package main

func platesBetweenCandles(s string, queries [][]int) []int {
	var (
		leftLatestCandle  = make([]int, len(s))
		rightLatestCandle = make([]int, len(s))
		preCandles        = make([]int, len(s))
		ans               = make([]int, len(queries))
		latestCandle      = -1
		candleCnt         = 0
	)

	for i, c := range s {
		if c == '|' {
			latestCandle = i
		}
		leftLatestCandle[i] = latestCandle
		if c == '*' {
			candleCnt++
		}
		preCandles[i] = candleCnt
	}

	latestCandle = -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '|' {
			latestCandle = i
		}
		rightLatestCandle[i] = latestCandle
	}

	for i, query := range queries {
		l, r := rightLatestCandle[query[0]], leftLatestCandle[query[1]]
		if l == -1 || r == -1 || l >= r {
			ans[i] = 0
		} else {
			ans[i] = preCandles[r] - preCandles[l]
		}
	}
	return ans
}

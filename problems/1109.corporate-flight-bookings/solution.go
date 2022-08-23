// 1109. 航班预订统计
// URL：https://leetcode-cn.com/problems/corporate-flight-bookings/
// 难度：中等
// 关键词：数组、差分
// 执行用时：108 ms, 在所有 Go 提交中击败了 89.81% 的用户
// 内存消耗：8.6 MB, 在所有 Go 提交中击败了 94.45% 的用户

package main

func corpFlightBookings(bookings [][]int, n int) []int {
	// 以差分的形式记录座位数量
	ans := make([]int, n)
	for _, booking := range bookings {
		first, last, seats := booking[0], booking[1], booking[2]
		ans[first-1] += seats
		if last < n {
			ans[last] -= seats
		}
	}
	// 将差分数据转换为实际的座位数量
	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}
	return ans
}

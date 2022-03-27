// 2028. 找出缺失的观测数据
// URL：https://leetcode-cn.com/problems/find-missing-observations/
// 难度：中等
// 关键词：数学
// 执行用时：144 ms, 在所有 Go 提交中击败了 46.15% 的用户
// 内存消耗：8.6 MB, 在所有 Go 提交中击败了 84.62% 的用户

package main

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	sn := mean * (m + n)
	for _, roll := range rolls {
		sn -= roll
	}

	expectation := float64(sn) / float64(n)
	if expectation < 1 || 6 < expectation {
		return []int(nil)
	}

	k := int(expectation)
	a := n*(k+1) - sn
	b := n - a
	ans := make([]int, 0, a+b)
	for i := 0; i < a; i++ {
		ans = append(ans, k)
	}
	for i := 0; i < b; i++ {
		ans = append(ans, k+1)
	}

	return ans
}

// 728. 自除数
// URL：https://leetcode-cn.com/problems/self-dividing-numbers/
// 难度：数学
// 关键词：字典树、哈希表、字符串
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 81.82% 的用户

package main

func IsSelfDividingNumber(num int) bool {
	tmp := num
	for tmp != 0 {
		m := tmp % 10
		if m == 0 || num%m != 0 {
			return false
		}
		tmp /= 10
	}
	return true
}

func selfDividingNumbers(left int, right int) []int {
	ans := []int{}
	for num := left; num <= right; num++ {
		if IsSelfDividingNumber(num) {
			ans = append(ans, num)
		}
	}
	return ans
}

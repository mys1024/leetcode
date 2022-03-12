// 2197. 替换数组中的非互质数
// URL：https://leetcode-cn.com/problems/replace-non-coprime-numbers-in-array/
// 难度：困难
// 关键词：数学、数论、最大公约数、最小公倍数、辗转相除法
// 执行用时：152 ms, 在所有 Go 提交中击败了 86.78% 的用户
// 内存消耗：8.8 MB, 在所有 Go 提交中击败了 81.82% 的用户

package main

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func replaceNonCoprimes(nums []int) []int {
	// 使用 nums 数组作为栈
	top := 0
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if gcd(nums[top], num) == 1 {
			top++
			nums[top] = num
			continue
		}
		for top >= 0 && gcd(nums[top], num) > 1 {
			num = lcm(nums[top], num)
			top--
		}
		top++
		nums[top] = num
	}
	return nums[:top+1]
}

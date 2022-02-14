// 735. 行星碰撞
// URL：https://leetcode-cn.com/problems/asteroid-collision/
// 难度：中等
// 关键词：栈、数组
// 执行用时：8 ms, 在所有 Go 提交中击败了 94.68% 的用户
// 内存消耗：4.3 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func IntAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func asteroidCollision(asteroids []int) []int {
	// 直接使用原数组作为栈进行操作，因为作为栈的空间
	// 不会超出原数组中已被遍历过的空间。
	head := 0
	for _, asteroid := range asteroids {
		if head == 0 || asteroids[head-1]*asteroid > 0 || asteroid > 0 {
			asteroids[head] = asteroid
			head++
			continue
		}
		for j := head - 1; j >= 0; j-- {
			if asteroids[j]*asteroid > 0 {
				asteroids[j+1] = asteroid
				head = j + 2
				break
			} else if IntAbs(asteroids[j]) > IntAbs(asteroid) {
				head = j + 1
				break
			} else if IntAbs(asteroids[j]) == IntAbs(asteroid) {
				head = j
				break
			} else if j == 0 {
				asteroids[0] = asteroid
				head = 1
			}
		}
	}
	return asteroids[:head]
}

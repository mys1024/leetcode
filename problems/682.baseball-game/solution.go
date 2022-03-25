// 682. 棒球比赛
// URL：https://leetcode-cn.com/problems/baseball-game/
// 难度：简单
// 关键词：栈
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.5 MB, 在所有 Go 提交中击败了 78.38% 的用户

package main

import "strconv"

func calPoints(ops []string) int {
	points := []int{}
	for _, o := range ops {
		n := len(points)
		if o == "+" {
			a, b := points[n-2], points[n-1]
			points = append(points, a+b)
		} else if o == "D" {
			a := points[n-1]
			points = append(points, a*2)
		} else if o == "C" {
			points = points[:n-1]
		} else {
			point, _ := strconv.Atoi(o)
			points = append(points, point)
		}
	}
	sum := 0
	for _, point := range points {
		sum += point
	}
	return sum
}

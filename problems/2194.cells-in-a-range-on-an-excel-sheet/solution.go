// 2194. Excel 表中某个范围内的单元格
// URL：https://leetcode-cn.com/problems/cells-in-a-range-on-an-excel-sheet/
// 难度：简单
// 关键词：字符串
// 执行用时：4 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：3 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func cellsInRange(s string) []string {
	col1, row1, col2, row2 := s[0], s[1], s[3], s[4]
	ans := []string{}
	cell := make([]byte, 2)
	for c := col1; c <= col2; c++ {
		for r := row1; r <= row2; r++ {
			cell[0], cell[1] = c, r
			ans = append(ans, string(cell))
		}
	}
	return ans
}

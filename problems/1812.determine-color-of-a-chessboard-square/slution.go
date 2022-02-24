// 1812. 判断国际象棋棋盘中一个格子的颜色
// URL：https://leetcode-cn.com/problems/determine-color-of-a-chessboard-square/
// 难度：简单
// 关键词：数学、字符串
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

func squareIsWhite(coordinates string) bool {
	return (coordinates[0]+coordinates[1])%2 == 1
}

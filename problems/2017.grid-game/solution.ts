// 2017. 网格游戏
// URL: https://leetcode.cn/problems/grid-game/
// 难度: 中等
// 关键词: 前缀和 数组
// 执行用时: 116 ms, 在所有 TypeScript 提交中击败了 33.33% 的用户
// 内存消耗: 57.5 MB, 在所有 TypeScript 提交中击败了 33.33% 的用户

function gridGame(grid: number[][]): number {
  const n = grid[0].length
  for (let i = 1; i < n; i++) {
    grid[0][n - 1 - i] += grid[0][n - i]
    grid[1][i] += grid[1][i - 1]
  }
  let min = Infinity
  for (let i = 0; i < n; i++) {
    let a = grid[1][i - 1]
    let b = grid[0][i + 1]
    min = Math.min(min, Math.max(a ? a : 0, b ? b : 0))
  }
  return min
}

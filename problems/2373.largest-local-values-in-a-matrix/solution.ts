// 2373. 矩阵中的局部最大值
// URL: https://leetcode.cn/problems/largest-local-values-in-a-matrix/
// 难度: 简单
// 关键词: 数组 矩阵
// 执行用时: 88 ms, 在所有 TypeScript 提交中击败了 28.95% 的用户
// 内存消耗: 45.3 MB, 在所有 TypeScript 提交中击败了 57.89% 的用户

function max(a: number, b: number): number {
  return a > b ? a : b
}

function maxOf3x3(grid: number[][], x: number, y: number) {
  let ret = grid[x][y]
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      ret = max(ret, grid[x + i][y + j])
    }
  }
  return ret
}

function largestLocal(grid: number[][]): number[][] {
  const n = grid.length
  const m = n - 2
  const ans: number[][] = new Array(m)
  for (let i = 0; i < m; i++) {
    ans[i] = new Array(m)
  }
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < m; j++) {
      ans[i][j] = maxOf3x3(grid, i, j)
    }
  }
  return ans
}

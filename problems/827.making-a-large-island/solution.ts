// 827. 最大人工岛
// URL: https://leetcode.cn/problems/making-a-large-island/
// 难度: 困难
// 关键词: 图 矩阵 深度优先搜索
// 执行用时: 212 ms, 在所有 TypeScript 提交中击败了 73.33% 的用户
// 内存消耗: 57.9 MB, 在所有 TypeScript 提交中击败了 86.67% 的用户

function largestIsland(grid: number[][]): number {
  const n = grid.length
  if (n === 0)
    return 0

  const directions: [number, number][] = [[0, 1], [0, -1], [1, 0], [-1, 0]]
  const sizes: number[] = [0, 0]

  function valid(x: number, y : number) {
    return 0 <= x && x < n && 0 <= y && y < n
  }
  
  function dfs(x: number, y : number, idx: number) {
    if (!valid(x, y) || grid[x][y] !== 1)
      return
    sizes[idx]++
    grid[x][y] = idx
    for (const [dx, dy] of directions) {
      dfs(x + dx, y + dy, idx)
    }
  }

  let idx = 2
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
      if (grid[i][j] === 1) {
        sizes[idx] = 0
        dfs(i, j, idx++)
      }
    }
  }

  if (sizes.length === 2)
    return 1

  const islands: Set<number> = new Set()
  let max = 0
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
      let size = 0
      if (grid[i][j] > 0) {
        size = sizes[grid[i][j]]
      } else {
        size = 1
        for (const [dx, dy] of directions) {
          const [x, y] = [i + dx, j + dy]
          if (!valid(x, y))
            continue
          const adjacency = grid[x][y]
          if (islands.has(adjacency))
            continue
          size += sizes[adjacency]
          islands.add(adjacency)
        }
        islands.clear()
      }
      if (size > max)
        max = size
    }
  }

  return max
}

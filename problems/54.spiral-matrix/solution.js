// 54. 螺旋矩阵
// URL：https://leetcode.cn/problems/spiral-matrix/
// 难度：中等
// 关键词：模拟、矩阵
// 执行用时：56 ms, 在所有 Go 提交中击败了 82.74% 的用户
// 内存消耗：40.9 MB, 在所有 Go 提交中击败了 33.98% 的用户

const MARK = 999

/**
 * @param {number[][]} matrix
 * @return {number[]}
 */
function spiralOrder(matrix) {
  const ans = [matrix[0][0]]
  matrix[0][0] = MARK

  let [x, y, di, flag] = [0, 0, 0, false]
  const directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]
  while (true) {
    const [nextX, nextY] = [directions[di][0]+x, directions[di][1]+y]
    if (legal(matrix, nextX, nextY)) {
      ans.push(matrix[nextX][nextY])
      matrix[nextX][nextY] = MARK
      x = nextX
      y = nextY
      flag = false
      continue
    }
    if (!flag) {
      flag = true
      di = (di+1) % 4
      continue
    }
    break
  }

  return ans
}

/**
 * @param {number[][]} matrix
 * @param {number} x
 * @param {number} y
 * @returns {boolean}
 */
function legal(matrix, x, y) {
  const m = matrix.length
  const n = matrix[0].length
  if (x < 0 || y < 0 || x >= m || y >= n) {
    return false
  }
  return matrix[x][y] !== MARK
}

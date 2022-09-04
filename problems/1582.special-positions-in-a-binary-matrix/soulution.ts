// 1582. 二进制矩阵中的特殊位置
// URL: https://leetcode.cn/problems/special-positions-in-a-binary-matrix/
// 难度: 中等
// 关键词: 数组 矩阵
// 执行用时: 68 ms, 在所有 TypeScript 提交中击败了 50.00% 的用户
// 内存消耗: 44.5 MB, 在所有 TypeScript 提交中击败了 16.67% 的用户

function numSpecial(mat: number[][]): number {
  const n = mat.length
  const m = mat[0].length

  const specialRow: boolean[] = new Array(n)
  const specialCol: boolean[] = new Array(m)
  let cnt = 0

  for (let i = 0; i < n; i++) {
    let sum = 0
    for (let j = 0; j < m; j++) {
      sum += mat[i][j]
      if (sum > 1)
        break
    }
    specialRow[i] = sum === 1 ? true : false
  }

  for (let j = 0; j < m; j++) {
    let sum = 0
    for (let i = 0; i < n; i++) {
      sum += mat[i][j]
      if (sum > 1)
        break
    }
    specialCol[j] = sum === 1 ? true : false
  }
  
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < m; j++) {
      if (mat[i][j] === 1 && specialRow[i] && specialCol[j])
        cnt++
    }
  }

  return cnt
}

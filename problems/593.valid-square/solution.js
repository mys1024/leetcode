// 593. 有效的正方形
// URL：https://leetcode.cn/problems/valid-square/
// 难度：中等
// 关键词：数学
// 执行用时：56 ms, 在所有 JavaScript 提交中击败了 92.37% 的用户
// 内存消耗：41 MB, 在所有 JavaScript 提交中击败了 94.92% 的用户

/**
 * @param {number[]} p1
 * @param {number[]} p2
 * @param {number[]} p3
 * @param {number[]} p4
 * @return {boolean}
 */
function validSquare(p1, p2, p3, p4) {
  return validDistance(p1, p2, p3, p4)
    && validDistance(p2, p1, p3, p4)
    && validDistance(p3, p1, p2, p4)
    && validDistance(p4, p1, p2, p3)
}

/**
 * @param {number[]} p1
 * @param {number[]} p2
 * @param {number[]} p3
 * @param {number[]} p4
 * @return {boolean}
 */
function validDistance(p1, p2, p3, p4) {
  const d1 = squareDistance(p1, p2)
  if (!d1)
    return false
  const d2 = squareDistance(p1, p3)
  if (!d2)
    return false
  const d3 = squareDistance(p1, p4)
  if (!d3)
    return false
  return Math.abs(Math.min(d1, d2, d3)*4 - (d1+d2+d3)) < 1e-9
}

/**
 * @param {number[]} p1
 * @param {number[]} p2
 * @return {number}
 */
function squareDistance(p1, p2) {
  return (p1[0]-p2[0])**2 + (p1[1]-p2[1])**2
}

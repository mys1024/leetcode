// 1331. 数组序号转换
// URL：https://leetcode.cn/problems/rank-transform-of-an-array/
// 难度：简单
// 关键词：数组、哈希表
// 执行用时：208 ms, 在所有 JavaScript 提交中击败了 47.79% 的用户
// 内存消耗：73.1 MB, 在所有 JavaScript 提交中击败了 9.56% 的用户

/**
 * @param {number[]} arr
 * @return {number[]}
 */
function arrayRankTransform(arr) {
  const order = {}
  let [prev, acc] = [Infinity, 1]
  for (const num of [...arr].sort((a, b) => a - b)) {
    if (num === prev) {
      continue
    }
    order[num] = acc
    prev = num
    acc++
  }
  return arr.map(n => order[n])
}

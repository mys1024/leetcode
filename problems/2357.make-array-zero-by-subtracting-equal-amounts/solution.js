// 2357. 使数组中所有元素都等于零
// URL：https://leetcode.cn/problems/make-array-zero-by-subtracting-equal-amounts/
// 难度：简单
// 关键词：数组、哈希表
// 执行用时：60 ms, 在所有 JavaScript 提交中击败了 73.73% 的用户
// 内存消耗：41.2 MB, 在所有 JavaScript 提交中击败了 43.13% 的用户

/**
 * @param {number[]} nums
 * @return {number}
 */
 function minimumOperations(nums) {
  const e = {}
  for (const num of nums) {
    if (num === 0) {
      continue
    }
    e[num] = true
  }
  return Object.keys(e).length
}

// 1217. 玩筹码
// URL：https://leetcode.cn/problems/minimum-cost-to-move-chips-to-the-same-position/
// 难度：简单
// 关键词：贪心、数组
// 执行用时：60 ms, 在所有 Go 提交中击败了 62.39% 的用户
// 内存消耗：41.1 MB, 在所有 Go 提交中击败了 32.11% 的用户

/**
 * @param {number[]} position
 * @return {number}
 */
var minCostToMoveChips = function(position) {
  let cnt0 = 0
  let cnt1 = 0
  for (const p of position) {
    if (p%2 == 0) {
      cnt0++
    } else {
      cnt1++
    }
  }
  return cnt0 < cnt1 ? cnt0 : cnt1
}

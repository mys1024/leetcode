// 1464. 数组中两元素的最大乘积
// URL：https://leetcode.cn/problems/maximum-product-of-two-elements-in-an-array/
// 难度：简单
// 关键词：数组
// 执行用时：60 ms, 在所有 TypeScript 提交中击败了 92.00% 的用户
// 内存消耗：44 MB, 在所有 TypeScript 提交中击败了 20.00% 的用户

function maxProduct(nums: number[]): number {
  let m1 = -1
  let m2 = -1
  for (const num of nums) {
    if (num > m1) {
      m2 = m1
      m1 = num
    } else if (num > m2) {
      m2 = num
    }
  }
  return (m1 - 1) * (m2 - 1)
}

// 1608. 特殊数组的特征值
// URL: https://leetcode.cn/problems/special-array-with-x-elements-greater-than-or-equal-x/
// 难度: 简单
// 关键词: 数组
// 执行用时: 60 ms, 在所有 TypeScript 提交中击败了 81.82% 的用户
// 内存消耗: 42.2 MB, 在所有 TypeScript 提交中击败了 81.82% 的用户

function specialArray(nums: number[]): number {
  const n = nums.length
  nums.sort((a, b) => a - b)
  let prev = -1
  for (let i = 0; i < n; i++) {
    const x = n - i
    const num = nums[i]
    if (num >= x && num > prev)
      return x
    prev = num
  }
  return -1
}

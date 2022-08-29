// 1470. 重新排列数组
// URL：https://leetcode.cn/problems/shuffle-the-array/
// 难度：简单
// 关键词：数组
// 执行用时：68 ms, 在所有 TypeScript 提交中击败了 87.50% 的用户
// 内存消耗：44 MB, 在所有 TypeScript 提交中击败了 17.50% 的用户

function shuffle(nums: number[], n: number): number[] {
  const ans: number[] = []
  for (let i = 0; i < n; i++) {
    ans.push(nums[i])
    ans.push(nums[n + i])
  }
  return ans
}

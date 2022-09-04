// 2367. 算术三元组的数目
// URL: https://leetcode.cn/problems/number-of-arithmetic-triplets/
// 难度: 简单
// 关键词: 数组 哈希表
// 执行用时: 72 ms, 在所有 TypeScript 提交中击败了 47.62% 的用户
// 内存消耗: 44.2 MB, 在所有 TypeScript 提交中击败了 6.35% 的用户

function arithmeticTriplets(nums: number[], diff: number): number {
  const n = nums.length
  const m: Record<number, number[]> = {}
  let cnt = 0

  for (const num of nums)
    m[num] = []

  for (let i = 0; i < n; i++) {
    const num = nums[i]
    for (let j = i + 1; j < n; j++) {
      if (nums[j] - num === diff)
        m[num].push(nums[j])
    }
  }

  for (const num of nums) {
    for (const next of m[num]) {
      cnt += m[num].length * m[next].length
    }
  }

  return cnt
}

// 1636. 按照频率将数组升序排序
// URL: https://leetcode.cn/problems/sort-array-by-increasing-frequency/
// 难度: 简单
// 关键词: 排序 数组
// 执行用时: 76 ms, 在所有 TypeScript 提交中击败了 61.54% 的用户
// 内存消耗: 44.5 MB, 在所有 TypeScript 提交中击败了 15.38% 的用户

function frequencySort(nums: number[]): number[] {
  const cnt: Record<number, number> = {}
  for (const num of nums) {
    if (!cnt[num])
      cnt[num] = 1
    else
      cnt[num]++
  }
  nums.sort((a, b) => {
    let d = cnt[a] - cnt[b]
    return d === 0 ? b - a : d
  })
  return nums
}

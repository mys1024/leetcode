// 1619. 删除某些元素后的数组均值
// URL: https://leetcode.cn/problems/mean-of-array-after-removing-some-elements/
// 难度: 简单
// 关键词: 数组
// 执行用时: 84 ms, 在所有 TypeScript 提交中击败了 9.09% 的用户
// 内存消耗: 44.1 MB, 在所有 TypeScript 提交中击败了 72.73% 的用户

function trimMean(arr: number[]): number {
  const n = arr.length
  const drop = n * 0.05
  arr.sort((a, b) => a - b)
  let s = 0
  for (let i = drop; i < n - drop; i++) {
    s += arr[i]
  }
  return s / (n - drop * 2)
}

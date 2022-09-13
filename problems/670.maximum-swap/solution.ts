// 670. 最大交换
// URL: https://leetcode.cn/problems/maximum-swap/
// 难度: 中等
// 关键词: 数组
// 执行用时: 84 ms, 在所有 TypeScript 提交中击败了 5.26% 的用户
// 内存消耗: 42.4 MB, 在所有 TypeScript 提交中击败了 47.37% 的用户

function maximumSwap(num: number): number {
  const digits = num.toString().split('').map(c => Number.parseInt(c))
  const n = digits.length
  const memo: Record<number, number[] | undefined> = {}
  for (let i = 0; i < n; i++) {
    const d = digits[i]
    if (!memo[d])
      memo[d] = new Array(9);
    (memo[d] as number[]).push(i)
  }
  for (let i = 0; i < n; i++) {
    const d = digits[i]
    for (let j = 9; j > d; j--) {
      const indexes = memo[j]
      if (!indexes || indexes.length === 0)
        continue
      const tmp = digits[i]
      const k = indexes[indexes.length - 1]
      digits[i] = digits[k]
      digits[k] = tmp
      return Number.parseInt(digits.join(''))
    }
    memo[d]?.shift()
  }
  return num
}

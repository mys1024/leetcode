// 1689. 十-二进制数的最少数目
// URL：https://leetcode.cn/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
// 难度：中等
// 关键词：字符串、贪心
// 执行用时：68 ms, 在所有 TypeScript 提交中击败了 85.71% 的用户
// 内存消耗：45.8 MB, 在所有 TypeScript 提交中击败了 85.71% 的用户

function minPartitions(n: string): number {
  for (let i = 9; i >= 0; i--) {
    if (n.includes(i.toString()))
      return i
  }
  return 0
}

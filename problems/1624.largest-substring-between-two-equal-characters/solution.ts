// 1624. 两个相同字符之间的最长子字符串
// URL: https://leetcode.cn/problems/largest-substring-between-two-equal-characters/
// 难度: 简单
// 关键词: 字符串
// 执行用时: 52 ms, 在所有 TypeScript 提交中击败了 100.00% 的用户
// 内存消耗: 42.3 MB, 在所有 TypeScript 提交中击败了 88.89% 的用户

function maxLengthBetweenEqualCharacters(s: string): number {
  const n = s.length
  const memo: Map<string, number> = new Map()

  let ans = -1
  for (let i = 0; i < n; i++) {
    const c = s[i]
    const first = memo.get(c)
    if (first === undefined) {
      memo.set(c, i)
    }
    else {
      const len = i - first - 1
      if (len > ans)
        ans = len
    }
  }

  return ans
}

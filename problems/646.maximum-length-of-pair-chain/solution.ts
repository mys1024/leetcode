// 646. 最长数对链
// URL：https://leetcode.cn/problems/maximum-length-of-pair-chain/
// 难度：中等
// 关键词：数组
// 执行用时：124 ms, 在所有 TypeScript 提交中击败了 40.00% 的用户
// 内存消耗：45.9 MB, 在所有 TypeScript 提交中击败了 80.00% 的用户

function max(a: number, b: number): number {
  return a > b ? a : b
}

function findLongestChain(pairs: number[][]): number {
  const n = pairs.length
  pairs.sort((a, b) => a[1] - b[1])

  const cnt = new Int32Array(n)
  let maxLen = 0

  for (let i = n - 1; i >= 0; i--) {
    for (let j = i + 1; j < n; j++) {
      if (pairs[i][1] < pairs[j][0]) {
        cnt[i] = max(cnt[i], cnt[j] + 1)
        maxLen = max(maxLen, cnt[i])
      }
    }
  }

  return maxLen + 1
}

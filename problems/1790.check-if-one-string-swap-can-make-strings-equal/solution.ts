// 1790. 仅执行一次字符串交换能否使两个字符串相等
// URL: https://leetcode.cn/problems/check-if-one-string-swap-can-make-strings-equal/
// 难度: 简单
// 关键词: 字符串
// 执行用时: 64 ms, 在所有 TypeScript 提交中击败了 63.33% 的用户
// 内存消耗: 44.1 MB, 在所有 TypeScript 提交中击败了 6.67% 的用户

function areAlmostEqual(s1: string, s2: string): boolean {
  const n = s1.length
  if (n !== s2.length)
    return false
  
  let flag = 0
  let [da, db] = ['', '']

  for (let i = 0; i < n; i++) {
    const [a, b] = [s1[i], s2[i]]
    if (a === b)
      continue
    switch (flag) {
      case 0:
        da = a
        db = b
        flag++
        break
      case 1:
        if (!(a === db && b === da))
          return false
        flag++
        break
      case 2:
        return false
    }
  }

  return flag !== 1
}

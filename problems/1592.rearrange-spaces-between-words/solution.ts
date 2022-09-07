// 1592. 重新排列单词间的空格
// URL: https://leetcode.cn/problems/rearrange-spaces-between-words/
// 难度: 简单
// 关键词: 字符串
// 执行用时: 64 ms, 在所有 TypeScript 提交中击败了 66.67% 的用户
// 内存消耗: 43.8 MB, 在所有 TypeScript 提交中击败了 33.33% 的用户

function reorderSpaces(text: string): string {
  let cnt = 0
  for (const c of text) {
    if (c === ' ')
      cnt++
  }
  const words = text.split(' ').filter(w => w !== '')
  if (words.length === 1) {
    return words[0] + ' '.repeat(cnt)
  } else {
    const tail = ' '.repeat(cnt % (words.length - 1))
    const sep = ' '.repeat(Math.floor(cnt / (words.length - 1)))
    return words.join(sep) + tail
  }
}

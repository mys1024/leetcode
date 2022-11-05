# 1668. 最大重复子字符串
# URL: https://leetcode.cn/problems/maximum-repeating-substring/
# 难度: 简单
# 关键词: 字符串
# 执行用时: 36 ms, 在所有 TypeScript 提交中击败了 74.17% 的用户
# 内存消耗: 15.1 MB, 在所有 TypeScript 提交中击败了 14.85% 的用户

def match(seq, i, word):
  n, m = len(seq), len(word)
  if n - i < m:
    return False
  for j in range(m):
    if seq[i + j] != word[j]:
      return False
  return True

class Solution:
  def maxRepeating(self, sequence: str, word: str) -> int:
    n = len(sequence)
    m = len(word)

    cnt = 0
    maximum = 0
    fallback = 0
    i = 0

    while i < n:
      if match(sequence, i, word):
        i += m
        cnt += 1
        maximum = max(maximum, cnt)
      else:
        i = fallback
        fallback += 1
        cnt = 0

    return maximum

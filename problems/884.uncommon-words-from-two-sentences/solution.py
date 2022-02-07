# 884. 两句话中的不常见单词
# URL：https://leetcode-cn.com/problems/uncommon-words-from-two-sentences/
# 难度：简单
# 关键词：计数
# 执行用时：32 ms, 在所有 Python3 提交中击败了 75.00% 的用户
# 内存消耗：15 MB, 在所有 Python3 提交中击败了 66.98% 的用户


from typing import List

class Solution:
  def uncommonFromSentences(self, s1: str, s2: str) -> List[str]:
    counter = {}
    words = s1.split(' ') + s2.split(' ')
    for word in words:
      if word in counter:
        counter[word] += 1
      else:
        counter.setdefault(word, 1)
    result = []
    for (word, count) in counter.items():
      if count == 1:
        result.append(word)
    return result

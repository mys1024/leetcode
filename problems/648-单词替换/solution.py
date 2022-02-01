# 648. 单词替换
# URL：https://leetcode-cn.com/problems/replace-words/
# 难度：中等
# 执行用时：52 ms, 在所有 Python3 提交中击败了 93.52% 的用户
# 内存消耗：28.8 MB, 在所有 Python3 提交中击败了 35.23% 的用户

from typing import List

class Solution:
  def replaceWords(self, dictionary: List[str], sentence: str) -> str:
    # 构造前缀树
    trie = {}
    for root in dictionary:
      current = trie
      for c in root:
        if c not in current:
          current[c] = {}
        elif 'END' in current:
          break
        current = current[c]
      current['END'] = True
    # 遍历每个单词，检测是否能在前缀树中匹配到词根
    words = sentence.split(' ')
    for i in range(len(words)):
      word = words[i]
      root = word
      tmp = ''
      current = trie
      for c in word:
        if c in current:
          current = current[c]
          tmp += c
          if 'END' in current:
            root = tmp
            break
        else:
          break
      words[i] = root
    return ' '.join(words)

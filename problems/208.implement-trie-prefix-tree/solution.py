# 208. 实现 Trie (前缀树)
# URL：https://leetcode-cn.com/problems/implement-trie-prefix-tree/
# 难度：中等
# 关键词：前缀树、字典树
# 执行用时：88 ms, 在所有 Python3 提交中击败了 99.60% 的用户
# 内存消耗：29.2 MB, 在所有 Python3 提交中击败了 70.92% 的用户

class Trie:
  def __init__(self):
    self.trie = {}

  def insert(self, word: str) -> None:
    current = self.trie
    for c in word:
      if c not in current:
        current[c] = {}
      current = current[c]
    current['#'] = True  # word's ending flag

  def search(self, word: str) -> bool:
    current = self.trie
    for c in word:
      if c not in current:
        return False
      current = current[c]
    return True if '#' in current else False

  def startsWith(self, prefix: str) -> bool:
    current = self.trie
    for c in prefix:
      if c not in current:
        return False
      current = current[c]
    return True

# 2156. 查找给定哈希值的子串
# URL：https://leetcode-cn.com/problems/find-substring-with-given-hash-value/
# 难度：中等
# 关键词：模运算、滑动窗口、反向滑窗、哈希函数、滚动哈希
# 执行用时：160 ms, 在所有 Python3 提交中击败了 98.38% 的用户
# 内存消耗：15.6 MB, 在所有 Python3 提交中击败了 30.85% 的用户

def generateCharValues(string: str):
  return list(map(lambda c: ord(c) - ord('a') + 1, string))

class Solution:
  def subStrHash(self, s: str, power: int, modulo: int, k: int, hashValue: int) -> str:
    charValues, length = generateCharValues(s), len(s)
    # special case
    if power % modulo == 0:
      for i in range(length - k + 1):
        if (charValues[i]) % modulo == hashValue:
          return s[i:i+k]
    # compute the first window hash
    basePower, maxPower, windowHash, position = power % modulo, 0, 0, -1
    for i in range(length - k, length):
      maxPower = (maxPower * basePower) % modulo if maxPower != 0 else 1
      windowHash = (windowHash + charValues[i] * maxPower) % modulo
    # forward search
    if windowHash == hashValue:
      position = length - k
    for i in range(length - k - 1, -1, -1):
      windowHash = (((windowHash - (charValues[i + k] * maxPower)) * basePower) + charValues[i]) % modulo
      if windowHash == hashValue:
        position = i
    return s[position:position+k]

# 6002. 设计位集
# URL：https://leetcode-cn.com/problems/design-bitset/
# 难度：中等
# 关键词：位运算
# 执行用时：4700 ms, 在所有 Python3 提交中击败了 33.33% 的用户
# 内存消耗：44.5 MB, 在所有 Python3 提交中击败了 100.00% 的用户

class Bitset:
  def __init__(self, size: int):
    self.size = size
    self.value = 0
    self.positiveCount = 0
    self.flipFlag = False

  def fix(self, idx: int) -> None:
    mask = 1 << idx
    if self.flipFlag:
      if self.value & mask:
        self.value ^= 1 << idx
        self.positiveCount -= 1
    else:
      if not (self.value & mask):
        self.value |= mask
        self.positiveCount += 1

  def unfix(self, idx: int) -> None:
    mask = 1 << idx
    if self.flipFlag:
      if not (self.value & mask):
        self.value |= mask
        self.positiveCount += 1
    else:
      if self.value & mask:
        self.value ^= 1 << idx
        self.positiveCount -= 1

  def realFlip(self) -> None:
    self.flipFlag = False
    self.value ^= (2 ** self.size) - 1
    self.positiveCount = self.size - self.positiveCount

  def flip(self) -> None:
    self.flipFlag = not self.flipFlag

  def all(self) -> bool:
    if self.flipFlag:
      return self.value == 0
    else:
      return self.value == (2 ** self.size) - 1

  def one(self) -> bool:
    if self.flipFlag:
      return self.value != (2 ** self.size) - 1
    else:
      return self.value > 0

  def count(self) -> int:
    if self.flipFlag:
      return self.size - self.positiveCount
    else:
      return self.positiveCount

  def toString(self) -> str:
    if self.flipFlag:
      self.realFlip()
    s = bin(self.value)[2:]
    s = '0' * (self.size - len(s)) + s
    s = s[::-1]
    return s

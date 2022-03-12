# 2166. 设计位集
# URL：https://leetcode-cn.com/problems/design-bitset/
# 难度：中等
# 关键词：位运算、懒操作
# 执行用时：764 ms, 在所有 Python3 提交中击败了 66.67% 的用户
# 内存消耗：44.6 MB, 在所有 Python3 提交中击败了 100.00% 的用户

class Bitset:
  def __init__(self, size: int):
    self.size = size
    self._value = 0        # 存储二进制数
    self._flipped = False  # 是否为翻转状态
    self._count = 0        # 1 的个数

  def _realFix(self, idx: int) -> None:
    mask = 1 << idx
    if not (self._value & mask):
      self._value |= mask
      self._count += 1

  def _realUnfix(self, idx: int) -> None:
    mask = 1 << idx
    if self._value & mask:
      self._value ^= mask
      self._count -= 1

  def _realFlip(self) -> None:
    self._flipped = False
    self._value ^= (2 ** self.size) - 1
    self._count = self.size - self._count

  def fix(self, idx: int) -> None:
    self._realUnfix(idx) if self._flipped else self._realFix(idx)

  def unfix(self, idx: int) -> None:
    self._realFix(idx) if self._flipped else self._realUnfix(idx)

  def flip(self) -> None:
    self._flipped = not self._flipped

  def all(self) -> bool:
    return (self._count == 0) if self._flipped else (self._count == self.size)

  def one(self) -> bool:
    return (self._count != self.size) if self._flipped else (self._count > 0)

  def count(self) -> int:
    return (self.size - self._count) if self._flipped else (self._count)

  def toString(self) -> str:
    if self._flipped:
      self._realFlip()
    s = bin(self._value)[:1:-1]
    s += '0' * (self.size - len(s))
    return s

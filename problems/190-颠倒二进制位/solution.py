# 190. 颠倒二进制位
# URL：https://leetcode-cn.com/problems/reverse-bits/
# 执行结果：通过
# 执行用时：32 ms, 在所有 Python3 提交中击败了 78.04% 的用户
# 内存消耗：14.6 MB, 在所有 Python3 提交中击败了 99.72% 的用户

class Solution:
  def reverseBits(self, n: int) -> int:
    result = 0
    for i in range(32):
      if n & (1 << i):
        result |= 1 << (31 - i)
    return result

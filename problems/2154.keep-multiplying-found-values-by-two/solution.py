# 2154. 将找到的值乘以 2
# URL：https://leetcode-cn.com/problems/keep-multiplying-found-values-by-two/
# 难度：简单
# 关键词：数组、集合、哈希表
# 执行用时：32 ms, 在所有 Python3 提交中击败了 92.63% 的用户
# 内存消耗：15 MB, 在所有 Python3 提交中击败了 96.75% 的用户

from typing import List

class Solution:
  def findFinalValue(self, nums: List[int], original: int) -> int:
    while True:
      if original in nums:
        original *= 2
      else:
        return original

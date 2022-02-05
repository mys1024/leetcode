# 1748. 唯一元素的和
# URL：https://leetcode-cn.com/problems/sum-of-unique-elements/
# 难度：简单
# 关键词：数组
# 执行用时：28 ms, 在所有 Python3 提交中击败了 91.79% 的用户
# 内存消耗：14.9 MB, 在所有 Python3 提交中击败了 84.33% 的用户

from typing import List, Dict

class Solution:
  def sumOfUnique(self, nums: List[int]) -> int:
    flags: Dict[int, bool] = {}
    s = 0
    for num in nums:
      try:
        if flags[num]:
          s -= num
          flags[num] = False
      except:
        flags[num] = True
        s += num
    return s

# 1748. 唯一元素的和
# URL：https://leetcode-cn.com/problems/sum-of-unique-elements/
# 难度：简单
# 关键词：数组
# 执行用时：24 ms, 在所有 Python3 提交中击败了 98.01% 的用户
# 内存消耗：14.8 MB, 在所有 Python3 提交中击败了 89.80% 的用户

from collections import Counter
from typing import List

class Solution:
  def sumOfUnique(self, nums: List[int]) -> int:
    counter = Counter(nums)
    return sum(map(lambda i: i[0], filter(lambda i: i[1] == 1, counter.items())))

# 5985. 根据给定数字划分数组
# URL：https://leetcode-cn.com/problems/partition-array-according-to-given-pivot/
# 难度：中等
# 关键词：数组、排序、分组
# 执行用时：116 ms, 在所有 Python3 提交中击败了 100.00% 的用户
# 内存消耗：28.6 MB, 在所有 Python3 提交中击败了 100.00% 的用户

from typing import List

class Solution:
  def pivotArray(self, nums: List[int], pivot: int) -> List[int]:
    small = []
    big = []
    eq = 0
    for num in nums:
      if num < pivot:
        small.append(num)
      elif num == pivot:
        eq += 1
      elif num > pivot:
        big.append(num)
    return small + [pivot] * eq + big

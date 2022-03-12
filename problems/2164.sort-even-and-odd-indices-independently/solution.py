# 2164. 对奇偶下标分别排序
# URL：https://leetcode-cn.com/problems/sort-even-and-odd-indices-independently/
# 难度：简单
# 关键词：数组、排序
# 执行用时：32 ms, 在所有 Python3 提交中击败了 50.00% 的用户
# 内存消耗：14.8 MB, 在所有 Python3 提交中击败了 100.00% 的用户

from typing import List

class Solution:
  def sortEvenOdd(self, nums: List[int]) -> List[int]:
    nums[::2] = sorted(nums[::2])
    nums[1::2] = sorted(nums[1::2], reverse=True)
    return nums

# 1985. 找出数组中的第 K 大整数
# URL：https://leetcode-cn.com/problems/find-the-kth-largest-integer-in-the-array/
# 难度：中等
# 执行用时：68 ms, 在所有 Python3 提交中击败了 85.65% 的用户
# 内存消耗：21.1 MB, 在所有 Python3 提交中击败了 12.96% 的用户

from typing import List

class Solution:
  def kthLargestNumber(self, nums: List[str], k: int) -> str:
    nums.sort(key=int, reverse=True)
    return nums[k - 1]

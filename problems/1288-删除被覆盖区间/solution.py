# 1288. 删除被覆盖区间
# URL：https://leetcode-cn.com/problems/remove-covered-intervals/
# 难度：中等
# 关键词：集合、区间
# 执行用时：36 ms, 在所有 Python3 提交中击败了 84.94% 的用户
# 内存消耗：15.4 MB, 在所有 Python3 提交中击败了 37.05% 的用户

from typing import List, Tuple

def isCovered(a: Tuple[int, int], b: Tuple[int, int]):
  if a == b:
    return False
  a_low, a_high = a
  b_low, b_high = b
  return a_low <= b_low and b_high <= a_high

class Solution:
  def removeCoveredIntervals(self, intervals: List[List[int]]) -> int:
    checkedIntervals = set()
    for interval in intervals:
      interval = tuple(interval)
      shouldAdd = True
      shouldRemove = set()
      for checkedInterval in checkedIntervals:
        if isCovered(interval, checkedInterval):
          shouldRemove.add(checkedInterval)
        elif isCovered(checkedInterval, interval):
          shouldAdd = False
          break
      checkedIntervals -= shouldRemove
      if shouldAdd:
        checkedIntervals.add(interval)
    return len(checkedIntervals)

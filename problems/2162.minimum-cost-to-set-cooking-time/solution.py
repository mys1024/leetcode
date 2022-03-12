# 2162. 设置时间的最少代价
# URL：https://leetcode-cn.com/problems/minimum-cost-to-set-cooking-time/
# 难度：中等
# 关键词：数组、排序、分组
# 执行用时：32 ms, 在所有 Python3 提交中击败了 100.00% 的用户
# 内存消耗：15 MB, 在所有 Python3 提交中击败了 33.33% 的用户

from typing import List, Tuple

def cost(startAt: int, moveCost: int, pushCost: int, plan: List[int]):
  c = 0
  for num in plan:
    if startAt != num:
      c += moveCost
      startAt = num
    c += pushCost
  return c

class Solution:
  def minCostSetTime(self, startAt: int, moveCost: int, pushCost: int, targetSeconds: int) -> int:
    minAndSecList: List[Tuple[int, int]] = []
    minutes, seconds = targetSeconds // 60, targetSeconds % 60
    if minutes < 100:
      minAndSecList.append((minutes, seconds))
    minutes, seconds = minutes - 1, seconds + 60
    if minutes > -1 and seconds < 100:
      minAndSecList.append((minutes, seconds))
    # generate cases
    cases: List[List[int]] = []
    for minutes, seconds  in minAndSecList:
      case = []
      if minutes > 0:
        case += list(map(int, str(minutes)))
        if seconds < 10:
          case.append(0)
      case += list(map(int, str(seconds)))
      cases.append(case)
    # compute the minimum cost
    minimum = 10 ** 8
    for case in cases:
      c = cost(startAt, moveCost, pushCost, case)
      minimum = c if c < minimum else minimum
    return minimum

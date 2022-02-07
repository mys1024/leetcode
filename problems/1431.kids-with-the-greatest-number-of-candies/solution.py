# 1431. 拥有最多糖果的孩子
# URL：https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/
# 难度：简单
# 关键词：数组
# 执行用时：36 ms, 在所有 Python3 提交中击败了 47.77% 的用户
# 内存消耗：14.9 MB, 在所有 Python3 提交中击败了 87.71% 的用户

from typing import List

class Solution:
  def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
    maxCandies = max(candies)
    return list(map(lambda x: x + extraCandies >= maxCandies, candies))

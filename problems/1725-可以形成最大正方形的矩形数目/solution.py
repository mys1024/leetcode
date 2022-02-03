# 1725. 可以形成最大正方形的矩形数目
# URL：https://leetcode-cn.com/problems/number-of-rectangles-that-can-form-the-largest-square/
# 难度：简单
# 关键词：数组
# 执行用时：44 ms, 在所有 Python3 提交中击败了 63.33% 的用户
# 内存消耗：15.3 MB, 在所有 Python3 提交中击败了 43.33% 的用户

from typing import List

class Solution:
  def countGoodRectangles(self, rectangles: List[List[int]]) -> int:
    maxSquareLen = -1
    count = 0
    for rectangle in rectangles:
      squareLen = min(rectangle[0], rectangle[1])
      if squareLen > maxSquareLen:
        maxSquareLen = squareLen
        count = 1
      elif squareLen == maxSquareLen:
        count += 1
    return count

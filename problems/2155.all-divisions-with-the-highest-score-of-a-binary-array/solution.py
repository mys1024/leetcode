# 2155. 分组得分最高的所有下标
# URL：https://leetcode-cn.com/problems/all-divisions-with-the-highest-score-of-a-binary-array/
# 难度：中等
# 关键词：前缀和、哈希表
# 执行用时：736 ms, 在所有 Python3 提交中击败了 45.83% 的用户
# 内存消耗：40.5 MB, 在所有 Python3 提交中击败了 12.32% 的用户

from typing import List, Dict

class Solution:
  def maxScoreIndices(self, nums: List[int]) -> List[int]:
    n = len(nums)
    # 计算左数组得分的前缀和
    left = [0]
    prev = 0
    for i in range(n):
      prev = left[i] + (1 if nums[i] == 0 else 0)
      left.append(prev)
    # 计算得分
    scores: Dict[int, List[int]] = {}
    maxScore = 0
    for i in range(n + 1):
      s = n - i - left[-1] + 2 *left[i] # 经推导得出该式
      if s < maxScore:
        continue
      maxScore = s
      if s in scores:
        scores[s].append(i)
      else:
        scores[s] = [i]
    return scores[maxScore]

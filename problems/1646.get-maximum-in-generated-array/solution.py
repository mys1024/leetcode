# 1646. 获取生成数组中的最大值
# URL：https://leetcode-cn.com/problems/get-maximum-in-generated-array/
# 难度：简单
# 关键词：数组
# 执行用时：36 ms, 在所有 Python3 提交中击败了 37.47% 的用户
# 内存消耗：14.9 MB, 在所有 Python3 提交中击败了 69.18% 的用户

class Solution:
  def getMaximumGenerated(self, n: int) -> int:
    if n < 2:
      return n
    nums = [0, 1] + [None] * (n - 1)
    for i in range(2, n + 1):
      mid = i // 2
      nums[i] = nums[mid] + (0 if i % 2 == 0 else nums[mid + 1])
    return max(nums)

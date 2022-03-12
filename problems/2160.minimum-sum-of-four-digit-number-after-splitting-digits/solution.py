# 2160. 拆分数位后四位数字的最小和
# URL：https://leetcode-cn.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/
# 难度：简单
# 关键词：数论
# 执行用时：32 ms, 在所有 Python3 提交中击败了 50.00% 的用户
# 内存消耗：14.9 MB, 在所有 Python3 提交中击败了 100.00% 的用户

class Solution:
  def minimumSum(self, num: int) -> int:
    nums = list(map(int, str(num)))
    nums.sort()
    return nums[0] * 10 + nums[2] + nums[1] * 10 + nums[3]

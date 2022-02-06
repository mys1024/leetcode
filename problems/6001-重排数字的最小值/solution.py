# 6001. 重排数字的最小值
# URL：https://leetcode-cn.com/problems/smallest-value-of-the-rearranged-number/
# 难度：中等
# 关键词：字符串、排序
# 执行用时：40 ms, 在所有 Python3 提交中击败了 100.00% 的用户
# 内存消耗：14.9 MB, 在所有 Python3 提交中击败了 100.00% 的用户

class Solution:
  def smallestNumber(self, num: int) -> int:
    negative = num < 0
    if negative:
      num = -num
    digits = list(map(int, str(num)))
    digits.sort(reverse=negative)
    if not negative:
      for i in range(len(digits)):
        if digits[i] != 0:
          digits[0], digits[i] = digits[i], digits[0]
          break
    result = int(''.join(map(str, digits)))
    return -result if negative else result

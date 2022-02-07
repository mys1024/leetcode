# 1414. 和为 K 的最少斐波那契数字数目
# URL：https://leetcode-cn.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/
# 难度：中等
# 关键词：贪心
# 执行用时：40 ms, 在所有 Python3 提交中击败了 44.59% 的用户
# 内存消耗：14.9 MB, 在所有 Python3 提交中击败了 72.97% 的用户

class Solution:
  def findMinFibonacciNumbers(self, k: int) -> int:
    # 生成小于 k 的斐波那契数
    fibNums = [1, 1]
    prev0, prev1 = 1, 1
    while True:
      fibNum = prev0 + prev1
      if fibNum < k:
        fibNums.append(fibNum)
        prev0, prev1 = prev1, fibNum
      elif fibNum == k:
        return 1
      else:
        break
    # 使用贪心策略开始计数
    fibNums.reverse()
    count = 0
    s = 0
    for fibNum in fibNums:
      while True:
        if s + fibNum < k:
          s += fibNum
          count += 1
        elif s + fibNum == k:
          return count + 1
        else:
          break

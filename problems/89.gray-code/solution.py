# 89. 格雷编码
# URL：https://leetcode-cn.com/problems/gray-code/
# 难度：中等
# 关键词：位操作、格雷码、镜像反射法
# 执行用时：52 ms, 在所有 Python3 提交中击败了 35.95% 的用户
# 内存消耗：18.4 MB, 在所有 Python3 提交中击败了 21.96% 的用户

# 题解参考：https://leetcode-cn.com/problems/gray-code/solution/gray-code-jing-xiang-fan-she-fa-by-jyd/

from typing import List

class Solution:
  def grayCode(self, n: int) -> List[int]:
    result = [0, 1]
    for i in range(1, n):
      cp = result.copy()
      cp.reverse()
      for j in range(len(cp)):
        cp[j] |= 1 << i
      result += cp
    return result

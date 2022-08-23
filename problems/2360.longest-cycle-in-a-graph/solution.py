# 2360. 图中的最长环
# URL：https://leetcode.cn/problems/longest-cycle-in-a-graph/
# 难度：困难
# 关键词：图、拓扑排序、广度优先搜索
# 执行用时：396 ms, 在所有 Python 提交中击败了 35.58% 的用户
# 内存消耗：24.7 MB, 在所有 Python 提交中击败了 99.91% 的用户

from typing import List

class Solution:
  def longestCycle(self, edges: List[int]) -> int:
    # 构建入度表
    inDegree: List[int] = [0] * len(edges)
    for i in range(len(edges)):
      t = edges[i]
      if (t != -1):
        inDegree[t] += 1

    # 删去不构成环的结点
    queue = []
    for i in range(len(inDegree)):
      if (inDegree[i] == 0):
        queue.append(i)
    while len(queue) > 0:
      i = queue.pop()
      t = edges[i]
      if (t == -1):
        continue
      edges[i] = -1
      inDegree[t] -= 1
      if (inDegree[t] == 0):
        queue.append(t)

    # 将环逐个消除并记录最大的环长度
    def bfs(start):
      queue = [(0, start)]
      while (len(queue) != 0):
        d, i = queue.pop(0)
        t = edges[i]
        if (t == -1):
          return d
        queue.append((d + 1, t))
        edges[i] = -1
        inDegree[t] -= 1
    maximum = -1
    for i in range(len(inDegree)):
      if (inDegree[i] != 0):
        d = bfs(i)
        maximum = max(maximum, d)

    # 返回答案
    return maximum

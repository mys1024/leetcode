# 107. 二叉树的层序遍历 II
# URL：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
# 难度：中等
# 关键词：树、二叉树、层序遍历、广度优先搜索
# 执行用时：40 ms, 在所有 Python3 提交中击败了 22.91% 的用户
# 内存消耗：15.3 MB, 在所有 Python3 提交中击败了 47.05% 的用户

from typing import List

# Definition for a binary tree node.
class TreeNode:
  def __init__(self, val=0, left=None, right=None):
    self.val = val
    self.left = left
    self.right = right

class Solution:
  def levelOrderBottom(self, root: TreeNode) -> List[List[int]]:
    if not root:
      return []
    queue, result = [(root, 1)], []
    currentNodes, currentLevel = [], 1
    while len(queue) > 0:
      node, level = queue.pop(0)
      if node.left:
        queue.append((node.left, level + 1))
      if node.right:
        queue.append((node.right, level + 1))
      if currentLevel != level:
        currentLevel = level
        result.append(currentNodes)
        currentNodes = []
      currentNodes.append(node.val)
    result.append(currentNodes)
    result.reverse()
    return result

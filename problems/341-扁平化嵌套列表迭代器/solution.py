# 341. 扁平化嵌套列表迭代器
# URL：https://leetcode-cn.com/problems/flatten-nested-list-iterator/
# 难度：中等
# 执行用时：56 ms, 在所有 Python3 提交中击败了 87.02% 的用户
# 内存消耗：18.6 MB, 在所有 Python3 提交中击败了 5.44% 的用户

from __future__ import annotations
from typing import List

class NestedInteger:
  '''
  题目中 NestedInteger 接口的一个实现。
  仅用于测试题解及提供类型提示，不是包含答案的一部分。
  '''
  def __init__(self, val: int | List[NestedInteger]) -> None:
    self.val = val
  def isInteger(self) -> bool:
    return not isinstance(self.val, list)
  def getInteger(self) -> int:
    return None if isinstance(self.val, list) else self.val
  def getList(self) -> List[NestedInteger]:
    return self.val if isinstance(self.val, list) else None

def spread(nestedList: List[NestedInteger]):
  result: List[int] = []
  for i in nestedList:
    if i.isInteger():
      result.append(i.getInteger())
    else:
      result += spread(i.getList())
  return result

class NestedIterator:
  def __init__(self, nestedList: List[NestedInteger]):
    self.list = spread(nestedList)
    self.length = len(self.list)
    self.index = 0
  def next(self) -> int:
    value = self.list[self.index]
    self.index += 1
    return value
  def hasNext(self) -> bool:
    return self.index != self.length

# testCase = [NestedInteger(5), NestedInteger(6), NestedInteger([])]
# it = NestedIterator(testCase)
# while it.hasNext():
#   print(it.next())

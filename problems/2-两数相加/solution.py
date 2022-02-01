# 2. 两数相加
# URL：https://leetcode-cn.com/problems/add-two-numbers/
# 执行结果：通过
# 执行用时：56 ms, 在所有 Python3 提交中击败了 63.00% 的用户
# 内存消耗：14.9 MB, 在所有 Python3 提交中击败了 80.69% 的用户

class ListNode:
  def __init__(self, val=0, next=None):
    self.val = val
    self.next = next

def toNumber(listNode: ListNode) -> int:
  num = 0
  count = 0
  while listNode:
    num += listNode.val * (10 ** count)
    count += 1
    listNode = listNode.next
  return num

def toNodeList(number: int) -> ListNode:
  head = None
  for d in map(int, str(number)):
    head = ListNode(d, head)
  return head

class Solution:
  def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
    num1 = toNumber(l1)
    num2 = toNumber(l2)
    return toNodeList(num1 + num2)

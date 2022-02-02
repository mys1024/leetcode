# 2000. 反转单词前缀
# URL：https://leetcode-cn.com/problems/reverse-prefix-of-word/
# 难度：简单
# 关键词：双指针、字符串
# 执行用时：36 ms, 在所有 Python3 提交中击败了 34.19% 的用户
# 内存消耗：15 MB, 在所有 Python3 提交中击败了 60.42% 的用户

def reverse(string: str):
  reversed = ''
  for c in string:
    reversed = c + reversed
  return reversed

class Solution:
  def reversePrefix(self, word: str, ch: str) -> str:
    position = str.find(word, ch)
    if position > 0:
      return reverse(word[0:position+1]) + word[position+1:]
    else:
      return word

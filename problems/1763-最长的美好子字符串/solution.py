# 1763. 最长的美好子字符串
# URL：https://leetcode-cn.com/problems/longest-nice-substring/
# 执行结果：通过
# 执行用时：152 ms, 在所有 Python3 提交中击败了 20.67% 的用户
# 内存消耗：15 MB, 在所有 Python3 提交中击败了 71.63% 的用户

def isNice(string: str, start: int, end: int) -> bool:
  string = string[start : end + 1]
  return len(set(string)) == len(set(string.lower())) * 2

class Solution:
  def longestNiceSubstring(self, s: str) -> str:
    length = len(s)
    longest = ''
    for i in range(length):
      for j in range(i, length):
        substr = s[i : j + 1]
        if isNice(s, i, j) and len(substr) > len(longest):
          longest = substr
    return longest

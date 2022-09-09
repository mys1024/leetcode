// 1598. 文件夹操作日志搜集器
// URL: https://leetcode.cn/problems/crawler-log-folder/
// 难度: 简单
// 关键词: 字符串
// 执行用时: 76 ms, 在所有 TypeScript 提交中击败了 12.50% 的用户
// 内存消耗: 43.9 MB, 在所有 TypeScript 提交中击败了 12.50% 的用户

function minOperations(logs: string[]): number {
  let cnt = 0
  for (const l of logs) {
    if (l === './') {}
    else if (l === '../') {
      if (cnt > 0)
        cnt -= 1
    }
    else {
      cnt++
    }
  }
  return cnt
}

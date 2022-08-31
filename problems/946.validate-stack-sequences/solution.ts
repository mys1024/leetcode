// 946. 验证栈序列
// URL：https://leetcode.cn/problems/validate-stack-sequences/
// 难度：中等
// 关键词：栈、数组
// 执行用时：72 ms, 在所有 TypeScript 提交中击败了 31.91% 的用户
// 内存消耗：44.2 MB, 在所有 TypeScript 提交中击败了 57.45% 的用户

function validateStackSequences(pushed: number[], popped: number[]): boolean {
  const stack: number[] = []
  let j = 0
  for (let i = 0; i < pushed.length; i++) {
    if (pushed[i] !== popped[j]) {
      stack.push(pushed[i])
      continue
    }
    j++
    for (let k = stack.length - 1; k >= 0 && j < popped.length; k--, j++) {
      if (stack[k] !== popped[j])
        break
      stack.pop()
    }
  }
  for (let k = stack.length - 1; k >= 0; k--, j++) {
    if (stack[k] !== popped[j])
      return false
  }
  return true
}

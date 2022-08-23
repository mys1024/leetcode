// 1460. 通过翻转子数组使两个数组相等
// URL：https://leetcode.cn/problems/make-two-arrays-equal-by-reversing-sub-arrays/
// 难度：简单
// 关键词：数组
// 执行用时：72 ms, 在所有 JavaScript 提交中击败了 74.07% 的用户
// 内存消耗：44.2 MB, 在所有 JavaScript 提交中击败了 62.96% 的用户

function same(arr1: any[], arr2: any[]) {
  if (arr1.length !== arr2.length) {
    return false
  }
  for (let i = 0; i < arr1.length; i++) {
    if (arr1[i] !== arr2[i]) {
      return false
    }
  }
  return true
}

function canBeEqual(target: number[], arr: number[]): boolean {
  return same(target.sort(), arr.sort())
}

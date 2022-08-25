// 658. 找到 K 个最接近的元素
// URL：https://leetcode.cn/problems/find-k-closest-elements/
// 难度：中等
// 关键词：数组
// 执行用时：144 ms, 在所有 TypeScript 提交中击败了 6.25% 的用户
// 内存消耗：48.8 MB, 在所有 TypeScript 提交中击败了 33.33% 的用户

function findClosestElements(arr: number[], k: number, x: number): number[] {
  const tmp = arr.map(num => [Math.abs(x - num), num])
  tmp.sort((a, b) => a[0] - b[0])

  const ans: number[] = []
  for (let i = 0; i < k; i++) {
      ans.push(tmp[i][1])
  }
  ans.sort((a, b) => a - b)

  return ans
}

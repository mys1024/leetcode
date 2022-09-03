// 881. 救生艇
// URL：https://leetcode.cn/problems/boats-to-save-people/
// 难度：中等
// 关键词：数组、贪心
// 执行用时：156 ms, 在所有 TypeScript 提交中击败了 28.57% 的用户
// 内存消耗：50.1 MB, 在所有 TypeScript 提交中击败了 50.00% 的用户

function numRescueBoats(people: number[], limit: number): number {
  const n = people.length
  people.sort((a, b) => a - b)

  let cnt = 0
  let i = 0
  let j = n - 1

  while (i <= j) {
    if (people[i] + people[j] <= limit)
      i++
    j--
    cnt++
  }

  return cnt
}

// 2303. 计算应缴税款总额
// URL: https://leetcode.cn/problems/calculate-amount-paid-in-taxes/
// 难度: 简单
// 关键词: 模拟
// 执行用时: 64ms, 在所有 TypeScript 提交中击败了 90.00% 的用户
// 内存消耗: 44.2 MB, 在所有 TypeScript 提交中击败了 30.00% 的用户

function calculateTax(brackets: number[][], income: number): number {
  let prevUpper = 0
  let tax = 0
  for (const [upper, percent] of brackets) {
    const step = upper - prevUpper
    prevUpper = upper
    if (income > step) {
      tax += step * percent / 100
      income -= step
    } else {
      tax += income * percent / 100
      income = 0
      break
    }
  }
  return tax
}

// 1475. 商品折扣后的最终价格
// URL：https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/
// 难度：简单
// 关键词：数组
// 执行用时：68 ms, 在所有 TypeScript 提交中击败了 93.94% 的用户
// 内存消耗：43.9 MB, 在所有 TypeScript 提交中击败了 48.48% 的用户

function finalPrices(prices: number[]): number[] {
  const n = prices.length
  for (let i = 0; i < n; i++) {
    for (let j = i + 1; j < n; j++) {
      if (prices[j] <= prices[i]) {
        prices[i] -= prices[j]
        break
      }
    }
  }
  return prices
}

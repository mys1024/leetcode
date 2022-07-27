// 592. 分数加减运算
// URL：https://leetcode.cn/problems/fraction-addition-and-subtraction/
// 难度：中等
// 关键词：数学、字符串、正则表达式
// 执行用时：64 ms, 在所有 JavaScript 提交中击败了 60.00% 的用户
// 内存消耗：42.4 MB, 在所有 JavaScript 提交中击败了 16.67% 的用户

/**
 * @param {string} expression
 * @return {string}
 */
function fractionAddition(expression) {
  const terms = expression.match(/(?:\+|-)?\d+\/\d+/g)
    .map(t => t.split('/').map(v => Number.parseInt(v)))

  let [a, b] = terms[0]
  for (const term of terms.slice(1)) {
    const [c, d] = term
    a = a*d + c*b
    b = b * d
    g = gcd(Math.abs(a), b)
    a /= g
    b /= g
  }
  return `${a}/${b}`
};

/**
 * @param {number} a
 * @param {number} b
 */
function gcd(a, b) {
  while (b != 0) {
    [a, b] = [b, a % b]
  }
  return a
}

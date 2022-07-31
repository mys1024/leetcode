// 1161. 最大层内元素和
// URL：https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/
// 难度：中等
// 关键词：二叉树、层序遍历
// 执行用时：168 ms, 在所有 JavaScript 提交中击败了 50.00% 的用户
// 内存消耗：61.3 MB, 在所有 JavaScript 提交中击败了 17.24% 的用户

class TreeNode {
  constructor(val, left, right) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
  }
}

/**
 * @param {TreeNode} root
 * @return {number}
 */
function maxLevelSum(root) {
  const queue = [[1, root]]
  const cnt = {}

  while (queue.length !== 0) {
    const [level, node] = queue.pop()
    cnt[level] = (cnt[level] ?? 0) + node.val
    if (node.left) {
      queue.push([level + 1, node.left])
    }
    if (node.right) {
      queue.push([level + 1, node.right])
    }
  }

  let [max, ans] = [-Infinity, 1]
  for (const key of Object.keys(cnt)) {
    if (cnt[key] > max) {
      max = cnt[key]
      ans = key
    }
  }

  return ans
}

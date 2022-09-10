// 669. 修剪二叉搜索树
// URL: https://leetcode.cn/problems/trim-a-binary-search-tree/
// 难度: 中等
// 关键词: 二叉树、深度优先遍历
// 执行用时: 80 ms, 在所有 TypeScript 提交中击败了 40.79% 的用户
// 内存消耗: 48 MB, 在所有 TypeScript 提交中击败了 7.89% 的用户

class TreeNode {
  val: number
  left: TreeNode | null
  right: TreeNode | null
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
  }
}

function trimBST(root: TreeNode | null, low: number, high: number): TreeNode | null {
  if (!root)
    return root
  const left = trimBST(root.left, low, high)
  const right = trimBST(root.right, low, high)
  if (low <= root.val && root.val <= high) {
    root.left = left
    root.right = right
    return root
  }
  return left || right
}

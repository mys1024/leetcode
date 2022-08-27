// 662. 二叉树最大宽度
// URL：https://leetcode.cn/problems/maximum-width-of-binary-tree/
// 难度：中等
// 关键词：二叉树、广度优先搜索、层序遍历
// 执行用时：96 ms, 在所有 Go 提交中击败了 11.11% 的用户
// 内存消耗：50.4 MB, 在所有 Go 提交中击败了 7.41% 的用户

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

function max(nums: bigint[]) {
  let m = -1n
  for (const num of nums) {
    if (num > m)
      m = num
  }
  return m
}

function widthOfBinaryTree(root: TreeNode | null): number {
  if (!root)
    return 0

  const queue: [TreeNode, number, bigint][] = [[root, 0, 0n]]
  const m: Record<number, [bigint, bigint]> = {}

  while (queue.length > 0) {
    const [node, depth, pos] = queue.shift() as [TreeNode, number, bigint]

    if (!m[depth])
      m[depth] = [pos, pos]
    else if (pos < m[depth][0])
      m[depth][0] = pos
    else if (pos > m[depth][1])
      m[depth][1] = pos

    if (node.left)
      queue.push([node.left, depth + 1, pos * 2n])
    if (node.right)
      queue.push([node.right, depth + 1, pos * 2n + 1n])
  }

  const widths = Object.values(m).map(([l, r]) => r - l + 1n)
  return Number(max(widths))
}

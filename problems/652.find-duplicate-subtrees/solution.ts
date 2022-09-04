// 652. 寻找重复的子树
// URL: https://leetcode.cn/problems/find-duplicate-subtrees/
// 难度: 中等
// 关键词: 树 深度优先搜索 序列化 哈希表
// 执行用时: 112 ms, 在所有 TypeScript 提交中击败了 51.02% 的用户
// 内存消耗: 49.3 MB, 在所有 TypeScript 提交中击败了 26.53% 的用户

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

function findDuplicateSubtrees(root: TreeNode | null): Array<TreeNode | null> {
  if (!root)
    return []

  const memo: Map<string, number> = new Map()
  const ans: TreeNode[] = []

  function dfs(node: TreeNode) {
    let s = `${node.val}`
    s += ',' + (node.left ? dfs(node.left) : 'n')
    s += ',' + (node.right ? dfs(node.right) : 'n')

    const flag = memo.get(s)
    if (!flag) {
      memo.set(s, 1)
    } else if (flag === 1) {
      memo.set(s, 2)
      ans.push(node)
    }

    return s
  }

  dfs(root)

  return ans
}

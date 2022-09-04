// 2368. 受限条件下可到达节点的数目
// URL: https://leetcode.cn/problems/reachable-nodes-with-restrictions/
// 难度: 中等
// 关键词: 图 深度优先搜索
// 执行用时: 560ms, 在所有 TypeScript 提交中击败了 30.23% 的用户
// 内存消耗: 111.4 MB, 在所有 TypeScript 提交中击败了 11.63% 的用户

function reachableNodes(n: number, edges: number[][], _restricted: number[]): number {
  const adjacencies: Record<number, number[]> = {}
  const visited: Record<number, boolean> = {}
  const restricted: Record<number, boolean> = {}
  let cnt = 0

  for (let node = 0; node < n; node++) {
    adjacencies[node] = []
    visited[node] = false
  }
  for (const node of _restricted) {
    restricted[node] = true
  }

  for (const [a, b] of edges) {
    adjacencies[a].push(b)
    adjacencies[b].push(a)
  }

  function dfs(node: number) {
    if (visited[node] || restricted[node])
      return
    cnt++
    visited[node] = true
    for (const child of adjacencies[node])
      dfs(child)
  }

  dfs(0)

  return cnt
}

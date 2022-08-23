// 2359. 找到离给定两个节点最近的节点
// URL：https://leetcode.cn/problems/find-closest-node-to-given-two-nodes/
// 难度：中等
// 关键词：图、广度优先搜索
// 执行用时：320 ms, 在所有 JavaScript 提交中击败了 27.67% 的用户
// 内存消耗：70.8 MB, 在所有 JavaScript 提交中击败了 50.32% 的用户

/**
 * @param {number[]} edges
 * @param {number} node1
 * @param {number} node2
 * @return {number}
 */
 function closestMeetingNode(edges, node1, node2) {
  function bfs(start) {
    const distance = []
    const queue = [[0, start]]
    while (queue.length !== 0) {
      const [d, node] = queue.pop()
      distance[node] = d
      if (edges[node] >= 0 && distance[edges[node]] === undefined) {
        queue.push([d + 1, edges[node]])
      }
    }
    return distance
  }

  const distance1 = bfs(node1)
  const distance2 = bfs(node2)

  let [ans, min] = [-1, Infinity]
  for (let i = 0; i < Math.max(distance1.length, distance2.length); i++) {
    const [d1, d2] = [distance1[i], distance2[i]]
    if (d1 === undefined || d2 === undefined) {
      continue
    }
    if (Math.max(d1, d2) < min) {
      min = Math.max(d1, d2)
      ans = i
    }
  }

  return ans
}

// 990. 等式方程的可满足性
// URL：https://leetcode-cn.com/problems/satisfiability-of-equality-equations/
// 难度：中等
// 关键词：并查集
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.6 MB, 在所有 Go 提交中击败了 96.58% 的用户

package main

func find(parents []int, idx int) int {
	if parents[idx] != idx {
		parents[idx] = find(parents, parents[idx]) // 路径压缩
	}
	return parents[idx]
}

func union(parents []int, idx1, idx2 int) {
	if idx1 == idx2 {
		return
	}
	parents[find(parents, idx1)] = find(parents, idx2)
}

func equationsPossible(equations []string) bool {
	// 定义并查集
	parents := make([]int, 26)
	for i := 0; i < 26; i++ {
		parents[i] = i
	}
	// 遍历所有表达式，将所有相等的变量放入相同的并查集
	for _, equation := range equations {
		if equation[1] == '=' {
			v1, v2 := equation[0]-'a', equation[3]-'a'
			union(parents, int(v1), int(v2))
		}
	}
	// 遍历所有表达式，发现不应该相等的变量出现在相同的并查集中，立即返回 false
	for _, equation := range equations {
		if equation[1] == '!' {
			v1, v2 := equation[0]-'a', equation[3]-'a'
			if find(parents, int(v1)) == find(parents, int(v2)) {
				return false
			}
		}
	}
	return true
}

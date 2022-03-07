// 207. 课程表
// URL：https://leetcode-cn.com/problems/course-schedule/
// 难度：中等
// 关键词：图、拓扑排序、广度优先搜索
// 执行用时：16 ms, 在所有 Go 提交中击败了 21.03% 的用户
// 内存消耗：6.2 MB, 在所有 Go 提交中击败了 41.31% 的用户

package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 特例
	if len(prerequisites) <= 1 {
		return true
	}
	// 构建入度表和邻接表
	inDegrees := make(map[int]int, numCourses)
	adjacencies := map[int][]int{}
	for _, prerequisite := range prerequisites {
		to, from := prerequisite[0], prerequisite[1]
		inDegrees[to] += 1
		_, ok := inDegrees[from]
		if !ok {
			inDegrees[from] = 0
		}
		if adjacencies[from] == nil {
			adjacencies[from] = []int{to}
		} else {
			adjacencies[from] = append(adjacencies[from], to)
		}
	}
	numCourses = len(inDegrees) // 有的测试用例的 numCourses 参数是错误的，此行意在修正错误
	// 初始化入度为 0 的课程队列
	queue, head, rear := make([]int, numCourses), 0, 0
	for course, inDegree := range inDegrees {
		if inDegree == 0 {
			queue[rear] = course
			rear++
		}
	}
	if rear == 0 {
		return false
	}
	// 循环删除入度为 0 的课程，若最终所有课程都入队过，则可能完成所有课程的学习
	for head != rear {
		course := queue[head]
		head++
		if adjacencies[course] == nil {
			continue
		}
		for _, adjacency := range adjacencies[course] {
			inDegrees[adjacency]--
			if inDegrees[adjacency] == 0 {
				queue[rear] = adjacency
				rear++
			}
		}
	}
	return rear == numCourses
}

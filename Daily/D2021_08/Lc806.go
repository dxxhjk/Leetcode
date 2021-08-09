package D2021_08

/*
847. 访问所有节点的最短路径
存在一个由 n 个节点组成的无向连通图，图中的节点按从 0 到 n - 1 编号。
给你一个数组 graph 表示这个图。其中，graph[i] 是一个列表，由所有与节点 i 直接相连的节点组成。
返回能够访问所有节点的最短路径的长度。你可以在任一节点开始和停止，也可以多次重访节点，并且可以重用边。
 */

/*
思路：
状态压缩：用一个数字表示目前已经有哪些节点被遛过了，由于 n < 12，所以一个 int 就能存。
广度优先搜索，队列里存的结构体，有三部分，当前节点，当前已经遛过的节点以及当前长度。
此外还需要一个哈希表，记录（当前节点，当前已经遛过的节点）确定的这个唯一状态是否搜索过了。
 */

type bfsStatus struct {
	status int
	length int
}

func toStatus(tmpNode, visitedNodes int) int {
	return tmpNode << 16 + visitedNodes
}

func toTempNode(status int) int {
	return status >> 16
}

func toVisitedNodes(status int) int {
	return status & 0xffff
}

func Lc806ShortestPathLength(graph [][]int) int {
	bfs := make([]bfsStatus, 0)
	visited := make(map[int]bool, 0)
	for i := 0; i < len(graph); i++ {
		bfs = append(bfs, bfsStatus{ toStatus(i, 1 << i), 0})
	}
	for len(bfs) != 0 {
		tempBfsStatus := bfs[0]
		if toVisitedNodes(tempBfsStatus.status) == 1<< len(graph) - 1 {
			return tempBfsStatus.length
		}
		bfs = bfs[1:]
		visited[toStatus(toTempNode(tempBfsStatus.status), toVisitedNodes(tempBfsStatus.status))] = true
		for _, target := range graph[toTempNode(tempBfsStatus.status)] {
			newBfsStatus := bfsStatus{toStatus(target, 1 << target | toVisitedNodes(tempBfsStatus.status)), tempBfsStatus.length + 1}
			if _, ok := visited[toStatus(toTempNode(newBfsStatus.status), toVisitedNodes(newBfsStatus.status))]; !ok {
				bfs = append(bfs, newBfsStatus)
			}
		}
	}
	return 0
}

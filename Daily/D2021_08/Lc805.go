package D2021_08

/*
802. 找到最终的安全状态
在有向图中，以某个节点为起始节点，从该点出发，每一步沿着图中的一条有向边行走。如果到达的节点是终点（即它没有连出的有向边），则停止。
对于一个起始节点，如果从该节点出发，无论每一步选择沿哪条有向边行走，最后必然在有限步内到达终点，则将该起始节点称作是 安全 的。
返回一个由图中所有安全的起始节点组成的数组作为答案。答案数组中的元素应当按 升序 排列。
该有向图有 n 个节点，按 0 到 n - 1 编号，其中 n 是 graph 的节点数。图以下述形式给出：graph[i] 是编号 j 节点的一个列表，满足 (i, j) 是图的一条有向边。
 */

/*
思路：
反向图加拓扑排序
 */

import (
	"sort"
)

func Lc805EventualSafeNodes(graph [][]int) []int {
	ans := make([]int, 0)
	reversedGraph := make([][]int, len(graph))
	outNum := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		reversedGraph[i] = make([]int, 0)
		outNum[i] = len(graph[i])
	}
	for i := 0; i < len(graph); i++ {
		for _, t := range graph[i] {
			reversedGraph[t] = append(reversedGraph[t], i)
		}
	}
	var topo []int
	for i := 0; i < len(graph); i++ {
		if outNum[i] == 0 {
			topo = append(topo, i)
		}
	}
	for len(topo) != 0 {
		tmpNode := topo[0]
		ans = append(ans, tmpNode)
		topo = topo[1:]
		for i := 0; i < len(reversedGraph[tmpNode]); i++ {
			outNum[reversedGraph[tmpNode][i]]--
			if outNum[reversedGraph[tmpNode][i]] == 0 {
				topo = append(topo, reversedGraph[tmpNode][i])
			}
		}
	}
	sort.Ints(ans)
	return ans
}
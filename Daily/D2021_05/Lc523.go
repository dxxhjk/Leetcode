package D2021_05

import (
	"fmt"
	"sort"
)

/*
1707. 与数组中元素的最大异或值
给你一个由非负整数组成的数组 nums。另有一个查询数组 queries，其中 queries[i] = [xi, mi] 。
第 i 个查询的答案是 xi 和任何 nums 数组中不超过 mi 的元素按位异或（XOR）得到的最大值。
换句话说，答案是 max(nums[j] XOR xi) ，其中所有 j 均满足 nums[j] <= mi 。
如果 nums 中的所有元素都大于 mi，最终答案就是 -1 。
返回一个整数数组 answer 作为查询的答案，其中 answer.length == queries.length 且 answer[i] 是第 i 个查询的答案。
 */

/*
思路：
字典树
先将 nums 按照从小到大的顺序排序，queries 按照 mi 从小到大排序，
每次进行一个查寻先将小于 mi 的 num 加入字典树，查找出与 xi 异或和最大的，记录到 ans 中。
 */

type myQuery523 struct {
	x int
	m int
	index int
}

type myQueries523 []myQuery523

func(mq myQueries523) Len() int {
	return len(mq)
}

func(mq myQueries523) Less(i, j int) bool {
	return mq[i].m < mq[j].m
}

func(mq myQueries523) Swap(i, j int) {
	mq[i], mq[j] = mq[j], mq[i]
}

type trieNode523 struct {
	children [2]*trieNode523
}

func(t *trieNode523) add(num int) {
	p := t
	for i := 31; i >= 0; i-- {
		a := (num >> i) & 1
		if p.children[a] == nil {
			newNode := new(trieNode523)
			p.children[a] = newNode
		}
		p = p.children[a]
	}
}

func(t *trieNode523) find(num int) int {
	if t.children[0] == nil && t.children[1] == nil {
		return -1
	}
	p := t
	ans := 0
	for i := 31; i>= 0; i-- {
		a := (num >> i) & 1
		b := 1 - a
		ans <<= 1
		if p.children[b] != nil {
			ans = ans + 1
			p = p.children[b]
		} else {
			p = p.children[a]
		}
	}
	return ans
}

func Lc523MaximizeXor(nums []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	myQueries := make(myQueries523, 0)
	for i, query := range queries {
		myQuery := new(myQuery523)
		myQuery.index = i
		myQuery.x = query[0]
		myQuery.m = query[1]
		myQueries = append(myQueries, *myQuery)
	}
	sort.Sort(myQueries)
	sort.Ints(nums)
	p := 0
	t := new(trieNode523)
	for _, myQuery := range myQueries {
		fmt.Println(myQuery)
		for ; p < len(nums) && nums[p] <= myQuery.m; p++ {
			t.add(nums[p])
		}
		ans[myQuery.index] = t.find(myQuery.x)
	}
	return ans
}
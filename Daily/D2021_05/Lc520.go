package D2021_05

import (
	"Leetcode/Structs"
	"sort"
)

/*
692. 前K个高频单词
给一非空的单词列表，返回前 k 个出现次数最多的单词。
返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。
 */

/*
思路：
哈希表存储词频，小根堆找K大
时间O(nlog(k))
空间O(n)
 */

type kv struct {
	k string
	v int
}

type kvs []kv

func(kvs kvs) Len() int {
	return len(kvs)
}

func(kvs kvs) Less(i, j int) bool {
	if kvs[i].v == kvs[j].v {
		return kvs[i].k < kvs[j].k
	}
	return kvs[i].v > kvs[j].v
}

func(kvs kvs) Swap(i, j int)  {
	kvs[i], kvs[j] = kvs[j], kvs[i]
}

func Lc520TopKFrequent(words []string, k int) []string {
	f := func(a, b interface{}) bool {
		aa, bb := a.(kv), b.(kv)
		if aa.v == bb.v {
			return aa.k < bb.k
		}
		return aa.v > bb.v
	}
	h := Structs.MakeHeap(k, f)
	m := make(map[string]int, 0)
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}
	for k, v := range m {
		h.Push(kv{k, v})
	}
	retInterface := h.GetAll()
	retKV := make(kvs, 0)
	for _, kvInterface := range retInterface {
		kv := kvInterface.(kv)
		retKV = append(retKV, kv)
	}
	sort.Sort(retKV)
	ans := make([]string, 0)
	for _, kv := range retKV {
		ans = append(ans, kv.k)
	}
	return ans
}
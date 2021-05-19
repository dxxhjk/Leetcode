package Structs

import "fmt"

/*
小根堆，找K大
 */

type Heap struct {
	value []int
	len int
}

func MakeHeap(len int) *Heap {
	heap := new(Heap)
	value := make([]int, 0)
	heap.value = value
	heap.len = len
	return heap
}

func(h *Heap) Show() {
	fmt.Println(h.value)
}

func(h *Heap) Push(num int) {
	if len(h.value) < h.len {
		h.value = append(h.value, num)
		p := len(h.value) - 1
		f := (p - 1) / 2
		for f >= 0 && h.value[p] < h.value[f] {
			h.value[p], h.value[f] = h.value[f], h.value[p]
			p = f
			f = (p - 1) / 2
		}
		return
	}
	if num < h.GetMin() {
		return
	}
	h.value[0] = num
	h.adjust(0)
}

func(h *Heap) adjust(p int) {
	l, r := 2 * p + 1, 2 * p + 2
	if l >= h.len {
		return
	}
	if r < h.len {
		if h.value[p] > h.value[l] {
			if h.value[p] > h.value[r] {
				if h.value[l] <= h.value[r] {
					h.value[p], h.value[l] = h.value[l], h.value[p]
					h.adjust(l)
				} else {
					h.value[p], h.value[r] = h.value[r], h.value[p]
					h.adjust(r)
				}
			} else {
				h.value[p], h.value[l] = h.value[l], h.value[p]
				h.adjust(l)
			}
		} else if h.value[p] > h.value[r] {
			h.value[p], h.value[r] = h.value[r], h.value[p]
			h.adjust(r)
		}
	} else {
		if h.value[p] > h.value[l] {
			h.value[p], h.value[l] = h.value[l], h.value[p]
		}
	}
}

func(h *Heap) GetMin() int {
	return h.value[0]
}



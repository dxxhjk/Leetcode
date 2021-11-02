package D2021_10

import "fmt"

type SummaryRanges struct {
	right []int
	left []int
}

func Constructor() SummaryRanges {
	s := new(SummaryRanges)
	s.right = make([]int, 10001)
	s.left = make([]int, 0)
	for i := 0; i < 10001; i++ {
		s.right[i] = -1
	}
	return *s
}

func (ranges *SummaryRanges) AddNum(val int) {
	if ranges.right[val] != -1 {
		return
	}
	ranges.right[val] = val
	ranges.left = append(ranges.left, val)
	index := len(ranges.left) - 1
	for index > 0 && ranges.left[index] < ranges.left[index - 1] {
		ranges.left[index], ranges.left[index - 1] = ranges.left[index - 1], ranges.left[index]
		index--
	}
	fmt.Println(ranges.left, ranges.right)
	if val < 10000 && ranges.right[val + 1] != -1 {
		ranges.right[val] = ranges.right[val + 1]
		if len(ranges.left) - 2 == index {
			ranges.left = ranges.left[:index + 1]
		} else {
			ranges.left = append(ranges.left[:index + 1], ranges.left[index + 2:]...)
		}
	}
	if val > 0 && ranges.right[val - 1] != -1 {
		for i := val - 1 ; i > 0 && ranges.right[i] != -1; i-- {
			ranges.right[i] = ranges.right[i + 1]
		}
		if len(ranges.left) == index + 1 {
			ranges.left = ranges.left[:index]
		} else {
			ranges.left = append(ranges.left[:index], ranges.left[index + 1:]...)
		}
	}
}

func (ranges *SummaryRanges) GetIntervals() [][]int {
	ans := make([][]int, 0)
	for _, r := range ranges.left {
		tempRange := make([]int, 2)
		tempRange[0] = r
		tempRange[1] = ranges.right[r]
		ans = append(ans, tempRange)
	}
	return ans
}
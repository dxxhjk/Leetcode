package D2021_11

/*

 */

type Lc1122Solution struct {
	nums []int
}


func Lc1122Constructor(nums []int) Lc1122Solution {
	solution := new(Lc1122Solution)
	solution.nums = nums
	return *solution
}


func (this *Lc1122Solution) Lc1122Reset() []int {
	return this.nums
}


func (this *Lc1122Solution) Lc1122Shuffle() []int {
	m := make(map[int]int)
	for _, num := range this.nums {
		m[num]++
	}
	ans := make([]int, 0)
	for k := range m {
		ans = append(ans, k)
	}
	return ans
}

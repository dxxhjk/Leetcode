package D2021_07

/*
剑指 Offer 42. 连续子数组的最大和
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
 */

func Lc717MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	ans := nums[0]
	for _, num := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += num
		if sum > ans {
			ans = sum
		}
	}
	return ans
}

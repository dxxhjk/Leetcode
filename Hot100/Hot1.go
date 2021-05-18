package Hot100

/*
1. 两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
 */

/*
思路：
用一个map存读过的数，往后读的时候到map里查有无需要的数
 */

func Hot1TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		_, ok := numMap[nums[i]]
		if ok {
			ans := make([]int, 0)
			ans = append(ans, numMap[nums[i]])
			ans = append(ans, i)
			return ans
		} else {
			numMap[target - nums[i]] = i
		}
	}
	return nil
}

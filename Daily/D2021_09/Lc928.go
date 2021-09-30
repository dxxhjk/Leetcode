package D2021_09

/*
437. 路径总和 III
给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
*/

import "Leetcode/Structs"

func Lc928PathSum(root *Structs.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	pre := make(map[int]int, 0)
	pre[0] = 1
	return dfs928(root, 0, targetSum, pre)
}

func dfs928(root *Structs.TreeNode, tempSum, targetSum int, pre map[int]int) int {
	if root == nil {
		return 0
	}
	tempSum += root.Val.(int)
	num := pre[targetSum-tempSum]
	pre[tempSum]++
	num += dfs928(root.Left, tempSum, targetSum, pre)
	num += dfs928(root.Right, tempSum, targetSum, pre)
	pre[tempSum]--
	return num
}

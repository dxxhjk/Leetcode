package D2021_11

import "Leetcode/Structs"

/*
给定一个二叉树，计算 整个树 的坡度 。
一个树的 节点的坡度 定义即为，该节点左子树的节点之和和右子树节点之和的 差的绝对值 。
如果没有左子树的话，左子树的节点之和为 0 ；没有右子树的话也是一样。空结点的坡度是 0 。
整个树的坡度就是其所有节点的坡度之和。
 */

func Lc1118FindTilt(root *Structs.TreeNode) int {
	ans, _ := findTiltAndSum(root)
	return ans
}

func findTiltAndSum(root *Structs.TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftTilt, leftSum := findTiltAndSum(root.Left)
	rightTilt, rightSum := findTiltAndSum(root.Right)
	return leftTilt + rightTilt + abs(leftSum - rightSum), root.Val.(int) + leftSum + rightSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

package D2021_11

import "Leetcode/Structs"

/*
700. 二叉搜索树中的搜索
给定二叉搜索树（BST）的根节点和一个值。
你需要在BST中找到节点值等于给定值的节点。
返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
 */

func Lc1126SearchBST(root *Structs.TreeNode, val int) *Structs.TreeNode {
	if root == nil {
		return nil
	}
	if val == root.Val {
		return root
	}
	if val > root.Val.(int) {
		return Lc1126SearchBST(root.Right, val)
	}
	return Lc1126SearchBST(root.Left, val)
}

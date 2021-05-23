package Daily

import "Leetcode/Structs"
/*
872. 叶子相似的树
请考虑一棵二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个叶值序列 。
举个例子，如上图所示，给定一棵叶值序列为 (6, 7, 4, 9, 8) 的树。
如果有两棵二叉树的叶值序列是相同，那么我们就认为它们是叶相似的。
如果给定的两个根结点分别为 root1 和 root2 的树是叶相似的，则返回 true；否则返回 false 。
 */

/*
思路：
深度优先遍历
 */

func dfs(leaf *[]int, root *Structs.TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leaf = append(*leaf, root.Val)
		return
	}
	dfs(leaf, root.Left)
	dfs(leaf, root.Right)
}

func Lc510LeafSimilar(root1, root2 *Structs.TreeNode) bool {
	leaf1, leaf2 := make([]int, 0), make([]int, 0)
	dfs(&leaf1, root1)
	dfs(&leaf2, root2)
	if len(leaf1) == len(leaf2) {
		for i, num := range leaf2 {
			if num != leaf1[i] {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
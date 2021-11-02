package D2021_05

import "Leetcode/Structs"

/*
993. 二叉树的堂兄弟节点
在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。
*/

/*
思路：
按层次遍历数，设置并维护层次标志符和兄弟标志符，同一层且不是兄弟即满足要求
 */

func Lc517IsCousins(root *Structs.TreeNode, x, y int) bool {
	if root == nil {
		return false
	}
	treeNodeArray := make([]*Structs.TreeNode, 0)
	treeNodeArray = append(treeNodeArray, root)
	brotherNode := new(Structs.TreeNode)
	treeNodeArray = append(treeNodeArray, brotherNode)
	brotherFlag := false
	levelNode := new(Structs.TreeNode)
	treeNodeArray = append(treeNodeArray, levelNode)
	findFlag := false
	tar := x
	for len(treeNodeArray) != 1 {
		if treeNodeArray[0] == nil {
			treeNodeArray = treeNodeArray[1:]
		} else if treeNodeArray[0] == brotherNode {
			brotherFlag = false
			treeNodeArray = treeNodeArray[1:]
		} else if treeNodeArray[0] == levelNode {
			if findFlag {
				return false
			}
			treeNodeArray = append(treeNodeArray, levelNode)
			treeNodeArray = treeNodeArray[1:]
		} else {
			if findFlag && !brotherFlag && treeNodeArray[0].Val == tar {
				return true
			}
			if treeNodeArray[0].Val == x {
				tar = y
				findFlag = true
				brotherFlag = true
			} else if treeNodeArray[0].Val == y {
				findFlag = true
				brotherFlag = true
			}
			treeNodeArray = append(treeNodeArray, treeNodeArray[0].Left)
			treeNodeArray = append(treeNodeArray, treeNodeArray[0].Right)
			treeNodeArray = append(treeNodeArray, brotherNode)
			treeNodeArray = treeNodeArray[1:]
		}
	}
	return false
}
package D2021_11

/*
559. N 叉树的最大深度
给定一个 N 叉树，找到其最大深度。
最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
 */

type Lc1121Node struct {
	Val int
	Children []*Lc1121Node
}

func Lc1121maxDepth(root *Lc1121Node) int {
	if len(root.Children) == 0 {
		return 1
	}
	maxLen := 0
	for _, child := range root.Children {
		tmpLen := Lc1121maxDepth(child)
		if maxLen < tmpLen {
			maxLen = tmpLen
		}
	}
	return maxLen + 1
}

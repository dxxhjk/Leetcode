package structs

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return new(TreeNode)
    }
    treeNodeArray := make([]*TreeNode, 0)
    for i := 0; i < len(nums); i++ {
        tempNode := new(TreeNode)
        tempNode.Val = nums[i]
        treeNodeArray = append(treeNodeArray, tempNode)
    }
    for i := 0; i < len(nums) / 2; i++ {
        leftIndex, rightIndex := i * 2 + 1, i * 2 + 2
        if leftIndex < len(nums) {
            treeNodeArray[i].Left = treeNodeArray[leftIndex]
        }
        if rightIndex < len(nums) {
            treeNodeArray[i].Right = treeNodeArray[rightIndex]
        }
    }
    return treeNodeArray[0]
}
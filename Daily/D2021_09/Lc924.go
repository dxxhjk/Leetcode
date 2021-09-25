package D2021_09

/*
430. 扁平化多级双向链表
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。
给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。
*/

type node struct {
	Val   int
	Prev  *node
	Next  *node
	Child *node
}

func Lc924Flatten(root *node) *node {
	s := make([]*node, 0)
	p := root
	if p == nil {
		return nil
	}
	for p.Next != nil || p.Child != nil || len(s) != 0 {
		if p.Child != nil {
			s = append(s, p.Next)
			p.Next = p.Child
			p.Child.Prev = p
			p.Child = nil
			p = p.Next
		} else if p.Next != nil {
			p = p.Next
		} else {
			p.Next = s[len(s)-1]
			if p.Next != nil {
				p.Next.Prev = p
				p = p.Next
			}
			s = s[:len(s)-1]
		}
	}
	return root
}

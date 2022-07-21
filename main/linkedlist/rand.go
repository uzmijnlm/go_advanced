package linkedlist

type NodeWithRand struct {
	Next  *NodeWithRand
	Value int
	Rand  *NodeWithRand
}

func NewNodeWithRand(val int) *NodeWithRand {
	return &NodeWithRand{Value: val}
}

func CopyNodeWithRand(node *NodeWithRand) *NodeWithRand {
	if node == nil {
		return nil
	}
	// 1.将新节点插入到老节点之间
	cur := node
	for cur != nil {
		next := cur.Next
		newNode := NewNodeWithRand(cur.Value)
		cur.Next = newNode
		newNode.Next = next
		cur = next
	}

	// 2.给新节点设置rand
	cur = node
	for cur != nil {
		oldNext := cur.Next.Next
		oldRand := cur.Rand
		newNode := cur.Next
		newNode.Rand = oldRand.Next
		cur = oldNext
	}

	// 3.分离新旧节点
	cur = node
	var result *NodeWithRand
	for cur != nil && cur.Next != nil {
		next := cur.Next
		if result == nil {
			result = next
		}
		cur.Next = cur.Next.Next
		cur = next
	}

	return result
}

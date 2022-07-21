package linkedlist

func IsPalindrome(node *Node) bool {
	if node == nil {
		return false
	}

	if node.Next == nil {
		return true
	}

	if node.Next.Next == nil {
		return node.Value == node.Next.Value
	}

	var slow *Node
	var fast *Node
	var prevSlow *Node

	// 用快慢指针，使慢指针指向中间右侧的节点
	slow = node
	fast = node
	for fast != nil && fast.Next != nil {
		prevSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	var isOdd bool
	if fast == nil { // 节点个数为偶数，slow指向中间右侧的节点
		isOdd = false
	} else { // 节点个数为奇数，slow指向中间节点
		// slow本来就指向中间右侧的节点，do nothing
		prevSlow = slow
		slow = slow.Next // 让其指向中间右侧的节点
		isOdd = true
	}

	prevSlow.Next = nil // 先切断前后两部分，以免翻转时对前半部分产生影响

	mid := slow // 提前备份中间右侧的节点

	// 翻转右半部分
	tail := ReverseSingly(slow)
	right := tail
	left := node

	isPalindrome := true
	for right != mid {
		if right.Value != left.Value {
			isPalindrome = false
		}
		right = right.Next
		left = left.Next
	}

	if right.Value != left.Value {
		isPalindrome = false
	}

	ReverseSingly(tail)

	if isOdd {
		left.Next.Next = mid
	} else {
		left.Next = mid
	}

	return isPalindrome
}

func SplitThreeParts(node *Node, num int) *Node {
	if node == nil {
		return nil
	}

	var lowHead *Node
	var lowTail *Node
	var midHead *Node
	var midTail *Node
	var highHead *Node
	var highTail *Node

	next := node.Next // 备份下一个节点
	for node != nil {
		node.Next = nil // 先切断这个节点和下一个节点的联系

		if node.Value < num {
			if lowHead == nil {
				lowHead = node
				lowTail = node
			} else {
				lowTail.Next = node
				lowTail = node
			}
		} else if node.Value == num {
			if midHead == nil {
				midHead = node
				midTail = node
			} else {
				midTail.Next = node
				midTail = node
			}
		} else {
			if highHead == nil {
				highHead = node
				highTail = node
			} else {
				highTail.Next = node
				highTail = node
			}
		}

		node = next      // 将节点指向下一个节点
		if node != nil { // 更新next的备份
			next = node.Next
		}
	}

	if lowTail != nil {
		if midHead != nil {
			lowTail.Next = midHead
		} else {
			if highHead != nil {
				lowTail.Next = highHead
			}
		}
	} else {
		if midTail != nil {
			if highHead != nil {
				midTail.Next = highHead
			}
		}
	}

	if lowHead != nil {
		return lowHead
	} else {
		if midHead != nil {
			return midHead
		} else {
			return highHead
		}
	}
}

func FindCrossPoint(node1 *Node, node2 *Node) *Node {
	hasLoop1 := HasLoop(node1)
	hasLoop2 := HasLoop(node2)

	if !hasLoop1 && !hasLoop2 { // 1.都无环，则先算出length1和length2，让长的先走|length1-length2|，然后一起走
		node, done := findCrossPointForBothNoLoop(node1, node2)
		if done {
			return node
		}
		return nil
	} else if (hasLoop1 && !hasLoop2) || (!hasLoop1 && hasLoop2) { // 2.一个有环一个无环，必然不相交
		return nil
	} else { // 3.两个都有环，分为3种情况
		// 分别找到入环节点
		entry1 := FindEntryForLoop(node1)
		entry2 := FindEntryForLoop(node2)
		if entry1 == entry2 { // 3.1 同一节点如环，思路同求两个无环链表交点
			node, done := findCrossPointForSameEntry(node1, node2, entry1)
			if done {
				return node
			}
			return nil
		} else {
			cur1 := entry1.Next
			for cur1 != entry1 {
				if cur1 == entry2 { // 3.2 不同节点如环
					return entry2
				}
				cur1 = cur1.Next
			}
			// 3.3 分别成环
			return nil
		}
	}
}

func findCrossPointForSameEntry(node1 *Node, node2 *Node, entry1 *Node) (*Node, bool) {
	cur1 := node1
	cur2 := node2
	size1 := 0
	size2 := 0
	for cur1 != entry1 {
		size1 += 1
		cur1 = cur1.Next
	}
	for cur2 != entry1 {
		size2 += 1
		cur2 = cur2.Next
	}
	cur1 = node1
	cur2 = node2
	if size1 >= size2 {
		diff := size1 - size2
		for i := 0; i < diff; i++ {
			cur1 = cur1.Next
		}
		for cur1 != cur2 {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
		return cur1, true
	} else {
		diff := size2 - size1
		for i := 0; i < diff; i++ {
			cur2 = cur2.Next
		}
		for cur1 != cur2 {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
		return cur1, true
	}
}

func findCrossPointForBothNoLoop(node1 *Node, node2 *Node) (*Node, bool) {
	size1 := 0
	size2 := 0
	cur1 := node1
	cur2 := node2

	for cur1 != nil {
		size1 += 1
		cur1 = cur1.Next
	}
	for cur2 != nil {
		size2 += 1
		cur2 = cur2.Next
	}

	if size1 >= size2 {
		cur1 = node1
		cur2 = node2
		diff := size1 - size2
		for i := 0; i < diff; i++ {
			cur1 = cur1.Next
		}
		for cur1 != nil && cur2 != nil {
			if cur1 == cur2 {
				return cur1, true
			} else {
				cur1 = cur1.Next
				cur2 = cur2.Next
			}
		}
	} else {
		cur1 = node1
		cur2 = node2
		diff := size2 - size1
		for i := 0; i < diff; i++ {
			cur2 = cur2.Next
		}
		for cur1 != nil && cur2 != nil {
			if cur1 == cur2 {
				return cur1, true
			} else {
				cur1 = cur1.Next
				cur2 = cur2.Next
			}
		}
	}
	return nil, false
}

func AddLinkedList(node1 *Node, node2 *Node) *Node {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	reversedNode1 := ReverseSingly(node1)
	cur1 := reversedNode1
	reversedNode2 := ReverseSingly(node2)
	cur2 := reversedNode2

	prevHead := NewNode(-1)
	cur := prevHead

	addOne := 0
	for cur1 != nil && cur2 != nil {
		val1 := cur1.Value
		val2 := cur2.Value
		sum := val1 + val2 + addOne
		if sum >= 10 {
			sum = sum - 10
			addOne = 1
		} else {
			addOne = 0
		}
		cur.Next = NewNode(sum)
		cur = cur.Next
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	for cur1 != nil {
		val := cur1.Value
		sum := val + addOne
		if sum >= 10 {
			sum = sum - 10
			addOne = 1
		} else {
			addOne = 0
		}
		cur.Next = NewNode(sum)
		cur = cur.Next
		cur1 = cur1.Next
	}

	for cur2 != nil {
		val := cur2.Value
		sum := val + addOne
		if sum >= 10 {
			sum = sum - 10
			addOne = 1
		} else {
			addOne = 0
		}
		cur.Next = NewNode(sum)
		cur = cur.Next
		cur2 = cur2.Next
	}

	if addOne == 1 {
		cur.Next = NewNode(1)
	}

	result := ReverseSingly(prevHead.Next)
	for result.Value == 0 {
		result = result.Next
	}
	return result
}

package linkedlist

import (
	"fmt"
	"math"
)

type Node struct {
	Next  *Node
	Value int
}

type DoubleNode struct {
	Next  *DoubleNode
	Prev  *DoubleNode
	Value int
}

func NewNode(val int) *Node {
	return &Node{
		Value: val,
	}
}

func NewDoubleNode(val int) *DoubleNode {
	return &DoubleNode{
		Value: val,
	}
}

func PrintLinkedList(node *Node) {
	fmt.Print("Linkedlist: ")
	for node != nil {
		fmt.Printf("%d->", node.Value)
		node = node.Next
	}
	fmt.Println("nil")
}

func PrintDoubleLinkedList(node *DoubleNode) {
	fmt.Print("Linkedlist: ")
	for node != nil {
		fmt.Printf("%d<->", node.Value)
		node = node.Next
	}
	fmt.Println("nil")
}

func CopyLinkedList(node *Node) *Node {
	var head *Node
	var newNode = NewNode(node.Value)
	head = newNode
	for node != nil && node.Next != nil {
		newNode.Next = NewNode(node.Next.Value)
		newNode = newNode.Next
		node = node.Next
	}
	return head
}

func CopyDoubleLinkedList(node *DoubleNode) *DoubleNode {
	var head *DoubleNode
	var newNode = NewDoubleNode(node.Value)
	head = newNode
	for node != nil && node.Next != nil {
		newNode.Next = NewDoubleNode(node.Next.Value)
		newNode.Next.Prev = newNode
		newNode = newNode.Next
		node = node.Next
	}
	return head
}

// ReverseSingly 核心思路是提前备份next节点、记录上一个遍历到的节点prev
func ReverseSingly(node *Node) *Node {
	var prev *Node
	var next *Node
	for node != nil {
		next = node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}

// ReverseDouble 在反转单链表的基础上多修改一个prev即可
func ReverseDouble(node *DoubleNode) *DoubleNode {
	var prev *DoubleNode
	var next *DoubleNode
	for node != nil {
		next = node.Next
		node.Next = prev
		node.Prev = next
		prev = node
		node = next
	}
	return prev
}

func GetCommonStart(node1 *Node, node2 *Node) *Node {
	if node1 == nil || node2 == nil {
		return nil
	}

	cur1 := node1
	size1 := 1
	cur2 := node2
	size2 := 1
	for cur1.Next != nil {
		cur1 = cur1.Next
		size1 += 1
	}
	for cur2.Next != nil {
		cur2 = cur2.Next
		size2 += 1
	}

	if cur1 != cur2 {
		return nil
	}

	cur1 = node1
	cur2 = node2
	var diff int
	if size1 >= size2 {
		diff = size1 - size2
		for i := 0; i < diff; i++ {
			cur1 = cur1.Next
		}
	} else {
		diff = size2 - size1
		for i := 0; i < diff; i++ {
			cur2 = cur2.Next
		}
	}

	for {
		if cur1 == cur2 {
			return cur1
		} else {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
	}
}

func HasLoop(node *Node) bool {
	if node == nil {
		return false
	}

	if node.Next == nil {
		return false
	}

	if node.Next.Next == nil {
		return false
	}

	slow := node.Next
	fast := node.Next.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func FindEntryForLoop(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.Next == nil {
		return nil
	}

	if node.Next.Next == nil {
		return nil
	}

	slow := node.Next
	fast := node.Next.Next
	hasLoop := false
	for fast != nil && fast.Next != nil {
		if slow == fast {
			hasLoop = true
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	if hasLoop {
		fast = node
		for fast != slow {
			fast = fast.Next
			slow = slow.Next
		}
		return slow
	} else {
		return nil
	}

}

// RemoveLastKthNodeSingly 删除单链表倒数第k个元素
// 从头开始遍历，每走一步（包括头节点）k的值减1：
// a.如果走到最后k的值等于0，则倒数第k个就是头节点，返回head.next
// b.如果走到最后k的值大于0，则说明没有倒数第k个节点，返回head
// c.如果走到最后k的值小于0，则从头再走一次，每走一步（包括头节点）k的值加1，k==0时结束，这个节点就是要删除的节点的前一个节点
func RemoveLastKthNodeSingly(node *Node, k int) *Node {
	lastK := k
	if k <= 0 || node == nil {
		return node
	}

	cur := node
	for cur != nil {
		lastK -= 1
		cur = cur.Next
	}
	if lastK == 0 {
		return node.Next
	} else if lastK >= 0 {
		return node
	} else {
		cur := node
		for cur != nil {
			lastK += 1
			if lastK == 0 {
				break
			}
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
		return node
	}
}

func RemoveLastKthNodeDouble(node *DoubleNode, k int) *DoubleNode {
	lastK := k
	if k <= 0 || node == nil {
		return node
	}

	cur := node
	for cur != nil {
		lastK -= 1
		cur = cur.Next
	}
	if lastK == 0 {
		next := node.Next
		next.Prev = nil
		return next
	} else if lastK >= 0 {
		return node
	} else {
		cur := node
		for cur != nil {
			lastK += 1
			if lastK == 0 {
				break
			}
			cur = cur.Next
		}
		newNext := cur.Next.Next
		cur.Next = newNext
		if newNext != nil {
			newNext.Prev = cur
		}
		return node
	}
}

func RemoveMidNode(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Next == nil {
		return nil
	}
	if node.Next.Next == nil {
		return node.Next
	}

	slow := node
	fast := node.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return node
}

// RemoveByRatio 删除位于a/b处的节点
// 例如链表一共7个节点，a=5，b=7，则要删除的是第 7*(5/7)=5 个节点
// 如果不能整除，如一共8个节点，a=5，b=7，则要删除的是 8*(5/7)=5.7 个节点，此时向上取整，应删除第6个节点
func RemoveByRatio(node *Node, a int, b int) *Node {
	cur := node
	size := 0
	for cur != nil {
		size += 1
		cur = cur.Next
	}

	if size == 0 || a > b {
		return node
	}

	if size == 1 {
		return nil
	}

	removeIndex := int(math.Ceil(float64(size)*float64(a)/float64(b))) - 1

	if removeIndex == 0 {
		return node.Next
	}

	prevRemoveIndex := removeIndex - 1

	index := -1
	cur = node
	for {
		index += 1
		if index == prevRemoveIndex {
			break
		}
		cur = cur.Next
	}

	cur.Next = cur.Next.Next
	return node
}

package linkedlist

import (
	"fmt"
	"github.com/uzmijnlm/go_advanced/main/basic"
	"github.com/uzmijnlm/go_advanced/main/linkedlist"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestReverseSingly(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1

		var head *linkedlist.Node
		var node = new(linkedlist.Node)
		head = node
		for i := 0; i < length; i++ {
			node.Value = rand.Intn(maxNum)
			node.Next = new(linkedlist.Node)
			node = node.Next
		}

		head1 := linkedlist.CopyLinkedList(head)
		head2 := linkedlist.CopyLinkedList(head)

		reversedHead1 := linkedlist.ReverseSingly(head1)
		reversedHead2 := reverseSingly(head2)

		for reversedHead1 != nil && reversedHead2 != nil {
			if reversedHead1.Value != reversedHead2.Value {
				t.Fail()
			}
			reversedHead1 = reversedHead1.Next
			reversedHead2 = reversedHead2.Next
		}

		if reversedHead1 != nil || reversedHead2 != nil {
			t.Fail()
		}
	}
}

func reverseSingly(head2 *linkedlist.Node) *linkedlist.Node {
	stack := basic.NewStack()
	head := head2
	for head != nil {
		stack.Push(head)
		head = head.Next
	}

	var newHead *linkedlist.Node
	var newNode *linkedlist.Node
	node := stack.Pop().(*linkedlist.Node)
	newNode = linkedlist.NewNode(node.Value)
	newHead = newNode
	for !stack.IsEmpty() {
		next := stack.Pop().(*linkedlist.Node)
		newNode.Next = linkedlist.NewNode(next.Value)
		newNode = newNode.Next
	}
	return newHead
}

func TestReverseDouble(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1

		var head *linkedlist.DoubleNode
		var node = linkedlist.NewDoubleNode(rand.Intn(maxNum))
		head = node
		for i := 0; i < length; i++ {
			node.Value = rand.Intn(maxNum)
			next := linkedlist.NewDoubleNode(rand.Intn(maxNum))
			node.Next = next
			next.Prev = node
			node = next
		}
		head1 := linkedlist.CopyDoubleLinkedList(head)
		head2 := linkedlist.CopyDoubleLinkedList(head)

		reversedHead1 := linkedlist.ReverseDouble(head1)
		reversedHead2 := reverseDouble(head2)

		for reversedHead1 != nil && reversedHead2 != nil {
			if reversedHead1.Value != reversedHead2.Value {
				t.Fail()
			}
			reversedHead1 = reversedHead1.Next
			reversedHead2 = reversedHead2.Next
			if reversedHead1 != nil {
				prev1 := reversedHead1.Prev
				prev2 := reversedHead2.Prev
				if prev1.Value != prev2.Value {
					t.Fail()
				}
			}

		}

		if reversedHead1 != nil || reversedHead2 != nil {
			t.Fail()
		}
	}
}

func reverseDouble(head2 *linkedlist.DoubleNode) *linkedlist.DoubleNode {
	stack := basic.NewStack()
	head := head2
	for head != nil {
		stack.Push(head)
		head = head.Next
	}

	var newHead *linkedlist.DoubleNode
	var newNode *linkedlist.DoubleNode
	node := stack.Pop().(*linkedlist.DoubleNode)
	newNode = linkedlist.NewDoubleNode(node.Value)
	newHead = newNode
	for !stack.IsEmpty() {
		next := stack.Pop().(*linkedlist.DoubleNode)
		newNext := linkedlist.NewDoubleNode(next.Value)
		newNode.Next = newNext
		newNext.Prev = newNode
		newNode = newNext
	}
	return newHead
}

func TestGetCommonStart(t *testing.T) {
	for i := 0; i < 1000; i++ {
		var node1Seg = new(linkedlist.Node)
		var node2Seg = new(linkedlist.Node)

		rand.Seed(time.Now().UnixNano())
		length1 := rand.Intn(100-10) + 10
		length2 := rand.Intn(100-10) + 10

		cur1 := node1Seg
		cur2 := node2Seg

		for i := 0; i < length1; i++ {
			newNode := new(linkedlist.Node)
			newNode.Value = i
			cur1.Next = newNode
			cur1 = cur1.Next
		}

		for i := 0; i < length2; i++ {
			newNode := new(linkedlist.Node)
			newNode.Value = i
			cur2.Next = newNode
			cur2 = cur2.Next
		}

		meet := rand.Intn(2)

		if meet == 0 { // 没有公共部分
			result := linkedlist.GetCommonStart(node1Seg, node2Seg)
			if result != nil {
				fmt.Println("没有公共部分")
				linkedlist.PrintLinkedList(node1Seg)
				linkedlist.PrintLinkedList(node2Seg)
				t.Fail()
			}
		} else { // 构建公共部分
			var commonStart = new(linkedlist.Node)
			commonStart.Value = -1
			lengthCommon := rand.Intn(100-10) + 10
			cur := commonStart

			for i := 0; i < lengthCommon; i++ {
				cur.Next = new(linkedlist.Node)
				cur = cur.Next
			}

			cur1.Next = commonStart
			cur2.Next = commonStart

			result := linkedlist.GetCommonStart(node1Seg, node2Seg)
			if result != commonStart {
				fmt.Println("有公共部分")
				linkedlist.PrintLinkedList(node1Seg)
				linkedlist.PrintLinkedList(node2Seg)
				t.Fail()
			}

		}

	}

}

func TestHasLoop(t *testing.T) {
	var node *linkedlist.Node
	if linkedlist.HasLoop(node) != false {
		t.Fail()
	}

	node = new(linkedlist.Node)
	cur := node
	cur.Next = new(linkedlist.Node)
	if linkedlist.HasLoop(node) != false {
		t.Fail()
	}

	cur = cur.Next
	cur.Next = new(linkedlist.Node)
	if linkedlist.HasLoop(node) != false {
		t.Fail()
	}

	cur = cur.Next
	cur.Next = new(linkedlist.Node)
	if linkedlist.HasLoop(node) != false {
		t.Fail()
	}

	entry := cur
	cur = cur.Next
	cur.Next = entry
	if linkedlist.HasLoop(node) != true {
		t.Fail()
	}

	cur.Next = new(linkedlist.Node)
	cur = cur.Next
	cur.Next = entry
	if linkedlist.HasLoop(node) != true {
		t.Fail()
	}

}

func TestFindEntryForLoop(t *testing.T) {
	var node *linkedlist.Node
	if linkedlist.FindEntryForLoop(node) != nil {
		t.Fail()
	}

	node = linkedlist.NewNode(0)
	cur := node
	cur.Next = linkedlist.NewNode(1)
	if linkedlist.FindEntryForLoop(node) != nil {
		t.Fail()
	}

	cur = cur.Next
	cur.Next = linkedlist.NewNode(2)
	if linkedlist.FindEntryForLoop(node) != nil {
		t.Fail()
	}

	cur = cur.Next
	cur.Next = linkedlist.NewNode(3)
	if linkedlist.FindEntryForLoop(node) != nil {
		t.Fail()
	}

	entry := cur
	cur = cur.Next
	cur.Next = entry
	if linkedlist.FindEntryForLoop(node) != entry {
		t.Fail()
	}

	cur.Next = linkedlist.NewNode(4)
	cur = cur.Next
	cur.Next = entry
	if linkedlist.FindEntryForLoop(node) != entry {
		t.Fail()
	}
}

func TestRemoveLastKthNodeSingly(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1

		head := linkedlist.NewNode(rand.Intn(maxNum))
		cur := head

		for j := 0; j < length; j++ {
			cur.Next = linkedlist.NewNode(rand.Intn(maxNum))
			cur = cur.Next
		}

		head1 := linkedlist.CopyLinkedList(head)
		head2 := linkedlist.CopyLinkedList(head)

		k := rand.Intn(length) + 1 // [1, length]随机

		res1 := linkedlist.RemoveLastKthNodeSingly(head1, k)
		res2 := removeLastKthNodeSingly(head2, k)

		for res1 != nil {
			if res1.Value != res2.Value {
				t.Fail()
			}
			res1 = res1.Next
			res2 = res2.Next
		}
		if res2 != nil {
			t.Fail()
		}
	}
}

func removeLastKthNodeSingly(head *linkedlist.Node, k int) *linkedlist.Node {
	size := 0
	cur := head
	for cur != nil {
		size += 1
		cur = cur.Next
	}
	if size < k {
		return head
	} else if size == k {
		return head.Next
	} else {
		n := size - k // 正数第n个就是要删除的前一个节点
		count := 0
		cur := head
		for {
			count += 1
			if count == n {
				break
			}
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
		return head
	}
}

func TestRemoveLastKthNodeDouble(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1

		head := linkedlist.NewDoubleNode(rand.Intn(maxNum))
		cur := head

		for j := 0; j < length; j++ {
			cur.Next = linkedlist.NewDoubleNode(rand.Intn(maxNum))
			cur = cur.Next
		}

		head1 := linkedlist.CopyDoubleLinkedList(head)
		head2 := linkedlist.CopyDoubleLinkedList(head)

		k := rand.Intn(length) + 1 // [1, length]随机

		res1 := linkedlist.RemoveLastKthNodeDouble(head1, k)
		res2 := removeLastKthNodeDouble(head2, k)

		for res1 != nil {
			if res1.Value != res2.Value {
				t.Fail()
			}
			res1 = res1.Next
			res2 = res2.Next
			if res1 != nil {
				prev1 := res1.Prev
				prev2 := res2.Prev
				if prev1.Value != prev2.Value {
					t.Fail()
				}
			}
		}
		if res2 != nil {
			t.Fail()
		}
	}
}

func removeLastKthNodeDouble(head *linkedlist.DoubleNode, k int) *linkedlist.DoubleNode {
	size := 0
	cur := head
	for cur != nil {
		size += 1
		cur = cur.Next
	}
	if size < k {
		return head
	} else if size == k {
		next := head.Next
		next.Prev = nil
		return next
	} else {
		n := size - k // 正数第n个就是要删除的前一个节点
		count := 0
		cur := head
		for {
			count += 1
			if count == n {
				break
			}
			cur = cur.Next
		}
		newNext := cur.Next.Next
		cur.Next = newNext
		if newNext != nil {
			newNext.Prev = cur
		}
		return head
	}
}

func TestRemoveMidNode(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1

		head := linkedlist.NewNode(rand.Intn(maxNum))
		cur := head

		for j := 0; j < length; j++ {
			cur.Next = linkedlist.NewNode(rand.Intn(maxNum))
			cur = cur.Next
		}

		head1 := linkedlist.CopyLinkedList(head)
		head2 := linkedlist.CopyLinkedList(head)

		res1 := linkedlist.RemoveMidNode(head1)
		res2 := removeMidNode(head2)

		for res1 != nil {
			if res1.Value != res2.Value {
				t.Fail()
			}
			res1 = res1.Next
			res2 = res2.Next
		}
		if res2 != nil {
			t.Fail()
		}

	}
}

func removeMidNode(head *linkedlist.Node) *linkedlist.Node {
	if head == nil {
		return nil
	}
	size := 0
	cur := head
	for cur != nil {
		size += 1
		cur = cur.Next
	}
	if size == 1 {
		return nil
	}
	if size == 2 {
		return head.Next
	}

	midIndex := (size - 1) / 2

	prevMid := midIndex - 1

	index := -1
	cur = head
	for {
		index += 1
		if index == prevMid {
			break
		}
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return head
}

func TestRemoveByRatio(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1

		head := linkedlist.NewNode(rand.Intn(maxNum))
		cur := head

		for j := 0; j < length; j++ {
			cur.Next = linkedlist.NewNode(rand.Intn(maxNum))
			cur = cur.Next
		}

		head1 := linkedlist.CopyLinkedList(head)
		head2 := linkedlist.CopyLinkedList(head)

		a := rand.Intn(100) + 1
		b := rand.Intn(100) + 1

		res1 := linkedlist.RemoveByRatio(head1, a, b)
		res2 := removeByRatio(head2, a, b)

		for res1 != nil {
			if res1.Value != res2.Value {
				t.Fail()
			}
			res1 = res1.Next
			res2 = res2.Next
		}
		if res2 != nil {
			t.Fail()
		}

	}
}

func removeByRatio(head *linkedlist.Node, a int, b int) *linkedlist.Node {
	if a < 1 || a > b {
		return head
	}
	n := 0
	cur := head
	for cur != nil {
		n += 1
		cur = cur.Next
	}
	n = int(math.Ceil(float64(a*n) / float64(b)))
	if n == 1 {
		head = head.Next
	}
	if n > 1 {
		cur = head
		for {
			n -= 1
			if n == 1 {
				break
			}
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
	}
	return head
}

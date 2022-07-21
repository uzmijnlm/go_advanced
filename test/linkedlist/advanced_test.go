package linkedlist

import (
	"github.com/uzmijnlm/go_advanced/main/linkedlist"
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	var node *linkedlist.Node
	if linkedlist.IsPalindrome(node) != false {
		t.Fail()
	}
	node = linkedlist.NewNode(-1)
	if linkedlist.IsPalindrome(node) != true {
		t.Fail()
	}
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		max := rand.Intn(100) + 10
		isPalindrome := rand.Intn(2)

		if isPalindrome == 0 { // 不构造回文
			firstNodeVal := rand.Intn(max)
			var node = linkedlist.NewNode(firstNodeVal)
			head := node
			for i := 1; i < length; i++ {
				val := rand.Intn(max)
				if i == length-1 {
					for val == firstNodeVal {
						val = rand.Intn(max)
					}
				}
				node.Next = linkedlist.NewNode(val)
				node = node.Next
			}

			if linkedlist.IsPalindrome(head) != false {
				t.Fail()
			}
		} else { // 构造回文
			var node = linkedlist.NewNode(0)
			head := node
			for i := 1; i < length; i++ {
				if i < length/2 {
					node.Next = linkedlist.NewNode(i)
				} else {
					node.Next = linkedlist.NewNode(length - i - 1)
				}
				node = node.Next
			}

			if linkedlist.IsPalindrome(head) != true {
				t.Fail()
			}

		}

	}

}

func TestSplitThreeParts(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(1) + 10
		max := rand.Intn(100) + 1

		num := rand.Intn(max)

		node := linkedlist.NewNode(rand.Intn(max))
		head := node
		for i := 0; i < length; i++ {
			node.Next = linkedlist.NewNode(rand.Intn(max))
			node = node.Next
		}

		result := linkedlist.SplitThreeParts(head, num)

		meetEqual := false
		meetLarger := false

		for result != nil {
			if !meetEqual && !meetLarger {
				if result.Value == num {
					meetEqual = true
				} else if result.Value > num {
					meetLarger = true
				}
				result = result.Next
				continue
			}
			if meetEqual {
				if result.Value > num {
					meetLarger = true
				} else if result.Value < num {
					t.Fail()
				}
				result = result.Next
				continue
			}
			if meetLarger {
				if result.Value <= num {
					t.Fail()
				}
				result = result.Next
				continue
			}
		}

	}

}

func TestFindCrossPoint(t *testing.T) {
	// 1：两个无环链表
	testBothNoLoop(t)

	// 2.一个有环一个无环
	testOnlyOneHasLoop(t)

	// 3.两个有环
	testBothHasLoop(t)
}

func testBothHasLoop(t *testing.T) {
	// 3.1 同一节点入环
	testSameEntryPoint(t)

	// 3.2 不同点如环
	testNotSameEntryPoint(t)

	// 3.3 分别成环
	testSeparateLoop(t)

}

func testSeparateLoop(t *testing.T) {
	for i := 0; i < 1000; i++ {
		length1 := rand.Intn(100-10) + 10
		length2 := rand.Intn(100-10) + 10

		node1 := new(linkedlist.Node)
		node1.Value = 0
		node2 := new(linkedlist.Node)
		node2.Value = 0
		cur1 := node1
		cur2 := node2

		for i := 0; i < length1; i++ {
			cur1.Next = new(linkedlist.Node)
			cur1 = cur1.Next
			cur1.Value = i + 1
		}
		for i := 0; i < length2; i++ {
			cur2.Next = new(linkedlist.Node)
			cur2 = cur2.Next
			cur2.Value = i + 1
		}

		entry1 := new(linkedlist.Node)
		entry1.Value = -1
		cur1.Next = entry1
		cur3 := entry1
		length3 := rand.Intn(100-10) + 10
		for i := 0; i < length3; i++ {
			cur3.Next = new(linkedlist.Node)
			cur3 = cur3.Next
			cur3.Value = i
		}
		cur3.Next = entry1

		entry2 := new(linkedlist.Node)
		entry2.Value = -1
		cur2.Next = entry2
		cur4 := entry2
		length4 := rand.Intn(100-10) + 10
		for i := 0; i < length4; i++ {
			cur4.Next = new(linkedlist.Node)
			cur4 = cur4.Next
			cur4.Value = i
		}
		cur4.Next = entry2

		if linkedlist.FindEntryForLoop(node1) != entry1 {
			t.Fail()
		}
		if linkedlist.FindEntryForLoop(node2) != entry2 {
			t.Fail()
		}
		if linkedlist.FindCrossPoint(node1, node2) != nil {
			t.Fail()
		}
	}
}

func testNotSameEntryPoint(t *testing.T) {
	for i := 0; i < 1000; i++ {
		length1 := rand.Intn(100-10) + 10
		length2 := rand.Intn(100-10) + 10

		node1 := new(linkedlist.Node)
		node1.Value = 0
		node2 := new(linkedlist.Node)
		node2.Value = 0
		cur1 := node1
		cur2 := node2

		for i := 0; i < length1; i++ {
			cur1.Next = new(linkedlist.Node)
			cur1 = cur1.Next
			cur1.Value = i + 1
		}
		for i := 0; i < length2; i++ {
			cur2.Next = new(linkedlist.Node)
			cur2 = cur2.Next
			cur2.Value = i + 1
		}

		entry := new(linkedlist.Node)
		entry.Value = -1
		cur1.Next = entry
		cur3 := entry
		length3 := rand.Intn(100-10) + 10

		container := make([]*linkedlist.Node, 0)

		for i := 0; i < length3; i++ {
			cur3.Next = new(linkedlist.Node)
			cur3 = cur3.Next
			container = append(container, cur3)
			cur3.Value = i
		}
		cur3.Next = entry

		var crossPoint *linkedlist.Node
		randNum := rand.Intn(len(container))
		crossPoint = container[randNum]
		cur2.Next = crossPoint

		if linkedlist.FindEntryForLoop(node1) != entry {
			t.Fail()
		}
		if linkedlist.FindEntryForLoop(node2) != crossPoint {
			t.Fail()
		}
		if linkedlist.FindCrossPoint(node1, node2) != crossPoint && linkedlist.FindCrossPoint(node1, node2) != entry {
			t.Fail()
		}
	}
}

func testSameEntryPoint(t *testing.T) {
	for i := 0; i < 1000; i++ {
		length1 := rand.Intn(100-10) + 10
		length2 := rand.Intn(100-10) + 10

		node1 := new(linkedlist.Node)
		node1.Value = 0
		node2 := new(linkedlist.Node)
		node2.Value = 0
		cur1 := node1
		cur2 := node2

		container := make([]*linkedlist.Node, 0)
		container = append(container, cur1)
		for i := 0; i < length1; i++ {
			cur1.Next = new(linkedlist.Node)
			cur1 = cur1.Next
			container = append(container, cur1)
			cur1.Value = i + 1
		}
		for i := 0; i < length2; i++ {
			cur2.Next = new(linkedlist.Node)
			cur2 = cur2.Next
			cur2.Value = i + 1
		}

		entry := new(linkedlist.Node)
		entry.Value = -1
		cur1.Next = entry
		cur3 := entry
		length3 := rand.Intn(100-10) + 10
		for i := 0; i < length3; i++ {
			cur3.Next = new(linkedlist.Node)
			cur3 = cur3.Next
			cur3.Value = i
		}
		cur3.Next = entry

		var crossPoint *linkedlist.Node
		container = append(container, entry)
		randNum := rand.Intn(len(container))
		crossPoint = container[randNum]
		cur2.Next = crossPoint

		if linkedlist.FindEntryForLoop(node1) != entry {
			t.Fail()
		}
		if linkedlist.FindEntryForLoop(node2) != entry {
			t.Fail()
		}
		if linkedlist.FindCrossPoint(node1, node2) != crossPoint {
			t.Fail()
		}
	}
}

func testOnlyOneHasLoop(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length1 := rand.Intn(100-10) + 10
		length2 := rand.Intn(100-10) + 10

		node1 := new(linkedlist.Node)
		node2 := new(linkedlist.Node)
		cur1 := node1
		cur2 := node2
		for i := 0; i < length1; i++ {
			cur1.Next = new(linkedlist.Node)
			cur1 = cur1.Next
		}
		for i := 0; i < length2; i++ {
			cur2.Next = new(linkedlist.Node)
			cur2 = cur2.Next
		}

		entry := new(linkedlist.Node)
		cur1.Next = entry
		cur3 := entry
		length3 := rand.Intn(100-10) + 10
		for i := 0; i < length3; i++ {
			cur3.Next = new(linkedlist.Node)
			cur3 = cur3.Next
		}
		cur3.Next = entry

		if linkedlist.FindEntryForLoop(node1) != entry {
			t.Fail()
		}

		if linkedlist.FindCrossPoint(node1, node2) != nil {
			t.Fail()
		}

	}
}

func testBothNoLoop(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length1 := rand.Intn(100-10) + 10
		length2 := rand.Intn(100-10) + 10

		meet := rand.Intn(2)

		node1 := new(linkedlist.Node)
		node2 := new(linkedlist.Node)
		cur1 := node1
		cur2 := node2
		for i := 0; i < length1; i++ {
			cur1.Next = new(linkedlist.Node)
			cur1 = cur1.Next
		}
		for i := 0; i < length2; i++ {
			cur2.Next = new(linkedlist.Node)
			cur2 = cur2.Next
		}

		if meet == 0 { // 1.1 构造两个相交的链表
			crossPoint := new(linkedlist.Node)
			cur3 := crossPoint
			length3 := rand.Intn(100-10) + 10
			for i := 0; i < length3; i++ {
				cur3.Next = new(linkedlist.Node)
				cur3 = cur3.Next
			}

			cur1.Next = crossPoint
			cur2.Next = crossPoint

			if linkedlist.FindCrossPoint(node1, node2) != crossPoint {
				t.Fail()
			}
		} else { // 1.2 构造两个不相交的链表
			if linkedlist.FindCrossPoint(node1, node2) != nil {
				t.Fail()
			}
		}

	}
}

func TestAddLinkedList(t *testing.T) {
	// 1. 至少一个为nil
	head1 := linkedlist.NewNode(1)
	cur1 := head1
	cur1.Next = linkedlist.NewNode(2)
	if linkedlist.AddLinkedList(head1, nil) != head1 {
		t.Fail()
	}
	if linkedlist.AddLinkedList(nil, head1) != head1 {
		t.Fail()
	}
	if linkedlist.AddLinkedList(nil, nil) != nil {
		t.Fail()
	}

	// 2. 都不为nil
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length1 := rand.Intn(8) + 1
		length2 := rand.Intn(8) + 1

		head1 := linkedlist.NewNode(rand.Intn(10))
		head2 := linkedlist.NewNode(rand.Intn(10))

		cur1 := head1
		cur2 := head2
		for i := 0; i < length1; i++ {
			cur1.Next = linkedlist.NewNode(rand.Intn(10))
			cur1 = cur1.Next
		}
		for i := 0; i < length2; i++ {
			cur2.Next = linkedlist.NewNode(rand.Intn(10))
			cur2 = cur2.Next
		}
		copiedHead1 := linkedlist.CopyLinkedList(head1)
		copiedHead2 := linkedlist.CopyLinkedList(head2)

		res := linkedlist.AddLinkedList(head1, head2)
		expectedRes := addLinkedList(copiedHead1, copiedHead2)

		for res != nil {
			if res.Value != expectedRes.Value {
				t.Fail()
			}
			res = res.Next
			expectedRes = expectedRes.Next
		}
		if expectedRes != nil {
			t.Fail()
		}

	}

}

func addLinkedList(head1 *linkedlist.Node, head2 *linkedlist.Node) *linkedlist.Node {
	val1 := 0
	val2 := 0
	cur1 := head1
	cur2 := head2
	for cur1 != nil {
		val1 = val1*10 + cur1.Value
		cur1 = cur1.Next
	}
	for cur2 != nil {
		val2 = val2*10 + cur2.Value
		cur2 = cur2.Next
	}

	sum := val1 + val2
	if sum == 0 {
		return linkedlist.NewNode(0)
	}

	var prevNode *linkedlist.Node
	var node *linkedlist.Node
	for sum != 0 {
		val := sum % 10
		node = linkedlist.NewNode(val)
		sum = sum / 10
		node.Next = prevNode
		prevNode = node
	}
	return node
}

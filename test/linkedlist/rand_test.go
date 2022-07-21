package linkedlist

import (
	"github.com/uzmijnlm/go_advanced/main/linkedlist"
	"math/rand"
	"testing"
	"time"
)

func TestCopyNodeWithRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1

		var container []*linkedlist.NodeWithRand

		var head *linkedlist.NodeWithRand
		var node = new(linkedlist.NodeWithRand)
		container = append(container, node)
		head = node
		for i := 0; i < length; i++ {
			node.Value = rand.Intn(maxNum)
			node.Next = new(linkedlist.NodeWithRand)
			node = node.Next
			container = append(container, node)
		}

		cur := head
		for cur != nil {
			cur.Rand = container[rand.Intn(length)]
			cur = cur.Next
		}

		copied := linkedlist.CopyNodeWithRand(head)

		cur1 := head
		cur2 := copied

		for cur1 != nil && cur2 != nil {
			if cur1.Value != cur2.Value {
				t.Fail()
			}
			if cur1.Rand.Value != cur2.Rand.Value {
				t.Fail()
			}
			cur1 = cur1.Next
			cur2 = cur2.Next
		}

		if cur1 != nil || cur2 != nil {
			t.Fail()
		}

	}

}

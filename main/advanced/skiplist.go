package advanced

import (
	"math/rand"
	"time"
)

const ( // 参考Redis中的常数取值
	P        = 1 / 4.0
	MaxLevel = 32
)

type skipListNode struct {
	value int             // 节点值
	next  []*skipListNode // 节点在不同层的下一个节点
}

func newSkipListNode(value int, size int) *skipListNode { // size表示当前节点在跳表索引几层
	return &skipListNode{
		value: value,
		next:  make([]*skipListNode, size),
	}
}

// 随机一个层数
// 算法并不是产生一个普通的服从均匀分布的随机数，其计算过程如下：
// 1.首先，每个节点肯定都有第1层指针，即都在第1层链表里
// 2.如果一个节点有第i层（i>=1)指针，即节点已经在第1层到第i层链表中），那么它有第i+1层指针的概率为P。
// 3.节点最大的层数不允许超过一个最大值，记为MaxLevel
func randomLevel() int {
	level := 1
	rand.Seed(time.Now().UnixNano())
	for rand.Float64() < P && level < MaxLevel {
		level += 1
	}
	return level
}

// 从startNode节点开始，找到level层最后一个小于value的节点，如果都小于value，则返回最后一个节点
func findClosest(startNode *skipListNode, level int, value int) *skipListNode {
	cur := startNode
	for cur.next[level] != nil && value > cur.next[level].value {
		cur = cur.next[level]
	}
	return cur
}

type SkipList struct {
	head  *skipListNode
	level int
}

func Constructor() *SkipList {
	return &SkipList{
		head:  newSkipListNode(-1, MaxLevel),
		level: 1,
	}
}

func (skipList *SkipList) Add(num int) {
	level := randomLevel()
	newNode := newSkipListNode(num, level)
	cur := skipList.head
	for i := skipList.level - 1; i >= 0; i-- { // i表示层数索引
		cur = findClosest(cur, i, num)
		if i < level { // 如果新节点的随机层数小于等于当前跳表的总层数，则用链表的插入方式将新节点插入即可
			if cur.next[i] == nil {
				cur.next[i] = newNode
			} else {
				oldNext := cur.next[i]
				cur.next[i] = newNode
				newNode.next[i] = oldNext
			}
		}
	}
	if level > skipList.level { // 如果随机层数大于当前跳表总层数，那么超过的层数head直接指向新节点
		for i := skipList.level; i < level; i++ {
			skipList.head.next[i] = newNode
		}
		skipList.level = level
	}
}

func (skipList *SkipList) Search(target int) bool {
	cur := skipList.head
	for i := skipList.level - 1; i >= 0; i-- {
		cur = findClosest(cur, i, target)
		if cur.next[i] != nil && cur.next[i].value == target {
			return true
		}
	}
	return false
}

func (skipList *SkipList) Erase(num int) bool {
	deleted := false
	cur := skipList.head
	level := 0
	for i := skipList.level - 1; i >= 0; i-- {
		cur = findClosest(cur, i, num)
		if cur.next[i] != nil && cur.next[i].value == num {
			cur.next[i] = cur.next[i].next[i]
			deleted = true
		}
		if skipList.head.next[i] != nil {
			level += 1
		}
	}
	skipList.level = level
	return deleted
}

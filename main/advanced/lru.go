package advanced

/*
	LRU Cache
	在构造时确定大小，为K

	有如下两个功能：
		1.set(key, value)：将记录插入该结构
		2.get(key)：返回key对应的value值

	要求：
		1.set和get的时间复杂度为O(1)
		2.某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的
		3.当缓存的大小超过K时，移除最不经常使用的记录

	思路：双向链表+hash即可。用双向链表表示谁最先操作、谁最后操作

	难点：自己实现双向链表，支持如下功能：
		1.新加一个节点到尾部，即addNode方法
		2.把一个链表中的节点头尾节点重连，将其放到尾部，即moveNodeToTail方法
		3.删除头节点，即removeHead方法
*/

type LRUNode struct {
	next  *LRUNode
	last  *LRUNode
	value interface{}
}

func newLRUNode(value interface{}) *LRUNode {
	return &LRUNode{value: value}
}

type LRUDoubleLinkedList struct {
	head *LRUNode
	tail *LRUNode
}

func newDoubleLinkedList() *LRUDoubleLinkedList {
	return &LRUDoubleLinkedList{}
}

func (list *LRUDoubleLinkedList) addNode(node *LRUNode) {
	if node == nil {
		return
	}
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		node.last = list.tail
		list.tail = node
	}
}

func (list *LRUDoubleLinkedList) moveNodeToTail(node *LRUNode) {
	if list.tail == node {
		return
	}
	if list.head == node {
		list.head = list.head.next
		list.head.last = nil
	} else {
		node.last.next = node.next
		node.next.last = node.last
	}
	node.last = list.tail
	node.next = nil
	list.tail.next = node
	list.tail = node
}

func (list *LRUDoubleLinkedList) removeHead() *LRUNode {
	if list.head == nil {
		return nil
	}
	oldHead := list.head
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		list.head = oldHead.next
		oldHead.next = nil
		list.head.last = nil
	}
	return oldHead
}

type LRUCache struct {
	key2Node map[interface{}]*LRUNode
	node2Key map[*LRUNode]interface{}
	nodeList *LRUDoubleLinkedList
	capacity int
	size     int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		key2Node: make(map[interface{}]*LRUNode),
		node2Key: make(map[*LRUNode]interface{}),
		nodeList: newDoubleLinkedList(),
		capacity: capacity,
	}
}

func (cache *LRUCache) Get(key interface{}) interface{} {
	if node, ok := cache.key2Node[key]; ok {
		cache.nodeList.moveNodeToTail(node)
		return node.value
	}
	return nil
}

func (cache *LRUCache) Set(key interface{}, value interface{}) {
	if _, ok := cache.key2Node[key]; ok {
		node := cache.key2Node[key]
		node.value = value
		cache.nodeList.moveNodeToTail(node)
	} else {
		node := newLRUNode(value)
		cache.node2Key[node] = key
		cache.key2Node[key] = node
		cache.nodeList.addNode(node)
		cache.size += 1
		if cache.size > cache.capacity {
			cache.removeMostUnusedCache()
		}
	}
}

func (cache *LRUCache) removeMostUnusedCache() {
	node := cache.nodeList.removeHead()
	key := cache.node2Key[node]
	delete(cache.node2Key, node)
	delete(cache.key2Node, key)
}

package advanced

/*
	LFU Cache
	在构造时确定大小，为K

	有如下两个功能：
		1.set(key, value)：将记录插入该结构
		2.get(key)：返回key对应的value值

	要求：
		1.set和get的时间复杂度为O(1)
		2.某个key的set或get操作一旦发生，将这个key的操作次数加1
		3.当缓存的大小超过K时，移除操作次数最少的key，如果有操作次数相同的key，选择上次调用发生最早的key进行移除

	思路：操作次数相同的key放到一个双向链表中，将其看作一个桶，桶与桶之间按照操作次数从小到大连接起来也形成一个双向链表
		 当一个key发生了操作，就从自己的桶里出来，进入次数+1的桶中，如果没有该桶就创建一个。桶空了就删除
 		 缓存达到上限时删除第一个桶的尾节点

	难点：始终维持桶之间、节点之间的双向链表关系
*/

type LFUNode struct {
	key      interface{}
	value    interface{}
	times    int
	lastNode *LFUNode
	nextNode *LFUNode
}

func newLFUNode(key interface{}, value interface{}, times int) *LFUNode {
	return &LFUNode{
		key:   key,
		value: value,
		times: times,
	}
}

type LFUNodeBucket struct {
	head       *LFUNode
	tail       *LFUNode
	lastBucket *LFUNodeBucket
	nextBucket *LFUNodeBucket
}

func newLFUNodeBucket(node *LFUNode) *LFUNodeBucket {
	return &LFUNodeBucket{
		head: node,
		tail: node,
	}
}

func (bucket *LFUNodeBucket) addNodeToHead(node *LFUNode) {
	node.nextNode = bucket.head
	bucket.head.lastNode = node
	bucket.head = node
}

func (bucket *LFUNodeBucket) isEmpty() bool {
	return bucket.head == nil
}

func (bucket *LFUNodeBucket) deleteNode(node *LFUNode) {
	if bucket.head == bucket.tail {
		bucket.head = nil
		bucket.tail = nil
	} else {
		if node == bucket.head {
			bucket.head = node.nextNode
			bucket.head.lastNode = nil
		} else if node == bucket.tail {
			bucket.tail = node.lastNode
			bucket.tail.nextNode = nil
		} else {
			node.lastNode.nextNode = node.nextNode
			node.nextNode.lastNode = node.lastNode
		}
	}
	node.lastNode = nil
	node.nextNode = nil
}

type LFUCache struct {
	capacity    int
	size        int
	key2Node    map[interface{}]*LFUNode
	node2Bucket map[*LFUNode]*LFUNodeBucket
	headBucket  *LFUNodeBucket
}

func newLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity:    capacity,
		key2Node:    make(map[interface{}]*LFUNode),
		node2Bucket: make(map[*LFUNode]*LFUNodeBucket),
	}
}

func (cache *LFUCache) Get(key interface{}) interface{} {
	if node, ok := cache.key2Node[key]; ok {
		node.times += 1
		curBucket := cache.node2Bucket[node]
		cache.moveForward(node, curBucket)
		return node.value
	}
	return nil
}

func (cache *LFUCache) Set(key interface{}, value interface{}) {
	if node, ok := cache.key2Node[key]; ok {
		node.value = value
		node.times += 1
		curBucket := cache.node2Bucket[node]
		cache.moveForward(node, curBucket)
	} else {
		if cache.size == cache.capacity {
			node := cache.headBucket.tail
			cache.headBucket.deleteNode(node)
			cache.modifyCurrentBucketIfEmpty(cache.headBucket)
			delete(cache.key2Node, node.key)
			cache.size -= 1
		}
		node := newLFUNode(key, value, 1)
		if cache.headBucket == nil {
			cache.headBucket = newLFUNodeBucket(node)
		} else {
			if cache.headBucket.head.times == node.times {
				cache.headBucket.addNodeToHead(node)
			} else {
				newBucket := newLFUNodeBucket(node)
				newBucket.nextBucket = cache.headBucket
				cache.headBucket.lastBucket = newBucket
				cache.headBucket = newBucket
			}
		}
		cache.key2Node[key] = node
		cache.node2Bucket[node] = cache.headBucket
		cache.size += 1
	}
}

// node节点操作次数增加了1，将它移到操作次数加1的桶
func (cache *LFUCache) moveForward(node *LFUNode, curBucket *LFUNodeBucket) {
	curBucket.deleteNode(node)

	var preBucket *LFUNodeBucket
	if cache.modifyCurrentBucketIfEmpty(curBucket) {
		preBucket = curBucket.lastBucket
	} else {
		preBucket = curBucket
	}

	nextBucket := curBucket.nextBucket
	if nextBucket == nil {
		newBucket := newLFUNodeBucket(node)
		if preBucket != nil {
			preBucket.nextBucket = newBucket
		}
		newBucket.lastBucket = preBucket
		if cache.headBucket == nil {
			cache.headBucket = newBucket
		}
		cache.node2Bucket[node] = newBucket
	} else {
		if nextBucket.head.times == node.times {
			nextBucket.addNodeToHead(node)
			cache.node2Bucket[node] = nextBucket
		} else {
			newBucket := newLFUNodeBucket(node)
			if preBucket != nil {
				preBucket.nextBucket = newBucket
			}
			newBucket.lastBucket = preBucket
			newBucket.nextBucket = nextBucket
			nextBucket.lastBucket = newBucket
			if cache.headBucket == nextBucket {
				cache.headBucket = newBucket
			}
			cache.node2Bucket[node] = newBucket
		}
	}
}

// 判断bucket是否为空。bucket: 刚刚减少了一个节点的桶
// 1.如果不空，则什么也不做
// 2.如果空了，且bucket是最左的桶，则删掉这个桶的同时让其下一个成为最左的桶
// 3.如果空了，且bucket不是最左的桶，则删掉这个桶的同时让前后维持双向链表
func (cache *LFUCache) modifyCurrentBucketIfEmpty(bucket *LFUNodeBucket) bool {
	if bucket.isEmpty() {
		if cache.headBucket == bucket {
			cache.headBucket = bucket.nextBucket
			if cache.headBucket != nil {
				cache.headBucket.lastBucket = nil
			}
		} else {
			bucket.lastBucket.nextBucket = bucket.nextBucket
			if bucket.nextBucket != nil {
				bucket.nextBucket.lastBucket = bucket.lastBucket
			}
		}
		return true
	} else {
		return false
	}
}

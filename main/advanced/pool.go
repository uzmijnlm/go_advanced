package advanced

import (
	"math/rand"
	"time"
)

/*
	RandomPool

	实现如下三个功能：
		1.insert(key)：将某个key加入到该结构，做到不重复加入
		2.delete(key)：将原本在结构中的某个key移除
		3.getRandom()：等概率随即返回结构中的任何一个key

	要求：三个操作时间复杂度都是O(1)
*/

type Pool struct {
	key2Index map[interface{}]int
	index2Key map[int]interface{}
	size      int
}

func NewPool() *Pool {
	return &Pool{
		key2Index: make(map[interface{}]int),
		index2Key: make(map[int]interface{}),
	}
}

func (pool *Pool) insert(key interface{}) {
	if _, ok := pool.key2Index[key]; !ok {
		pool.key2Index[key] = pool.size
		pool.index2Key[pool.size] = key
		pool.size += 1
	}
}

func (pool *Pool) delete(key interface{}) {
	if index, ok := pool.key2Index[key]; ok {
		lastIndex := pool.size - 1
		lastKey := pool.index2Key[lastIndex]
		pool.key2Index[lastKey] = index
		pool.index2Key[index] = lastKey
		delete(pool.index2Key, lastIndex)
		delete(pool.key2Index, key)
	}
}

func (pool *Pool) getRandom() interface{} {
	if pool.size == 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(pool.size)
	return pool.index2Key[randIndex]
}

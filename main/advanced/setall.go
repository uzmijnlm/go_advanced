package advanced

import "time"

/*
	setAll hash

	有如下四个功能：
		1.put(key, value)：将记录插入该结构
		2.get(key)：返回key对应的value值
		3.containsKey(key)：判断是否包含某个key
		4.setAll(value)：将所有key的值统一设置成value

	要求：四个操作的时间复杂度为O(1)

	思路：加入一个时间戳结构，步骤如下：
		1.把每一个记录都加上一个时间，标记每条记录是何时建立的
		2.设置一个setAll记录，也加上时间戳，标记setAll记录的建立时间
		3.查询时对时间戳进行判断，以此为依据决定返回本身的记录还是setAll的记录
*/

type valueWithTime struct {
	value interface{}
	time  int64
}

func newValueWithTime(value interface{}, time int64) *valueWithTime {
	return &valueWithTime{
		value: value,
		time:  time,
	}
}

type SetAllHash struct {
	baseMap     map[interface{}]*valueWithTime
	setAllValue *valueWithTime
}

func NewSetAllHash() *SetAllHash {
	return &SetAllHash{
		baseMap:     make(map[interface{}]*valueWithTime),
		setAllValue: newValueWithTime(nil, time.Now().UnixNano()),
	}
}

func (hash *SetAllHash) containsKey(key interface{}) bool {
	if _, ok := hash.baseMap[key]; ok {
		return true
	} else {
		return false
	}
}

func (hash *SetAllHash) put(key interface{}, value interface{}) {
	hash.baseMap[key] = newValueWithTime(value, time.Now().UnixNano())
}

func (hash *SetAllHash) setAll(value interface{}) {
	hash.setAllValue = newValueWithTime(value, time.Now().UnixNano())
}

func (hash *SetAllHash) get(key interface{}) interface{} {
	if hash.containsKey(key) {
		if hash.baseMap[key].time > hash.setAllValue.time {
			return hash.baseMap[key].value
		} else {
			return hash.setAllValue.value
		}
	} else {
		return nil
	}
}

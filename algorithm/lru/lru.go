package lru

import (
	alg "GoLearn/algorithm"
	"container/list"
	"errors"
	"unsafe"
)

type LRU struct {
	maxbytes int
	cache    map[string]*list.Element
	size     int
	list     *list.List
}

func NewLRU(maxbytes int) *LRU {

	return &LRU{
		maxbytes: maxbytes,
		cache:    make(map[string]*list.Element),
		size:     0,
		list:     list.New(),
	}
}

func (lru *LRU) Get(key string) (interface{}, error) {
	if element, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(element)
		return element.Value.(alg.Entry), nil
	}
	return nil, errors.New("cannot find")
}

func (lru *LRU) Put(key string, value interface{}) error {

	// 如果已经存在，则更新值，顺序
	if element, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(element)
		element.Value = value
		lru.cache[key] = element
	} else {
		// 如果不存在，则添加到链表头部,更改内存大小
		entry := alg.Entry{Key: key, Val: value}
		s := int(unsafe.Sizeof(entry))
		// 单个数据超过内存
		if s > lru.maxbytes {
			return errors.New("single data out of maxbytes")
		}
		ele := lru.list.PushFront(entry)
		lru.cache[key] = ele
		lru.size += s
	}
	for lru.size > lru.maxbytes {
		lru.RemoveLastElement()
	}
	return nil
}

func (lru *LRU) RemoveLastElement() {
	ele := lru.list.Back()
	entry := ele.Value.(alg.Entry)
	lru.size -= int(unsafe.Sizeof(entry))
	lru.list.Remove(ele)
}

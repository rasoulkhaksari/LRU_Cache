package LRU

import "container/list"

type LRUCache2 struct {
	capacity int
	list     *list.List
	m        map[int]*list.Element
}

type KeyVal struct {
	key int
	val int
}

func New(capacity int) LRUCache2 {
	var cache LRUCache2
	cache.capacity = capacity
	cache.list = list.New()
	cache.m = make(map[int]*list.Element)
	return cache
}

func (lruc *LRUCache2) Get(key int) int {
	elem, ok := lruc.m[key]
	if !ok {
		return -1
	}
	lruc.list.MoveToFront(elem)
	return elem.Value.(*KeyVal).val
}

func (lruc *LRUCache2) Put(key int, value int) {
	_, ok := lruc.m[key]
	if ok {
		elem := lruc.m[key]
		lruc.list.MoveToFront(elem)
		elem.Value = &KeyVal{key, value}
	} else {
		if lruc.list.Len() == lruc.capacity {
			elem := lruc.list.Back()
			mkey := elem.Value.(*KeyVal).key
			delete(lruc.m, mkey)
			lruc.list.Remove(elem)
		}
		lruc.list.PushFront(&KeyVal{key, value})
		lruc.m[key] = lruc.list.Front()
	}

}

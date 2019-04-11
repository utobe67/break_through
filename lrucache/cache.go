package lrucache

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	mutex      sync.RWMutex
	size       int
	items      map[interface{}]*list.Element
	linkedList *list.List
}

type entry struct {
	key   interface{}
	value interface{}
}

func NewCache(size int) LRUCache {
	return LRUCache{
		size:       size,
		items:      make(map[interface{}]*list.Element),
		linkedList: list.New(),
	}
}

func (this *LRUCache) Get(key interface{}) (val interface{}, found bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	e, exists := this.items[key]
	if exists {
		this.linkedList.MoveToFront(e)
		return e.Value.(entry).value, true
	}
	return
}

func (this *LRUCache) Set(key, val interface{}) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	item := entry{
		key:   key,
		value: val,
	}

	e, exists := this.items[key]
	if exists {
		this.linkedList.MoveToFront(e)
		this.items[key].Value = item
	} else {
		this.items[key] = this.linkedList.PushFront(item)

		if len(this.items) > this.size {
			e := this.linkedList.Back()
			delete(this.items, e.Value.(entry).key)
			this.linkedList.Remove(e)
		}
	}
}

func (this *LRUCache) Peek(key interface{}) (val interface{}, found bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	e, exists := this.items[key]
	if exists {
		return e.Value.(entry).value, true
	}
	return
}

func (this *LRUCache) Remove(key interface{}) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	e, exists := this.items[key]
	if exists {
		this.linkedList.Remove(e)
		delete(this.items, e.Value.(entry).key)
	}
}

func (this *LRUCache) Keys() []interface{} {
	this.mutex.RLock()
	i := 0
	keys := make([]interface{}, len(this.items))
	for e := this.linkedList.Front(); e != nil; e = e.Next() {
		keys[i] = e.Value.(entry).key
		i++
	}
	this.mutex.RUnlock()
	return keys
}

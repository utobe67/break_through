package lrucache

/*
import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	valueMap map[int]int
	order    *list.List
	indexMap map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		valueMap: make(map[int]int),
		order:    list.New(),
		indexMap: make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.valueMap[key]
	if ok {
		this.order.MoveToBack(this.indexMap[key])
		return val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.valueMap[key]
	if ok {
		this.order.MoveToBack(this.indexMap[key])
		this.valueMap[key] = value
	} else {
		if len(this.valueMap) < this.capacity {
			this.valueMap[key] = value
			this.indexMap[key] = this.order.PushBack(key)
		} else {
			val := this.order.Front().Value.(int)
			this.order.Remove(this.order.Front())
			delete(this.valueMap, val)

			this.valueMap[key] = value
			this.indexMap[key] = this.order.PushBack(key)
		}
	}
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Printf("%d\n", cache.Get(1))
	cache.Put(3, 3)
	fmt.Printf("%d\n", cache.Get(2))
	cache.Put(4, 4)
	fmt.Printf("%d\n", cache.Get(1))
	fmt.Printf("%d\n", cache.Get(3))
	fmt.Printf("%d\n", cache.Get(4))
}
*/

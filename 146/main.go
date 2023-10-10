package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/lru-cache/
func main() {

	lRUCache := Constructor(2)

	lRUCache.Put(1, 1)           // cache is {1=1}
	lRUCache.Put(2, 2)           // cache is {1=1, 2=2}
	fmt.Println(lRUCache.Get(1)) // return 1
	lRUCache.Put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lRUCache.Get(2)) // returns -1 (not found)
	lRUCache.Put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lRUCache.Get(1)) // return -1 (not found)
	fmt.Println(lRUCache.Get(3)) // return 3
	fmt.Println(lRUCache.Get(4)) // return 4
}

type LRUCache struct {
	Capacity int
	Items    map[int]*list.Element
	Queue    *list.List
}

type Item struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Items:    make(map[int]*list.Element, capacity),
		Queue:    list.New(),
	}
}

func (this *LRUCache) Get(key int) int {

	element, exists := this.Items[key]
	if exists == false {
		return -1
	}
	this.Queue.MoveToFront(element)
	return element.Value.(*Item).Value
}

func (this *LRUCache) Put(key int, value int) {
	if element, exists := this.Items[key]; exists == true {
		this.Queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return
	}

	if this.Queue.Len() == this.Capacity {
		if element := this.Queue.Back(); element != nil {
			item := this.Queue.Remove(element).(*Item)
			delete(this.Items, item.Key)
		}
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	element := this.Queue.PushFront(item)
	this.Items[item.Key] = element

	return
}

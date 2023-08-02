package lrucache

import (
	"time"
)

type Node[T any, K comparable] struct {
	key     K
	value   T
	created time.Time
	timeout int // seconds
	prev,
	next *Node[T, K]
}

type LRUCache[K comparable, T any] struct {
	space      int
	cache      map[K]*Node[T, K]
	head, tail *Node[T, K]
}

func New[K comparable, T any](capacity, timeout int) LRUCache[K, T] {
	cache := map[K]*Node[T, K]{}
	var zero T
	var zeroKey K
	head, tail := &Node[T, K]{zeroKey, zero, time.Now(), timeout, nil, nil}, &Node[T, K]{zeroKey, zero, time.Now(), timeout, nil, nil}
	head.next = tail
	tail.prev = head
	return LRUCache[K, T]{capacity, cache, head, tail}
}

func (lru *LRUCache[K, T]) AddNode(node *Node[T, K]) {
	node.next = lru.head.next
	node.prev = lru.head
	node.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache[K, T]) RemoveNode(node *Node[T, K]) {
	prev, next := node.prev, node.next
	node.prev.next = next
	node.next.prev = prev
}

func (lru *LRUCache[K, T]) Get(key K) (T, bool) {
	if node, ok := lru.cache[key]; ok {
		lru.RemoveNode(node)
		lru.AddNode(node)
		return node.value, true
	} else {
		var res T
		return res, false
	}
}

func (lru *LRUCache[K, T]) Put(key K, value T) {
	if node, ok := lru.cache[key]; ok {
		node.value = value
		lru.RemoveNode(node)
		lru.AddNode(node)
	} else {
		if lru.space == 0 {
			toRemove := lru.tail.prev
			lru.RemoveNode(toRemove)
			delete(lru.cache, toRemove.key)
		} else {
			lru.space--
		}
		newNode := &Node[T, K]{key, value, time.Now(), 0, nil, nil}
		lru.AddNode(newNode)
		lru.cache[key] = newNode
	}
}

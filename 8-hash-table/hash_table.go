package hashtable

import (
	"hash/fnv"
	linkedlist "new-way/2-linked-list"
)

const DefaultCapacity = 3
const DefaultLoadFactor = 0.75

func defaultHash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))

	return h.Sum64()
}

type hashEntity struct {
	key   string
	value interface{}
}

type HashTable interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Unset(key string)
	Size() int
}

func New(capacity int) HashTable {
	if capacity < DefaultCapacity {
		capacity = DefaultCapacity
	}

	return &HashTableSeperateChaining{
		capacity:      capacity,
		hash:          defaultHash,
		loadFactor:    DefaultLoadFactor,
		resizeTrigger: int(float64(capacity) * DefaultLoadFactor),
		bucket:        make([]*linkedlist.DoublyLinkedList, capacity),
	}
}

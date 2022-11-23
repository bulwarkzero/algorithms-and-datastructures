package hashtable

import (
	"bytes"
	"fmt"
	linkedlist "new-way/2-linked-list"
	"strings"
)

type HashTableSeperateChaining struct {
	bucket        []*linkedlist.DoublyLinkedList
	capacity      int
	loadFactor    float64
	resizeTrigger int
	hash          func(s string) uint64
	size          int
}

func (h *HashTableSeperateChaining) keyIndex(key string) int {
	return int(h.hash(key) % uint64(h.capacity))
}

func (h *HashTableSeperateChaining) needsResize() bool {
	return h.size >= h.resizeTrigger
}

func (h *HashTableSeperateChaining) resize() {
	h.capacity *= 2
	h.resizeTrigger = int(float64(h.capacity) * h.loadFactor)

	newBucket := make([]*linkedlist.DoublyLinkedList, h.capacity)

	for i := 0; i < len(h.bucket); i++ {
		bucketValue := h.bucket[i]
		if bucketValue == nil {
			continue
		}

		bucketNode := bucketValue.Head()
		for bucketNode != nil {
			newIndex := h.keyIndex(bucketNode.Value().(hashEntity).key)
			if collisionalEntity := newBucket[newIndex]; collisionalEntity != nil {
				collisionalEntity.Add(bucketNode.Value())
			} else {
				newBucket[newIndex] = linkedlist.New()
				newBucket[newIndex].Add(bucketNode.Value())
			}
			bucketNode = bucketNode.Next()
		}
	}

	h.bucket = newBucket
}

func (h *HashTableSeperateChaining) getEntity(key string, index int) *linkedlist.Node {
	if index >= len(h.bucket) || h.bucket[index] == nil {
		return nil
	}

	bucketItem := h.bucket[index].Head()

	for bucketItem != nil {
		if bucketItem.Value().(hashEntity).key == key {
			return bucketItem
		}

		bucketItem = bucketItem.Next()
	}

	return nil
}

func (h *HashTableSeperateChaining) Set(key string, value interface{}) {
	index := h.keyIndex(key)

	// empty spot
	if h.bucket[index] == nil {
		ll := linkedlist.New()
		ll.Add(hashEntity{key: key, value: value})
		h.bucket[index] = ll
		h.size++

		return
	}

	// key exists, we need to update value
	if node := h.getEntity(key, index); node != nil {
		node.SetValue(hashEntity{key: key, value: value})

		return
	}

	// key collision
	h.bucket[index].Add(hashEntity{key: key, value: value})
	h.size++

	if h.needsResize() {
		h.resize()
	}
}

func (h *HashTableSeperateChaining) Get(key string) interface{} {
	index := h.keyIndex(key)

	entity := h.getEntity(key, index)
	if entity != nil {
		return entity.Value().(hashEntity).value
	}

	return nil
}

func (h *HashTableSeperateChaining) Unset(key string) {
	index := h.keyIndex(key)

	entity := h.getEntity(key, index)
	if entity != nil {
		h.bucket[index].RemoveNode(entity)
	}

	h.size--
}

func (h *HashTableSeperateChaining) Size() int {
	return h.size
}

// returns the stringified view of values distribution in hash backend
func (h *HashTableSeperateChaining) Distribution() string {
	buffer := bytes.NewBuffer([]byte(""))

	for i := 0; i < len(h.bucket); i++ {
		value := h.bucket[i]
		if value == nil {
			buffer.WriteString(fmt.Sprintf("%d => nil\r\n", i))
			continue
		}

		var valuesString string

		currentNode := value.Head()
		for currentNode != nil {
			nodeValue := currentNode.Value().(hashEntity)
			valuesString += fmt.Sprintf("{%v => %v}, ", nodeValue.key, nodeValue.value)

			currentNode = currentNode.Next()
		}

		valuesString = strings.TrimRight(valuesString, ", ")

		buffer.WriteString(fmt.Sprintf("%d => [%s]\r\n", i, valuesString))
	}

	return buffer.String()
}

func (h *HashTableSeperateChaining) String() string {
	buffer := bytes.NewBuffer([]byte(""))

	for i := 0; i < len(h.bucket); i++ {
		value := h.bucket[i]
		if value == nil {
			continue
		}

		currentNode := value.Head()
		for currentNode != nil {
			nodeValue := currentNode.Value().(hashEntity)
			buffer.WriteString(fmt.Sprintf("{%v => %v}, ", nodeValue.key, nodeValue.value))

			currentNode = currentNode.Next()
		}
	}

	return strings.TrimRight(buffer.String(), ", ")
}

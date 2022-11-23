package linkedlist

import (
	"fmt"
	"strings"
)

type Node struct {
	value interface{}
	prev,
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) SetValue(value interface{}) {
	n.value = value
}

type DoublyLinkedList struct {
	size int
	head,
	tail *Node
}

func (l *DoublyLinkedList) Head() *Node {
	return l.head
}

func (l *DoublyLinkedList) Tail() *Node {
	return l.tail
}

func (l *DoublyLinkedList) Size() int {
	return l.size
}

func (l *DoublyLinkedList) IsEmpty() bool {
	return l.size < 1
}

func (l *DoublyLinkedList) Add(val interface{}) *DoublyLinkedList {
	newNode := &Node{
		value: val,
		prev:  l.tail,
		next:  nil,
	}

	if l.IsEmpty() {
		l.head = newNode
		l.tail = newNode
		l.size++

		return l
	}

	l.tail.next = newNode
	l.tail = l.tail.next
	l.size++

	return l
}

func (l *DoublyLinkedList) RemoveFirst() interface{} {
	if l.IsEmpty() {
		panic("remove from empty list is not possible")
	}

	value := l.head.value
	l.head = l.head.next
	l.size--

	if l.IsEmpty() {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	return value
}

func (l *DoublyLinkedList) RemoveLast() interface{} {
	if l.IsEmpty() {
		panic("remove from empty list is not possible")
	}

	value := l.tail.value
	l.tail = l.tail.prev
	l.size--

	if l.IsEmpty() {
		l.head = nil
	} else {
		l.tail.next = nil
	}

	return value
}

func (l *DoublyLinkedList) RemoveNode(node *Node) (value interface{}) {
	// it's head
	if node.prev == nil {
		return l.RemoveFirst()
	}

	// it's tail
	if node.next == nil {
		return l.RemoveLast()
	}

	value = node.value

	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--

	return value
}

func (l *DoublyLinkedList) RemoveAt(index int) interface{} {
	if index < 0 || index > l.Size()-1 {
		panic("invalid index")
	}

	node := l.head
	var i int

	for node != nil && i < index {
		node = node.next
		i++
	}

	return l.RemoveNode(node)
}

func (l *DoublyLinkedList) String() string {
	el := l.head
	str := ""

	for el != nil {
		str += fmt.Sprintf("(%d) -> ", el.value)
		el = el.next
	}

	return strings.TrimRight(str, " -> ")
}

func New() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

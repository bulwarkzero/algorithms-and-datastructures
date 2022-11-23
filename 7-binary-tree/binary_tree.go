package binarytree

import queue "new-way/4-queue"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
	len  int
}

func (tree *BinaryTree) traversePreOrder(root *Node, f func(*Node)) {
	if root == nil {
		return
	}

	if f != nil {
		f(root)
	}

	tree.traversePreOrder(root.left, f)
	tree.traversePreOrder(root.right, f)
}

func (tree *BinaryTree) traversePostOrder(root *Node, f func(*Node)) {
	if root == nil {
		return
	}

	tree.traversePreOrder(root.left, f)
	tree.traversePreOrder(root.right, f)

	if f != nil {
		f(root)
	}
}

func (tree *BinaryTree) traverseInOrder(root *Node, f func(*Node)) {
	if root == nil {
		return
	}

	tree.traversePreOrder(root.left, f)
	if f != nil {
		f(root)
	}
	tree.traversePreOrder(root.right, f)
}

func (tree *BinaryTree) Insert(value int) {
	tree.len++

	if tree.root == nil {
		tree.root = &Node{value: value}

		return
	}

	currentNode := tree.root
	for {
		// we didn't want duplicate values in tree
		if value == currentNode.value {
			return
		}

		// right nodes are holding greater values
		if value > currentNode.value {
			if currentNode.right == nil {
				currentNode.right = &Node{value: value}
				return
			}

			currentNode = currentNode.right
			continue
		}

		// left nodes are holding less values
		if value < currentNode.value {
			if currentNode.left == nil {
				currentNode.left = &Node{value: value}
				return
			}

			currentNode = currentNode.left
			continue
		}
	}
}

func (tree *BinaryTree) Find(value int) (node *Node) {
	currentNode := tree.root

	for currentNode != nil && value != currentNode.value {
		if value > currentNode.value {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
		}
	}

	return currentNode
}

func (tree *BinaryTree) minNode(root *Node) *Node {
	minNode := root
	for root.left != nil {
		minNode = root.left
		root = root.left
	}

	return minNode
}

func (tree *BinaryTree) minValue(root *Node) int {
	return tree.minNode(root).value
}

func (tree *BinaryTree) remove(root *Node, value int) *Node {
	if root == nil {
		return root
	}

	if value > root.value {
		root.right = tree.remove(root.right, value)
	} else if value < root.value {
		root.left = tree.remove(root.left, value)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		// node has two children
		root.value = tree.minValue(root.right)
		root.right = tree.remove(root.right, root.value)
	}

	return root
}

func (tree *BinaryTree) Remove(value int) {
	tree.len--

	tree.root = tree.remove(tree.root, value)
}

func (tree *BinaryTree) Len() int {
	return tree.len
}

func (tree *BinaryTree) TraversePreOrder(f func(*Node)) {
	tree.traversePreOrder(tree.root, f)
}

func (tree *BinaryTree) TraversePostOrder(f func(*Node)) {
	tree.traversePostOrder(tree.root, f)
}

// traverse items in order from smaller to larger
func (tree *BinaryTree) TraverseInOrder(f func(*Node)) {
	tree.traverseInOrder(tree.root, f)
}

// traverse items in each tree level
func (tree *BinaryTree) TraverseLevelOrder(f func(*Node)) {
	queue := queue.NewArrayBasedQueue()

	root := tree.root
	queue.Enqueue(root)

	for queue.Len() > 0 {
		root = queue.Dequeue().(*Node)

		f(root)

		if root.left != nil {
			queue.Enqueue(root.left)
		}

		if root.right != nil {
			queue.Enqueue(root.right)
		}
	}
}

func New() *BinaryTree {
	return &BinaryTree{}
}

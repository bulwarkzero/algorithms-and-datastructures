package avltree

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

type Node struct {
	value  int
	left   *Node
	right  *Node
	height int // node height (max of left and right node height)
	bf     int // balance factor
}

type AVLTree struct {
	root *Node
	len  int
}

func (tree *AVLTree) minNode(root *Node) *Node {
	minNode := root
	for root.left != nil {
		minNode = root.left
		root = root.left
	}

	return minNode
}

func (tree *AVLTree) minValue(root *Node) int {
	return tree.minNode(root).value
}

func (tree *AVLTree) insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{value: value}
	}

	if value > node.value {
		node.right = tree.insert(node.right, value)
	} else {
		node.left = tree.insert(node.left, value)
	}

	tree.update(node)

	return tree.balance(node)
}

func (tree *AVLTree) update(node *Node) {
	leftNodeHeight := -1
	rightNodeHeight := -1

	if node.left != nil {
		leftNodeHeight = node.left.height
	}

	if node.right != nil {
		rightNodeHeight = node.right.height
	}

	node.height = 1 + max(leftNodeHeight, rightNodeHeight)

	node.bf = rightNodeHeight - leftNodeHeight
}

func (tree *AVLTree) leftRotate(node *Node) *Node {
	b := node.right
	node.right = b.left
	b.left = node

	tree.update(node)
	tree.update(b)

	return b
}

func (tree *AVLTree) rightRotate(node *Node) *Node {
	b := node.left
	node.left = b.right
	b.right = node

	tree.update(node)
	tree.update(b)

	return b
}

func (tree *AVLTree) leftLeft(node *Node) *Node {
	return tree.rightRotate(node)
}

func (tree *AVLTree) leftRight(node *Node) *Node {
	node.left = tree.leftRotate(node.left)

	return tree.leftLeft(node)
}

func (tree *AVLTree) rightRight(node *Node) *Node {
	return tree.leftRotate(node)
}

func (tree *AVLTree) rightLeft(node *Node) *Node {
	node.right = tree.rightRotate(node.right)

	return tree.rightRight(node)
}

// Re-balance a tree if it's balance factor is +2 or -2
func (tree *AVLTree) balance(node *Node) *Node {
	// Left heavy
	if node.bf == -2 {
		// Left Left
		if node.left.bf <= 0 {
			return tree.leftLeft(node)
			// Left Right
		} else {
			return tree.leftRight(node)
		}
		// Right heavy
	} else if node.bf == 2 {
		// Right Right
		if node.right.bf >= 0 {
			return tree.rightRight(node)
			// Right Left
		} else {
			return tree.rightLeft(node)
		}
	}

	return node
}

func (tree *AVLTree) remove(root *Node, value int) *Node {
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

	tree.update(root)

	return tree.balance(root)
}

func (tree *AVLTree) Len() int {
	return tree.len
}

func (tree *AVLTree) IsEmpty() bool {
	return tree.len < 1
}

func (tree *AVLTree) Contains(value int) bool {
	if tree.IsEmpty() {
		return false
	}

	currentNode := tree.root
	for currentNode != nil {
		if value == currentNode.value {
			return true
		}

		if value > currentNode.value {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
		}
	}

	return false
}

func (tree *AVLTree) Find(value int) (node *Node) {
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

func (tree *AVLTree) Insert(value int) bool {
	// we don't want duplicate values
	if tree.Contains(value) {
		return false
	}

	tree.root = tree.insert(tree.root, value)
	tree.len++

	return true
}

func (tree *AVLTree) Remove(value int) {
	tree.len--

	tree.root = tree.remove(tree.root, value)
}

func New() *AVLTree {
	return &AVLTree{}
}

package tree

import "fmt"

//Node tree node
type Node struct {
	value int
	left  *Node
	right *Node
}

//BinarySearchTree a tree
type BinarySearchTree struct {
	root *Node
}

//Insert the new value in tree
func (bst *BinarySearchTree) Insert(value int) {
	newNode := &Node{value, nil, nil}
	if bst.root == nil {
		bst.root = newNode
	} else {
		insertNode(bst.root, newNode)
	}
}

func insertNode(root, newNode *Node) {
	if newNode.value < root.value {
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}
	} else {
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}
}

//Search node from tree by value
func (bst *BinarySearchTree) Search(value int) *Node {
	return search(bst.root, value)
}

func search(node *Node, value int) *Node {
	fmt.Println(node.value, &node)
	if node == nil {
		return nil
	}

	if node.value == value {
		return node
	}

	if value < node.value {
		return search(node.left, value)
	} else if value > node.value {
		return search(node.right, value)
	}

	return nil
}

//Remove a value from tree
func (bst *BinarySearchTree) Remove(value int) bool {
	if bst.root == nil {
		return false
	}
	fmt.Println(bst.root.value, &bst.root)
	node := search(bst.root, value)
	if node == nil {
		return false
	}
	fmt.Println(&node, &bst.root.left.left)
	return removeNode(node)
}

func removeNode(node *Node) bool {
	//case 1 node has no child
	if node.left == nil && node.right == nil {
		node = nil
		return true
	}
	//case 2 node has one child
	if node.left == nil {
		node = node.right
		return true
	} else if node.right == nil {
		node = node.left
		return true
	}
	//case 3 node has two child
	//find minimum value from right subtree
	rightMinNode := node.right
	for {
		if rightMinNode != nil && rightMinNode.left != nil {
			rightMinNode = rightMinNode.left
		} else {
			break
		}
	}
	node.value = rightMinNode.value
	rightMinNode = nil
	return true
}

//ToString show a tree by string
func (bst *BinarySearchTree) ToString() {
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.value)
		stringify(n.right, level)
	}
}

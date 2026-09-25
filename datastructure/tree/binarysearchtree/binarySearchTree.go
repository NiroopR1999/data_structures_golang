package binarysearchtree

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func insertNode(root, newNode *Node) *Node {
	// If we reach an empty position, this is where
	// the new node should be inserted.
	if root == nil {
		root = newNode
	} else if newNode.value < root.value {
		// Smaller values belong in the left subtree
		// according to the BST property.
		root.left = insertNode(root.left, newNode)
	} else {
		// Greater or equal values go to the right subtree.
		root.right = insertNode(root.right, newNode)
	}

	return root
}

func (bst *BinarySearchTree) Insert(value int) {
	newNode := &Node{value: value}

	// We assign the returned node to root because if the tree
	// is empty, the new node becomes the root.
	bst.root = insertNode(bst.root, newNode)
}

func (bst *BinarySearchTree) IsEmpty() bool {
	// An empty BST has no root node.
	return bst.root == nil
}

func (bst *BinarySearchTree) Min(root *Node) int {
	if root == nil {
		return 0
	}

	// In a BST, the smallest value is the leftmost node.
	if root.left == nil {
		return root.value
	}

	return bst.Min(root.left)
}

func (bst *BinarySearchTree) Max(root *Node) int {
	if root == nil {
		return 0
	}

	// In a BST, the largest value is the rightmost node.
	if root.right == nil {
		return root.value
	}

	return bst.Max(root.right)
}

func (bst *BinarySearchTree) Search(root *Node, value int) bool {
	if root == nil {
		// We reached an empty position, so the value doesn't exist.
		return false
	}

	if root.value == value {
		return true
	} else if value < root.value {
		// Because this is a BST, a smaller value can only
		// exist in the left subtree.
		return bst.Search(root.left, value)
	} else {
		// A larger value can only exist in the right subtree.
		return bst.Search(root.right, value)
	}
}

func (bst *BinarySearchTree) PreOrder(root *Node) {
	if root == nil {
		return
	}

	// Root → Left → Right
	fmt.Print(root.value, " ")

	bst.PreOrder(root.left)
	bst.PreOrder(root.right)
}

func (bst *BinarySearchTree) InOrder(root *Node) {
	if root == nil {
		return
	}

	// Left → Root → Right
	//
	// This produces sorted order for a BST.
	bst.InOrder(root.left)

	fmt.Print(root.value, " ")

	bst.InOrder(root.right)
}

func (bst *BinarySearchTree) PostOrder(root *Node) {
	if root == nil {
		return
	}

	// Left → Right → Root
	bst.PostOrder(root.left)
	bst.PostOrder(root.right)

	fmt.Print(root.value, " ")
}

func (bst *BinarySearchTree) LevelOrder(root *Node) {
	if root == nil {
		return
	}

	// Level-order traversal requires BFS,
	// so we use a queue to process nodes level by level.
	queue := []*Node{root}

	for len(queue) > 0 {
		// Take the first node from the queue.
		res := queue[0]
		queue = queue[1:]

		fmt.Print(res.value, " ")

		// Add children to the queue so they are processed
		// after all nodes currently ahead of them.
		if res.left != nil {
			queue = append(queue, res.left)
		}

		if res.right != nil {
			queue = append(queue, res.right)
		}
	}
}

func deleteNode(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	if value < root.value {
		// Search for the value in the left subtree.
		root.left = deleteNode(root.left, value)
	} else if value > root.value {
		// Search for the value in the right subtree.
		root.right = deleteNode(root.right, value)
	} else {
		// We found the node that needs to be deleted.

		// Case 1:
		// No right child.
		//
		// The left child can replace the current node.
		if root.right == nil {
			return root.left
		}

		// Case 2:
		// No left child.
		//
		// The right child can replace the current node.
		if root.left == nil {
			return root.right
		}

		// Case 3:
		// The node has two children.
		//
		// Find the smallest node in the right subtree.
		// This is called the in-order successor.
		successor := root.right

		for successor.left != nil {
			successor = successor.left
		}

		// Copy the successor's value into the node
		// we actually wanted to delete.
		root.value = successor.value

		// Now remove the duplicate successor node
		// from the right subtree.
		root.right = deleteNode(root.right, successor.value)
	}

	return root
}

func (bst *BinarySearchTree) Delete(value int) {
	// The returned node matters because deleting the root
	// can change the tree's root.
	bst.root = deleteNode(bst.root, value)
}
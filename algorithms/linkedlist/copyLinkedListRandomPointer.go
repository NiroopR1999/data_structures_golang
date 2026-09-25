package linkedlist

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}	


	// Map each original node to its corresponding copied node.
	// We need this so we can find the copied version of any
	// node referenced by a Random pointer.
	nodeMap := make(map[*Node]*Node)

	// First pass: create a copy of every node.
	// We don't connect Next or Random yet because we first
	// need copies of all nodes to exist.
	current := head
	for current != nil {
		nodeMap[current] = &Node{
			Val: current.Val,
		}
		current = current.Next
	}

	// Second pass: connect Next and Random pointers.
	// nodeMap[original] gives us the corresponding copied node.
	current = head
	for current != nil {
		copyNode := nodeMap[current]

		// Connect Next to the copied version of the original Next.
		copyNode.Next = nodeMap[current.Next]

		// Connect Random to the copied version of the original Random.
		// If Random is nil, map lookup also gives nil.
		copyNode.Random = nodeMap[current.Random]

		current = current.Next
	}

	return nodeMap[head]
}

package btree

func binarySearch(input []int, target int) int {
	low, high := 0, len(input)
	for low < high {
		mid := (high + low) / 2
		if input[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func insertNewKey(input *[]int, target int) {
	s := *input
	if len(s) > 0 {
		insertIndex := binarySearch(s, target)
		if insertIndex != len(s) {
			for i := insertIndex; i < len(s); i++ {
				var temp int = s[i]
				s[i] = target
				target = temp
			}
		}
	}
	s = append(s, target)
	*input = s
}

func insertNewChildren(input *[]*bTreeNode, index int, target *bTreeNode) {
	s := *input
	if len(s) != index {
		for i := index; i < len(*input); i++ {
			var temp *bTreeNode = s[i]
			s[i] = target
			target = temp
		}
	}
	s = append(s, target)
	*input = s
}

func GetCorrectKeyToInsert(children []*bTreeNode, data int) int {
	for _, child := range children {
		if child.Keys[0] > data {
			return child.Keys[0]
		}
	}
	return data
}

func handleNodeSplit(node *bTreeNode, data int) (*bTreeNode, *bTreeNode, int) {
	newNode := newBTreeNode()
	splitPoint := len(node.Keys) / 2
	medianData := node.Keys[splitPoint]
	newNode.Keys = make([]int, len(node.Keys[splitPoint:]))
	copy(newNode.Keys, node.Keys[splitPoint:])
	node.Keys = node.Keys[:splitPoint]

	newNode.PrevSibling = node
	if node.NextSibling != nil{
		node.NextSibling.PrevSibling = newNode
	}
	newNode.NextSibling = node.NextSibling
	node.NextSibling = newNode

	if data < medianData {
		insertNewKey(&node.Keys, data)
	} else {
		insertNewKey(&newNode.Keys, data)
	}
	return node, newNode, medianData
}

func handleRootNodeSplit(node *bTreeNode, data int) *bTreeNode {
	node, newNode, medianData := handleNodeSplit(node, data)
	rootNode := newBTreeNode()
	rootNode.Keys = append(rootNode.Keys, medianData)
	rootNode.Children = append(rootNode.Children, node)
	rootNode.Children = append(rootNode.Children, newNode)

	//Update parent and sibling nodes of the newly created nodes.
	newNode.Parent = rootNode
	node.Parent = rootNode

	return rootNode
}

func GetAllChildren(node *bTreeNode, data int, index int) []*bTreeNode {
	nodeToSplit := node.Children[index]
	children := node.Children[:]
	newNode := newBTreeNode()
	newNode.Parent = node
	splitPoint := len(nodeToSplit.Keys) / 2
	medianData := nodeToSplit.Keys[splitPoint]
	newNode.Keys = make([]int, len(nodeToSplit.Keys[splitPoint:]))
	copy(newNode.Keys, nodeToSplit.Keys[splitPoint:])
	nodeToSplit.Keys = nodeToSplit.Keys[:splitPoint]

	nodeToSplit.NextSibling = newNode
	newNode.PrevSibling = node

	if data < medianData {
		insertNewKey(&nodeToSplit.Keys, data)
	} else {
		insertNewKey(&newNode.Keys, data)
	}
	insertNewChildren(&children, index+1, newNode)
	return children
}

// If neighboring nodes have two few values, the sibling nodes are merged.
// This situation is called underflow.
// This function checks for underflow and returns with which sibling
// it underflows i.e; with previous or next sibling.
func checkUnderFlow(node *bTreeNode, fanout int) string {
	if node.PrevSibling != nil {
		if len(node.Keys)+len(node.PrevSibling.Keys) <= fanout {
			return "prev"
		}
	}

	//If this is the last child node of a node, next sibling means
	//we are searching on the another node hierarchy.
	if node.NextSibling != nil && indexOfChildPointer(node, node.Parent) != len(node.Children) {
		if len(node.Keys)+len(node.NextSibling.Keys) <= fanout {
			return "next"
		}
	}

	return ""
}

// searchIndex function returns the index of the element
// in the array, returns -1, if the target element is not there in the array.
func searchIndex(input []int, target int) int {
	low, high := 0, len(input)-1
	for low <= high {
		mid := (low + high) / 2
		if input[mid] == target {
			return mid
		} else if input[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// This function returns the index of a child given one of the child
// pointers of a node.
func indexOfChildPointer(child *bTreeNode, node *bTreeNode) int {
	for i, children := range node.Children {
		if child == children {
			return i
		}
	}
	return -1
}

package trees

type avlNode struct {
	Data   int
	Height int
	Left   *avlNode
	Right  *avlNode
	Parent *avlNode
}

func newAVLNode(data int) *avlNode {
	node := &avlNode{
		Data:   data,
		Height: 1,
	}
	return node
}

func (node *avlNode) balanceFactor() int {
	return height(node.Left) - height(node.Right)
}

// Deletion helper function.
func (node *avlNode) predecessor() *avlNode {
	if node.Left == nil {
		return nil
	}
	node = node.Left
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (node *avlNode) leafNode() bool {
	return (node.Left == nil && node.Right == nil)
}

func (node *avlNode) isLeftChild(data int) bool{
	if node.Left != nil && node.Left.Data == data{
			return true
	}
	return false
}

func (node *avlNode) isRightChild(data int) bool{
	if node.Right != nil && node.Right.Data == data{
			return true
	}
	return false
}

type AVLTree struct {
	Root *avlNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (tree *AVLTree) Insert(data int) {
	newNode := newAVLNode(data)
	//If the tree is empty, create the root node
	if tree.Root == nil {
		tree.Root = newNode
	} else {
		node := tree.Root
		var prevNode *avlNode
		for node != nil {
			prevNode = node
			if node.Data > data {
				node = node.Left
			} else {
				node = node.Right
			}
		}
		if prevNode.Data > newNode.Data {
			prevNode.Left = newNode
			newNode.Parent = prevNode
		} else {
			prevNode.Right = newNode
			newNode.Parent = prevNode
		}
		//once a node is inserted, balance the tree.
		tree.updateHeight(newNode)
	}
}

func (tree *AVLTree) Delete(data int) {
	if tree.Root == nil {
		panic("You cannot delete an element from the empty tree.")
	}
	node := tree.Root
	var parent *avlNode
	for node != nil {
		if node.Data == data {
			if node.leafNode() {
				//If the node to delete is a leaf node.
				parent = node.Parent
				if parent.isLeftChild(data){
					parent.Left = nil
				}else if parent.isRightChild(data){
					parent.Right = nil
				}
				node = nil
			} else if node.Left != nil && node.Right == nil {
				//If the node has no right child, set parent of the deleted node to left
				//child of deleted node
				parent = node.Parent
				node.Left.Parent = node.Parent
				parent.Right = node.Left
				node = nil
			} else if node.Left == nil && node.Right != nil {
				//If the node has no left child, set parent of the deleted node to
				//right child of deleted node.
				parent = node.Parent
				node.Right.Parent = node.Parent
				parent.Right = node.Right
				node = nil
			} else {
				//If the node has both childs
				//Find the largest node in left subtree
				predecessor := node.predecessor()
				//copy the value of largest subtree into node to delete.
				node.Data = predecessor.Data
				parent = predecessor.Parent
				//check if the predecessor has any left child(There should be no right child)
				if predecessor.Left != nil {
					predecessor.Left.Parent = parent
					parent.Right = predecessor.Left
				}
				node = nil
			}
			//Correcting the parent height, after deleting the node.
			parent.Height = 1 + max(height(parent.Left), height(parent.Right))
			tree.updateHeight(parent)
		} else if node.Data > data {
			node = node.Left
		} else {
			node = node.Right
		}
	}
}

func height(node *avlNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func (tree *AVLTree) updateHeight(node *avlNode) {
	node = node.Parent
	for node != nil {
		node.Height = 1 + max(height(node.Left), height(node.Right))
		isBalanced := tree.checkForBalance(node)
		if !isBalanced {
			switch bf := node.balanceFactor(); bf {
			case 2:
				if node.Left.balanceFactor() >= 0 {
					node = tree.singleRightRotate(node)
				} else {
					node = tree.doubleRotateRight(node)
				}
			case -2:
				if node.Right.balanceFactor() <= 0 {
					node = tree.singleLeftRotate(node)
				} else {
					node = tree.doubleRotateLeft(node)
				}
			}
		}
		node = node.Parent
	}
}

func (tree *AVLTree) checkForBalance(node *avlNode) bool {
	return (node.balanceFactor() >= -1 && node.balanceFactor() <= 1)
}

func (tree *AVLTree) singleLeftRotate(node *avlNode) *avlNode {
	isRootNode := (tree.Root == node)
	temp := node
	node = node.Right
	temp.Right = node.Left
	node.Left = temp
	//Update Parents
	if temp.Parent != nil {
		temp.Parent.Right = node
	}
	node.Parent = temp.Parent
	temp.Parent = node
	//Update heights
	temp.Height = 1 + max(height(temp.Left), height(temp.Right))
	node.Height = 1 + max(height(node.Left), height(node.Right))
	if isRootNode {
		tree.Root = node
	}
	return node
}

func (tree *AVLTree) singleRightRotate(node *avlNode) *avlNode {
	isRootNode := (tree.Root == node)
	temp := node
	node = node.Left
	temp.Left = node.Right
	node.Right = temp
	//Update parents
	if temp.Parent != nil {
		temp.Parent.Left = node
	}
	node.Parent = temp.Parent
	temp.Parent = node
	//Update heights
	temp.Height = 1 + max(height(temp.Left), height(temp.Right))
	node.Height = 1 + max(height(node.Left), height(node.Right))
	if isRootNode {
		tree.Root = node
	}
	return node
}

func (tree *AVLTree) doubleRotateLeft(node *avlNode) *avlNode {
	isRootNode := (tree.Root == node)
	temp := node
	right := temp.Right
	node = right.Left
	temp.Right = node.Left
	right.Left = node.Right
	node.Left = temp
	node.Right = right
	//Updating parents
	if temp.Parent != nil {
		temp.Parent.Right = node
	}
	node.Parent = temp.Parent
	temp.Parent = node
	right.Parent = node
	if temp.Right != nil {
		temp.Right.Parent = temp
	}
	if right.Left != nil {
		right.Left.Parent = right
	}
	//Updating the height
	temp.Height = 1 + max(height(temp.Left), height(temp.Right))
	right.Height = 1 + max(height(right.Left), height(right.Right))
	node.Height = 1 + max(height(node.Left), height(node.Right))
	if isRootNode {
		tree.Root = node
	}
	return node
}

func (tree *AVLTree) doubleRotateRight(node *avlNode) *avlNode {
	isRootNode := (tree.Root == node)
	temp := node
	left := temp.Left
	node = left.Right
	temp.Left = node.Right
	left.Right = node.Left
	node.Right = temp
	node.Left = left
	//Updating parents
	if temp.Parent != nil {
		temp.Parent.Left = node
	}
	node.Parent = temp.Parent
	temp.Parent = node
	left.Parent = node
	if temp.Right != nil {
		temp.Right.Parent = temp
	}
	if left.Right != nil {
		left.Right.Parent = left
	}
	//Updating the height
	temp.Height = 1 + max(height(temp.Left), height(temp.Right))
	left.Height = 1 + max(height(left.Left), height(left.Right))
	node.Height = 1 + max(height(node.Left), height(node.Right))
	if isRootNode {
		tree.Root = node
	}
	return node
}

func (tree AVLTree) Height() int {
	return tree.Root.Height
}

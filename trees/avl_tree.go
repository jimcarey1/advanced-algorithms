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
	if temp.Parent != nil {
		temp.Parent.Right = node
	}
	node.Parent = temp.Parent
	temp.Parent = node
	temp.Height = temp.Height - 2
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
	node.Parent = temp.Parent
	temp.Parent = node
	temp.Height = temp.Height - 2
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
	node.Left = temp
	node.Right = right
	node.Parent = temp.Parent
	temp.Parent = node
	right.Parent = node
	temp.Height = temp.Height - 2
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
	node.Right = temp
	node.Left = left
	node.Parent = temp.Parent
	temp.Parent = node
	left.Parent = node
	if isRootNode {
		tree.Root = node
	}
	return node
}

func (tree AVLTree) Height() int {
	return tree.Root.Height
}

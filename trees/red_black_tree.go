package trees

type redBlackNode struct{
	Data int
	Left *redBlackNode
	Right *redBlackNode
	Parent *redBlackNode
	BlackHeight int
	Color string
}

func newRedBlackNode(data int) *redBlackNode{
	return &redBlackNode{
		Data: data,
		BlackHeight: 0,
		Color: "red",
	}
}

func (node *redBlackNode) isLeftChild() bool{
	return (node != nil) && (node.Parent.Left == node)
}

func (node *redBlackNode) isRightChild() bool{
	return (node != nil) && (node.Parent.Right == node)
}

//This function checks whether the node is red.
func (node *redBlackNode) IsRedNode() bool{
	return (node != nil) && (node.Color == "red")
}

//This function checks whether the node is black.
//null nodes are black.
func (node *redBlackNode) IsBlackNode() bool{
	return (node == nil) || (node.Color == "black")
}

type redBlackTree struct{
	Root *redBlackNode
}

func NewRedBlackTree() *redBlackTree{
	return &redBlackTree{}
}

//Insert operation.
func (tree *redBlackTree) Insert(data int) {
	newNode := newRedBlackNode(data)
	//If the root node is nil, make the newNode the root node.
	if tree.Root == nil{
		newNode.Color = "black"
		tree.Root = newNode
		return
	}

	//Find the position to insert the newNode.
	node := tree.Root
	var prevNode *redBlackNode
	for node != nil{
		prevNode = node
		if node.Data > newNode.Data{
			//Check in its left subtree.
			node = node.Left
		}else{
			node = node.Right
		}
	}
	//Insert the newNode and populate its parent data.
	newNode.Parent = prevNode
	if prevNode.Data > newNode.Data{
		prevNode.Left = newNode
	}else{
		prevNode.Right = newNode
	}

	//check for the balancing conditions.
	node = newNode
	var uncle *redBlackNode
	for node.Parent != nil{
		parent := node.Parent
		//case1: If the new node's parent is black, there is no violation.
		if parent.Color == "black"{
			return
		}

		grandparent := parent.Parent
		//If there is no grandparent, make the parent node black.
		if grandparent == nil{
			parent.Color = "black"
			return
		}

		//Case 5:
		//node is the right child of the left child of the grandparent
		if node.isRightChild() && parent.isLeftChild(){
			uncle = grandparent.Right
			if uncle == nil || uncle.Color == "black"{
				grandparent.Left = node
				node.Parent = grandparent
				parent.Right = node.Left
				if node.Left != nil{
					node.Left.Parent = parent
				}
				parent.Parent = node
				parent = node
			}
		}

		//node is the left child of the right child of the grandparent
		if node.isLeftChild() && parent.isRightChild(){
			uncle = grandparent.Left
			if uncle == nil || uncle.Color == "black"{
				grandparent.Right = node
				node.Parent = grandparent
				parent.Left = node.Right
				if node.Right != nil{
					node.Right.Parent = parent
				}
				parent.Parent = node
				parent = node
			}
		}

		//case 6:
		if parent.isLeftChild(){
			node = tree.singleRightRotate(grandparent)
			node.Color = "red"
			parent.Color = "black"
			return
		}else if parent.isRightChild(){
			node = tree.singleLeftRotate(grandparent)
			node.Color = "red"
			parent.Color = "black"
			return
		}

		//case2:
		//uncle is red, parent is red and node is red.
		//recoloring will work.
		parent.Color = "black"
		uncle.Color = "black"
		grandparent.Color = "red"
		node = grandparent
	}
}

func (tree *redBlackTree) singleLeftRotate(node *redBlackNode) *redBlackNode {
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
	if isRootNode {
		tree.Root = node
	}
	return node
}

func (tree *redBlackTree) singleRightRotate(node *redBlackNode) *redBlackNode {
	isRootNode := (tree.Root == node)
	temp := node
	node = node.Left
	temp.Left = node.Right
	node.Right = temp
	//Update parents
	if temp.Parent != nil{
		temp.Parent.Left = node
	}
	node.Parent = temp.Parent
	temp.Parent = node
	if isRootNode{
		tree.Root = node
	}
	return node
}

//Remove operation
func (tree *redBlackTree) Remove(data int) {

}
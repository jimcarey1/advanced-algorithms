//The algorithm starts from the root and performs a binary search comparing the search key
//with the keys stored in the root until it finds the first separator key is greater than
//the searched value. This locates a searched subtree. As we've discussed previously, index
//keys split the tree into subtrees with boundaries between two neighboring keys. As soon
//as we find the subtree, we follow the pointer that corresponds to it and continue the same
//search process(locate the separator key and follow the pointer) until we reach a target leaf
//node, where we either find the searched key or conclude it is not present by locating
//its predecessor.

package btree

import (
	"slices"
)

type bTree struct {
	Root   *bTreeNode
	fanout int
}

// This creates an tree with root node initialized.
func Constructor(fanout int) *bTree {
	return &bTree{
		Root:   newBTreeNode(),
		fanout: fanout,
	}
}

func (bt *bTree) Insert(data int) {
	//Our goal is to reach the leaf node.
	node := bt.Root
	var index int
	for !node.LeafNode() {
		index = binarySearch(node.Keys, data)
		node = node.Children[index]
	}

	var rootNodeEffected bool
	//This for loop runs as long as root is not affected
	//and the node is not nill.
	for node != nil && !rootNodeEffected {
		node, data, rootNodeEffected = node.split(data, index, bt.fanout)
	}
	//If the root node is effected, then update the rootnode.
	if rootNodeEffected {
		bt.Root = node
	}
}

func (bt *bTree) Remove(data int) {
	node := bt.Root
	var index int
	for !node.LeafNode() {
		index = binarySearch(node.Keys, data)
		node = node.Children[index]
	}
	removeIndex := searchIndex(node.Keys, data)
	//If the index is -1, then the element we are trying to
	//search is not there in the tree.
	if removeIndex != -1 {
		node.Keys = slices.Delete(node.Keys, removeIndex, removeIndex+1)
		side := checkUnderFlow(node, bt.fanout)
		for side != "" {
			if side == "next" {
				node = node.Merge(bt.fanout)
			} else {
				node = node.PrevSibling.Merge(bt.fanout)
			}
			side = checkUnderFlow(node, bt.fanout)
		}
	}
	//We have reached the root node and actually should check if the root node should be updated.
	if node.Parent == nil && len(node.Keys) == 0 {
		bt.Root = node.Children[0]
	}
}

func (bn *bTreeNode) Merge(fanout int) *bTreeNode {
	bn.Keys = append(bn.Keys, bn.NextSibling.Keys...)
	bn.Children = append(bn.Children, bn.NextSibling.Children...)
	index := indexOfChildPointer(bn, bn.Parent)

	if bn.LeafNode() {
		bn.Parent.Keys = slices.Delete(bn.Parent.Keys, index, index+1)
	} else {
		insertNewKey(&bn.Keys, bn.Parent.Keys[index])
		bn.Parent.Keys = slices.Delete(bn.Parent.Keys, index, index+1)
	}
	//Before deleting the right node, I have to update the
	//sibling nodes.
	if bn.NextSibling != nil && bn.NextSibling.NextSibling != nil {
		bn.NextSibling.NextSibling.PrevSibling = bn
	}
	bn.Parent.Children = slices.Delete(bn.Parent.Children, index+1, index+2)
	return bn.Parent
}

func (bt *bTree) Search(data int) bool {
	node := bt.Root
	var index int
	for !node.LeafNode() {
		index = binarySearch(node.Keys, data)
		node = node.Children[index]
	}
	return false
}

func (bn *bTreeNode) split(data, index, fanout int) (*bTreeNode, int, bool) {
	if len(bn.Keys) < fanout {
		insertNewKey(&bn.Keys, data)
		return nil, 0, false
	} else {
		//This is the root node.
		if bn.Parent == nil {
			return handleRootNodeSplit(bn, data), 0, true
		} else {
			if len(bn.Parent.Children) < fanout+1 {
				bn, newNode, medianData := handleNodeSplit(bn, data)
				newNode.Parent = bn.Parent
				insertNewChildren(&bn.Parent.Children, index+1, newNode)
				return bn.Parent, medianData, false
			} else if len(bn.Parent.Children) == fanout+1 {
				//This is the children node that we should split.
				children := GetAllChildren(bn.Parent, data, index)
				node := bn.Parent
				keys := node.Keys
				splitPoint := len(keys) / 2
				var isRoot bool = node.RootNode()

				//Create internal nodes and share the children between these
				//two nodes.
				newNode := newBTreeNode()

				//Update the parent and sibling pointers.
				newNode.Parent = node.Parent
				newNode.PrevSibling = node
				newNode.NextSibling = node.NextSibling
				node.NextSibling = newNode

				if data < node.Keys[0] {
					node.Keys = node.Keys[:splitPoint]
					//This step is wrong, we have to get the key of the
					//first child node that is greater than data.
					insertNewKey(&node.Keys, GetCorrectKeyToInsert(children, data))
					node.Children = children[:len(node.Keys)+1]
					newNode.Children = children[len(node.Keys)+1:]
					for _, child := range newNode.Children {
						child.Parent = newNode
						newNode.Keys = append(newNode.Keys, child.Keys[0])
					}
				} else {
					node.Keys = node.Keys[:splitPoint]
					node.Children = children[:len(node.Keys)+1]
					newNode.Children = children[len(node.Children):]
					for _, child := range newNode.Children {
						child.Parent = newNode
						newNode.Keys = append(newNode.Keys, child.Keys[0])
					}
				}
				var element int = newNode.Keys[0]
				newNode.Keys = newNode.Keys[1:]
				if isRoot {
					//Create the upper node, that contains these two nodes as children.
					upperNode := newBTreeNode()
					upperNode.Parent = bn.Parent.Parent
					upperNode.Keys = append(upperNode.Keys, element)

					//Update the parent of the two split nodes as the upper node.
					newNode.Parent = upperNode
					bn.Parent.Parent = upperNode

					//Update the upper node's children.
					upperNode.Children = append(upperNode.Children, node)
					upperNode.Children = append(upperNode.Children, newNode)
					return upperNode, 0, true
				} else {
					return node.Parent, element, false
				}
			}
		}
	}
	return nil, 0, false
}

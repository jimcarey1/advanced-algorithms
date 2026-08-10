package btree

type bTreeNode struct {
	Keys        []int
	Children    []*bTreeNode
	PrevSibling *bTreeNode
	NextSibling *bTreeNode
	Parent      *bTreeNode
}

func newBTreeNode() *bTreeNode {
	return &bTreeNode{}
}

// This function checks if the current node is leaf node.
func (b *bTreeNode) LeafNode() bool {
	return len(b.Children) == 0
}

// This function checks if the current node is root node.
func (b *bTreeNode) RootNode() bool {
	return b.Parent == nil
}
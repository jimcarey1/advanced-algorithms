package main

import (
	"fmt"

	"github.com/jimcarey1/advanced_algorithms/trees"
)

func main() {
	avlTree := trees.NewAVLTree()
	avlTree.Insert(1)
	avlTree.Insert(2)
	avlTree.Insert(3)
	avlTree.Insert(4)
	avlTree.Insert(5)
	avlTree.Insert(2)
	fmt.Println(avlTree.Root)
	fmt.Println(avlTree.Root.Left)
	fmt.Println(avlTree.Root.Left.Left)
	fmt.Println(avlTree.Root.Left.Right)
	fmt.Println(avlTree.Root.Right)
	fmt.Println(avlTree.Root.Right.Right)
}

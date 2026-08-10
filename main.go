package main

import (
	"fmt"

	"github.com/jimcarey1/advanced_algorithms/btree"
)

func main() {
	btree := btree.Constructor(4)
	btree.Insert(155)
	btree.Insert(585)
	btree.Insert(748)
	btree.Insert(781)
	btree.Insert(373)
	btree.Insert(480)
	btree.Insert(797)
	btree.Insert(838)
	btree.Insert(743)
	fmt.Println(btree.Root.Keys)
	fmt.Println(btree.Root.Children[0].Keys)
	fmt.Println(btree.Root.Children[1].Keys)
	fmt.Println(btree.Root.Children[2].Keys)
	node := btree.Root.Children[0]
	for node != nil{
		fmt.Println(node.Keys)
		node = node.NextSibling
	}
}

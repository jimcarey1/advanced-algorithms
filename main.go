package main

import (
	"fmt"

	"github.com/jimcarey1/advanced_algorithms/lists"
)

func main() {
	skipList := lists.NewSkipList()
	skipList.Insert(100)
	PrintSkipList(skipList)
	skipList.Insert(90)
	PrintSkipList(skipList)
	skipList.Insert(80)
	PrintSkipList(skipList)
	skipList.Insert(70)
	PrintSkipList(skipList)
	skipList.Insert(60)
	PrintSkipList(skipList)
	skipList.Insert(50)
	PrintSkipList(skipList)
	skipList.Insert(40)
	PrintSkipList(skipList)
	skipList.Insert(40)
	fmt.Printf("The code is reaching this line.")
	PrintSkipList(skipList)
}

func PrintLevel(node *lists.Node){
	node = node.Next
	for node != nil{
		fmt.Printf("%d->", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func PrintSkipList(list *lists.SkipList){
	for _, node := range list.Levels{
		PrintLevel(node)
	}
	fmt.Println()
}
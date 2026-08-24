package lists

import (
	"math"
	"math/rand"
)

type Node struct {
	Val  int
	Prev *Node
	Next *Node
	Up   *Node
	Down *Node
}

type SkipList struct {
	Levels []*Node
}

func NewSkipList() *SkipList {
	node := &Node{Val: int(math.Inf(-1))}
	return &SkipList{
		Levels: []*Node{node},
	}
}

// This method adds a newlevel to the list.
func (list *SkipList) AddNewLevel() {
	node := &Node{Val: int(math.Inf(-1))}
	list.Levels = append(list.Levels, node)
	//After adding a level, We have to update the Up and Down pointers.
	noOfLevels := len(list.Levels)
	list.Levels[noOfLevels-1].Down = list.Levels[noOfLevels-2]
	list.Levels[noOfLevels-2].Up = list.Levels[noOfLevels-1]
}

// This method adds a new value to the skip list.
func (list *SkipList) Insert(val int) {
	trackPrevLevelSteps := []*Node{}
	currentLevel := len(list.Levels) - 1
	node := list.Levels[currentLevel]
	for currentLevel > 0 {
		if node.Next != nil && node.Next.Val < val {
			node = node.Next
		} else if node.Next == nil || node.Next.Val > val {
			trackPrevLevelSteps = append(trackPrevLevelSteps, node)
			node = node.Down
			currentLevel = currentLevel - 1
		}
	}
	for node.Next != nil && node.Next.Val < val {
		node = node.Next
	}
	// Inserting the value at the base level and,
	// Updating the Previous and Next pointers.
	nextNode := node.Next
	newNode := &Node{Val: val}
	newNode.Prev = node
	newNode.Next = nextNode
	node.Next = newNode
	if nextNode != nil {
		nextNode.Prev = newNode
	}

	//Now, we have to randomize adding the node to upper levels.
	//We will insert, as long as ShouldInsert is true.
	var AtLevel int = 1 //Currently we are at level 1.
	var currentHeight int = len(list.Levels)
	for ShouldInsert() {
		if AtLevel < currentHeight {
			newNode := &Node{Val: val}
			//We need access to previous node at that particular level.
			prevNode := trackPrevLevelSteps[len(trackPrevLevelSteps)-1]
			//Delete the last element from the slice(we longer need it.)
			trackPrevLevelSteps = trackPrevLevelSteps[:len(trackPrevLevelSteps)-1]

			//Insert the node at the level and update the prev and next nodes.
			nextNode := prevNode.Next
			newNode.Next = nextNode
			prevNode.Next = newNode
			newNode.Prev = prevNode
			if nextNode != nil {
				nextNode.Prev = newNode
			}

			//Update the Up and Down pointers.
			newNode.Down = node
			node.Up = newNode
			node = newNode
			AtLevel = AtLevel + 1
		} else {
			//If we are on to add a new level.
			list.AddNewLevel()
			upperNode := list.Levels[len(list.Levels)-1]
			newNode := &Node{Val: val}
			upperNode.Next = newNode
			newNode.Prev = upperNode
			newNode.Down = node
			node.Up = newNode
			node = newNode
		}
	}
}

// This method removes a value from the skip list
func (list *SkipList) Remove(val int) {

}

// This method searches, if an element exists in the skip list.
func (list *SkipList) Search(val int) bool {
	currentLevel := len(list.Levels) - 1
	node := list.Levels[currentLevel]
	for currentLevel > 0 {
		if node.Val == val {
			return true
		} else {
			if node.Next != nil && node.Next.Val < val {
				node = node.Next
			} else if node.Next == nil || node.Next.Val > val {
				node = node.Down
				currentLevel = currentLevel - 1
			}
		}
	}
	for node != nil {
		if node.Val == val {
			return true
		}
		node = node.Next
	}
	return false
}

// This function randomly generates true or false.
func ShouldInsert() bool {
	randomValue := rand.Uint32()
	return randomValue%2 == 0
}

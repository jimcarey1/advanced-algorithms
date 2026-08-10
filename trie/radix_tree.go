package trie

import (
	"fmt"
	"slices"
)

type Trie struct {
	Val       string
	Children  []*Trie
	Completed bool
}

func Constructor() *Trie {
	return &Trie{
		Children: []*Trie{},
	}
}

func (t *Trie) Insert(word string) {
	//If the trie is empty
	if len(t.Children) == 0 {
		t.Children = append(t.Children, &Trie{Val: word, Completed: true, Children: []*Trie{}})
		return
	}

	prefix, trie, _ := getMatching(word, t.Children)
	//If there are some children, but no prefix.
	if trie == nil {
		t.Children = append(t.Children, &Trie{Val: word, Completed: true, Children: []*Trie{}})
		return
	}
	//The trie node is incomplete, the prefix and the word both matches the trie node.
	//The trie node has data "tap"(incomplete), the word to insert is "tap"
	//We simply update the complete property to true.
	if !trie.Completed && prefix == trie.Val && prefix == word {
		trie.Completed = true
		return
	}
	//The trie node is a complete word, and has a prefix with the word to insert.
	//trie node has data "ten" and completed and I want to insert "tea"
	//The prefix is "te"(not completed) and the children are "a"(compelted) and "n"(completed)
	if trie.Completed && trie.Val != prefix && word != prefix {
		temp := trie.Val
		tempChildren := trie.Children
		fmt.Println(tempChildren)
		fmt.Println(len(tempChildren))
		trie.Completed = false
		trie.Val = prefix
		trie.Children = trie.Children[:0]
		trie.Children = append(trie.Children, &Trie{Val: temp[len(prefix):], Completed: true, Children: tempChildren})
		trie.Children = append(trie.Children, &Trie{Val: word[len(prefix):], Completed: true, Children: []*Trie{}})
		fmt.Println(trie.Children)
		return
	}

	//If trie node is not completed but shares a prefix with the word to insert.
	//trie node has data "te" and not completed, I want to insert "top"
	//The prefix is t and the children are "e"(not completed) and "op"(completed)
	if !trie.Completed && prefix != word && prefix != trie.Val {
		tempVal := trie.Val
		tempChildern := trie.Children
		trieNode := &Trie{Val: tempVal[len(prefix):], Completed: false, Children: []*Trie{}}
		trieNode.Children = tempChildern
		trie.Val = prefix
		trie.Children = trie.Children[:0]
		trie.Children = append(trie.Children, &Trie{Val: word[len(prefix):], Completed: true, Children: []*Trie{}})
		trie.Children = append(trie.Children, trieNode)
		return
	}

	//The trie node is a completed word, and prefix matches the trie node but the word does not.
	//trie node has data "tap" and completed, I want to insert "tape"
	//The prefix is "tap"(completed) and the children will be "tape"(completed)
	if trie.Completed && trie.Val == prefix && word != prefix {
		if len(trie.Children) == 0 {
			trie.Children = append(trie.Children, &Trie{Val: word[len(prefix):], Completed: true, Children: []*Trie{}})
			return
		}
	}
	//The trie node is a complete word, and prefix matches the word but no the trie node data(opposite to one at the top)
	//trie node has data "tape" and completed, I want to insert "tap"
	//We will update the existing trie node data to "tap"(completed) and add a children with
	//data "tap(e)" (completed)
	if trie.Completed && prefix == word && prefix != trie.Val {
		tempChildren := trie.Children
		trie.Children = trie.Children[:0]
		trie.Children = append(trie.Children, &Trie{Val: trie.Val[len(prefix):], Completed: true, Children: tempChildren})
		trie.Val = prefix
		return
	}
	//Walk down the trie with the remaining part of the word.
	trie.Insert(word[len(prefix):])

}

func (t *Trie) Remove(word string) {
	prefix, trie, index := getMatching(word, t.Children)
	if trie == nil {
		fmt.Printf("No word named: %s\n", word)
	}
	//If the word and the prefix are the same, and the trie node is complete.
	//If the word is "a", trie node data is "a"(completed), prefix is also a
	//We remove that child, from the slice and update it's parent, If it has
	//no other children, If it has no children we will not do anything.
	if prefix == word && trie.Completed && len(trie.Children) == 0 {
		t.Children = slices.Delete(t.Children, index, index+1)
		if len(t.Children) == 0 {
			t.Val = t.Val + trie.Val
		}
		return
	}
	//If the word and the prefix are the same and the trie node is complete.
	//If the word is "tap", trie node is "tap"(completed), prefix is also "tap"
	//But the "tap" trie node has a children "tap(e)" (completed)
	if prefix == word && trie.Completed && len(trie.Children) > 0 {
		//If there is only one child, we will remove that child and attach its
		//data to the current trie.
		//In the above example, we will remove "tap(e)" child of the "tap",
		//we will want to remove and change the data "tap" node to "tape".
		if len(trie.Children) == 1 {
			tempChildData := trie.Children[0].Val
			trie.Children = trie.Children[:0]
			trie.Val = trie.Val + tempChildData
		} else {
			//If there is more than one node, we do not have to do anything other
			//than setting the node complete attribute to false.
			trie.Completed = false
		}
		return
	}
}

func (t *Trie) Search(word string) bool {
	prefix, trie, _ := getMatching(word, t.Children)
	if trie == nil {
		return false
	}
	if prefix == word && prefix == trie.Val && !trie.Completed {
		return false
	}
	if prefix == word && prefix == trie.Val && trie.Completed {
		return true
	}
	return trie.Search(word[len(prefix):])
}

func (t *Trie) StartsWith(prefix string) bool {
	prefix2, trie, _ := getMatching(prefix, t.Children)
	if trie == nil {
		return false
	}
	if prefix2 == prefix {
		return true
	}
	return trie.StartsWith(prefix[len(prefix2):])
}

func getMatching(word string, children []*Trie) (string, *Trie, int) {
	for i := 0; i < len(children); i++ {
		prefix := prefix(word, children[i].Val)
		if prefix != "" {
			return prefix, children[i], i
		}
	}
	return "", nil, -1
}

func prefix(s1 string, s2 string) string {
	result := make([]byte, 0)
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			break
		} else {
			result = append(result, s1[i])
		}
	}
	return string(result)
}

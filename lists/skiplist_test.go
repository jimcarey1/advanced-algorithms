package lists

import (
	"testing"
)

func TestInsert(t *testing.T) {
	skipList := NewSkipList()
	skipList.Insert(50)
	skipList.Insert(80)
	// skipList.Insert(10)
	// skipList.Insert(1)
	// skipList.Insert(100)
	// skipList.Insert(88)
	node := skipList.Levels[0]
	if node.Next.Val != 50 {
		t.Errorf("Expected 50, got %d.\n", node.Next.Val)
	}
	if node.Next.Next.Val != 80 {
		t.Errorf("Expected 80, got %d\n", node.Next.Next.Val)
	}
}

package binTree

import (
	"testing"
)

func TestCreateAndInsert(t *testing.T) {
	root := Node{ Key: 0, Value: 0 }
	root.Insert(-1, 0)
	root.Insert(1, 1)
	root.Insert(2, 2)
	root.Insert(3, 3)

	if root.Value != 0 {
		t.Errorf("root not 0")
	}
	if root.Right.Value != 1 {
		t.Errorf("child root not 1")
	}
}

func TestSearch(t *testing.T) {
	root := Node{ Key: 0, Value: 0 }
	root.Insert(-1, 0)
	root.Insert(1, 1)
	root.Insert(2, 2)
	root.Insert(3, 3)

	s1 := root.Search(-1)
	if s1.Key != -1 {
		t.Errorf("s1 doesn't have the right key")
	}
	s2 := root.Search(10)
	if s2 != nil {
		t.Errorf("s2 is not nil")
	}
}

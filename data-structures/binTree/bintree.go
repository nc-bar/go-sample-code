package binTree

type Node struct {
	Key	int
	Value	int
	Left	*Node
	Right	*Node
	Parent	*Node
}

func (tree *Node) Insert(key, value int) *Node {
	if tree == nil {
		return &Node {
			Key: key,
			Value: value,
		}
	}
	// tree != nil
	if key < tree.Key {
		if tree.Left == nil {
			tree.Left = &Node {
				Key: key,
				Value: value, //fix this
				Parent: tree,
			}
			return tree.Left
		} else {
			return tree.Left.Insert(key, value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &Node {
				Key: key,
				Value: value,
				Parent: tree,
			}
			return tree.Right
		} else {
			return tree.Right.Insert(key, value)
		}
	}
	return nil
}

func (tree *Node) Search(key int) *Node {
	if tree == nil {
		return nil 
	}
	// tree != nil
	if tree.Key == key {
		return tree
	}

	if key < tree.Key {
		return tree.Left.Search(key)
	} else {
		return tree.Right.Search(key)
	}
}

// not tested | not debugged
func (tree *Node) Delete(key int) {
	toDelete := tree.Search(key)
	if toDelete == nil {
		return 
	}

	// node to delete found, search for the smallest
	// key at it's right
	r := toDelete.Right
	for r.Left != nil {
		r = r.Left
	}

	// set r where toDelete is
	if r.Parent != nil {
		// r can't be root?
		r.Parent.Left= r.Right
	}
	r.Parent = toDelete.Parent
	r.Left = toDelete.Left
	r.Right = toDelete.Right
}

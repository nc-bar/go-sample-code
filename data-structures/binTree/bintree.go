package binTree

type Node struct {
	Key		int
	Value		int
	Left, Right	*Node
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


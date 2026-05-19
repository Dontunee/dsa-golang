package dataStructure

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Search checks if the current node is nil and returns false
// checks if the search parameter is greater than the current node key and move to the right node
// else if the search parameter is less than the current node key , move to the left
// if not it is equal so return true
func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}
	if key > n.Key {
		return n.Right.Search(key)
	} else if key < n.Key {
		return n.Left.Search(key)
	}
	return true
}

//Insert checks if value to be added is greater than current node value
// perform nil check and insert a new node else perform insert operation on the right node
//do the same for the left node if the value to be added is less than the current node value
//if it is the same value perform insert operation on left value
func (n *Node) Insert(key int) {
	if key > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	} else if key < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	} else {
		n.Left.Insert(key)
	}
}

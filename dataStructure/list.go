package dataStructure

type list struct {
	data int
	next *list
}

func (l *list) Search(x int) bool {
	if l == nil {
		return false
	}

	if l.data == x {
		return true
	}
	return l.next.Search(x)
}

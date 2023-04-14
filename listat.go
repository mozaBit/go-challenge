package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	if l.Next == nil {
		return l
	}
	iter := l
	for i := 1; i <= pos; i++ {
		if iter.Next == nil {
			return nil
		}
		iter = iter.Next
	}
	return iter
}

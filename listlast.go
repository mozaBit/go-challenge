package piscine

func ListLast(l *List) interface{} {
	if l.Head != nil {
		iter := l.Head
		for ; iter.Next != nil; iter = iter.Next {
		}
		return iter.Data
	} else {
		return nil
	}
}

package piscine

func ListSize(l *List) int {
	count := 0
	if l.Head == nil {
		return 0
	} else {
		for iter := l.Head; iter != nil; iter = iter.Next {
			count += 1
		}
	}
	return count
}

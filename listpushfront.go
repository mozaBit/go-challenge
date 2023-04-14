package piscine

func ListPushFront(l *List, data interface{}) {
	x := new(NodeL)
	x.Data = data
	if l.Head != nil {
		x.Next = l.Head
		l.Head = x
	}
	if l.Head == nil {
		l.Head = x
		x.Next = nil
	}
}

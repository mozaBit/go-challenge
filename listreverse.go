package piscine

func ListReverse(l *List) {
	var prev *NodeL = nil
	if l.Head == nil {
	} else {
		for l.Head != nil {
			next_node := l.Head.Next
			l.Head.Next = prev
			prev = l.Head
			l.Head = next_node
		}
		l.Head = prev
		iter := l.Head
		for ; iter.Next != nil; iter = iter.Next {
		}
		l.Tail = iter
	}
}

// func ListReverse(l *List) {
// 	l2 := new(List)
// 	l3 := new(NodeL)
// 	iter := l.Head
// 	for i := 1; i <= ListSize2(l); i++ {
// 		for j := ListSize2(l) - 1; j >= 1; j-- {
// 			if l2.Head == nil {
// 				l2.Head = iter
// 				l3 = l2.Head
// 			} else {
// 				l3.Next = iter
// 			}
// 		}
// 		iter = iter.Next
// 	}
// 	l3.Next = nil
// 	l2.Tail = l3
// 	l = l2
// }

// func ListSize2(l *List) int {
// 	count := 0
// 	if l.Head == nil {
// 		return 0
// 	} else {
// 		for iter := l.Head; iter != nil; iter = iter.Next {
// 			count += 1
// 		}
// 	}
// 	return count
// }

// func ListReverse(l *List) {
// 	for i := 0; i < ListSize(l); i++ {
// 		jter := l.Head
// 		for ; jter.Next != nil; jter = jter.Next {
// 		}
// 		x := new(NodeL)
// 		x.Data = l.Head.Data
// 		x.Next = nil
// 		jter.Next = x
// 		l.Head = l.Head.Next
// 	}
// 	iter := l.Head
// 	for ; iter.Next != nil; iter = iter.Next {
// 	}
// 	l.Tail = iter
// }

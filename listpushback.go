package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	x := new(NodeL)
	x.Data = data
	x.Next = nil
	if l.Head == nil {
		l.Head = x
	} else {
		iP := l.Head
		for ; iP.Next != nil; iP = iP.Next {
		}
		iP.Next = x
	}
}

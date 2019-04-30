package link

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type Link struct {
	head *Node
	tail *Node
}

func (l *Link) InsertHead(data interface{}) {
	node := &Node{
		data:data,
		next:nil,
	}
	if l.tail == nil && l.head == nil {
		l.head = node
		l.tail = node
		return
	}
	node.next = l.head
	l.head = node
}

func (l *Link) InsertTail(data interface{}) {
	node := &Node{
		data:data,
		next:nil,
	}
	if l.tail == nil && l.head == nil {
		l.head = node
		l.tail = node
		return
	}
	l.tail.next = node
	l.tail = node
}

func (l *Link) Trans() {
	p := l.head
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}

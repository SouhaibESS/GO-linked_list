package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

type list struct {
	first *node
	last  *node
	len   int
}

func (l *list) add(value int) {
	newNode := new(node)
	newNode.value = value

	if l.len == 0 {
		l.first = newNode
		l.last = newNode
	} else if l.len == 1 {
		l.last = newNode
		l.first.next = l.last
	} else {
		l.last.next = newNode
		l.last = newNode
	}
	l.len++
}

func (l list) get(value int) *node {
	if l.first == nil {
		return nil
	}
	elmnt := l.first
	for {
		if elmnt == nil {
			return nil
		}
		if value != elmnt.value {
			elmnt = elmnt.next
		} else {
			return elmnt
		}
	}
}

func (l *list) remove(value int) {
	if l.first == nil {
		fmt.Println("The list is empty!")
		return
	}

	if l.first.value == value {
		l.first = l.first.next
	}

	elmnt := l.first
	for {
		if elmnt.next == nil {
			return
		}
		if elmnt.next.value == value {
			elmnt.next = elmnt.next.next
			break
		}
		elmnt = elmnt.next
	}
}

func (l list) String() string {
	ret := strings.Builder{}
	elmnt := l.first
	for {
		if elmnt != nil {
			ret.WriteString(fmt.Sprintf("%d -> ", elmnt.value))
			elmnt = elmnt.next
		} else {
			ret.WriteString(fmt.Sprintf("nil."))
			break
		}
	}

	return ret.String()
}

func main() {
	l := list{}
	l.add(1)
	l.add(2)
	l.add(3)

	l.remove(3)

	fmt.Println(l)
}

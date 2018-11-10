package client

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/ottojo/blnk2/vector"
	"net"
)

type Optimization int

const (
	NONE             Optimization = iota
	PREVENTDUPLICATE Optimization = iota
	PACKETSIZE       Optimization = iota
)

type Client struct {
	Length        int
	FirstLed      *LedListElement
	FirstStageLed *LedListElement
	Socket        net.Conn
	Opt           Optimization
	Address       string
	Id            string
	Connecting    bool
}

type LedList struct {
	First *LedListElement
}

func (l *LedList) Length() int {
	e := l.First
	i := 0
	for e != nil {
		i++
		e = e.Next
	}
	return i
}

func (l *LedList) get(i int) *LedListElement {
	element := l.First
	for b := 0; b < i; b++ {
		element = element.Next
	}
	return element
}

func (l *LedList) Add(led Led) *LedListElement {
	element := LedListElement{Data: led}
	if l.First == nil {
		l.First = &element
	} else {
		last := l.First
		for last.Next != nil {
			last = last.Next
		}
		last.Next = &element
	}
	return &element
}

func (l *LedListElement) GetOffset(i int) *LedListElement {
	element := l
	for b := 0; b < i; b++ {
		element = element.Next
	}
	return element
}

type LedListElement struct {
	Data Led
	Next *LedListElement
}

type Led struct {
	Color    colorful.Color
	Position vector.Vec3
}

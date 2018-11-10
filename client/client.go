package client

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/ottojo/blnk2/vector"
	"log"
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
	connecting    bool
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

func (c *Client) Connect() {
	if c.Socket == nil && c.Address != "" && !c.connecting {
		c.connecting = true
		var err error
		log.Printf("%s connecting to %s\n", c.Id, c.Address)
		c.Socket, err = net.Dial("tcp", c.Address)
		LogError(err)
		c.connecting = false

	}
}

func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}

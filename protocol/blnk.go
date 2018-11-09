package protocol

import (
	"github.com/ottojo/blnk/color"
	"github.com/ottojo/blnk/neoPixel"
	"log"
	"net"
)

func SetPixelsMessage(colors []color.Color8bit) BlnkPacket {
	var message = make([]byte, 3*len(colors))
	for i, c := range colors {
		message[3*i] = c.R
		message[3*i+1] = c.G
		message[3*i+2] = c.B

	}
	message = append([]byte{1}, message...)
	return newPacket(message)
}

func SetIntervalMessage(colors []color.Color8bit, startIndex int) BlnkPacket {
	var message = make([]byte, 3*len(colors)+2)
	message[0] = byte(startIndex >> 8)
	message[1] = byte(startIndex)
	for i, c := range colors {
		message[3*i+2] = c.R
		message[3*i+3] = c.G
		message[3*i+4] = c.B
	}
	message = append([]byte{2}, message...)
	return newPacket(message)
}

type BlnkPacket []byte

func newPacket(payload []byte) BlnkPacket {
	return append([]byte{0xaf, byte(len(payload) >> 8), byte(len(payload))}, payload...)
}

type NeoClient struct {
	ID      string
	Address string
	Conn    net.Conn
	Strip   neoPixel.Strip
}

func (n *NeoClient) Commit() {
	if n.Conn == nil {
		var err error
		n.Conn, err = net.Dial("tcp", n.Address)
		if err != nil {
			log.Println(err)
		}
	}

	colors := make([]color.Color8bit, len(n.Strip.NeoPixels))
	for i, n := range n.Strip.NeoPixels {
		colors[i] = n.Color()
	}
	m := SetPixelsMessage(colors)
	_, err := n.Conn.Write(m)
	if err != nil {
		log.Println(err)
	}
}

func (n *NeoClient) Disconnect() {
	if n.Conn != nil {
		n.Conn.Close()
	}
}


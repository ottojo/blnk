package blnkProtocol

import (
	"github.com/ottojo/blnkServer/color"
	"github.com/ottojo/blnkServer/neoPixel"
	"log"
	"net"
)

func setPixelsMessage(colors []color.Color8bit) []byte {
	var message = make([]byte, 3*len(colors))
	for i, c := range colors {
		message[3*i] = c.R
		message[3*i+1] = c.G
		message[3*i+2] = c.B

	}
	message = append([]byte{1}, message...)
	return append([]byte{0xaf, byte(len(message) >> 8), byte(len(message))}, message...)
}

type NeoClient struct {
	ID      string
	Address string
	Conn    net.Conn
	Strip   neoPixel.NeoPixelStrip
}

func (n NeoClient) Commit() {
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
	m := setPixelsMessage(colors)
	n.Conn.Write(m)
}

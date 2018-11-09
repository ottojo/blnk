package blnk

import (
	"github.com/ottojo/blnk/color"
	"github.com/ottojo/blnk/neoPixel"
	"github.com/ottojo/blnk/protocol"
	"github.com/ottojo/blnk/storage"
	"io/ioutil"
	"log"
	"net"
)

type BlnkSystem struct {
	clients      []*protocol.NeoClient
	PixelStage   []*neoPixel.NeoPixel
	Optimization Optimization
}








type Optimization int

const (
	NONE       Optimization = iota
	PACKETSIZE Optimization = iota
)

func CreateBlnkSystem(config []byte, o Optimization) (BlnkSystem, error) {
	var b BlnkSystem
	var err error
	b.clients, err = storage.DecodeClients(config)
	if err != nil {
		return b, err
	}

	b.Optimization = o

	for clientIndex := range b.clients {
		for pixelIndex := range b.clients[clientIndex].Strip.NeoPixels {
			b.PixelStage = append(b.PixelStage, &b.clients[clientIndex].Strip.NeoPixels[pixelIndex])
		}
	}

	return b, nil
}

func LoadBlnkSystemFromFile(filename string) (BlnkSystem, error) {
	var b BlnkSystem
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		return b, err
	}
	b, err = CreateBlnkSystem(d)
	return b, err
}

func (b *BlnkSystem) Commit() {

	for _, n := range b.clients {
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

		switch b.Optimization {
		case NONE:
			m := protocol.SetPixelsMessage(colors)
			_, err := n.Conn.Write(m)
			if err != nil {
				log.Println(err)
			}
			break
		case PACKETSIZE:

			var startIndex int

			for ; ; {

			}

			break
		}

	}
}

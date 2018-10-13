package blnk

import (
	"github.com/ottojo/blnk/neoPixel"
	"github.com/ottojo/blnk/protocol"
	"github.com/ottojo/blnk/storage"
	"io/ioutil"
)

type BlnkSystem struct {
	Clients []*protocol.NeoClient
	Pixels  []*neoPixel.NeoPixel
}

func CreateBlnkSystem(config []byte) (BlnkSystem, error) {
	var b BlnkSystem
	var err error
	b.Clients, err = storage.DecodeClients(config)
	if err != nil {
		return b, err
	}

	for clientIndex := range b.Clients {
		for pixelIndex := range b.Clients[clientIndex].Strip.NeoPixels {
			b.Pixels = append(b.Pixels, &b.Clients[clientIndex].Strip.NeoPixels[pixelIndex])
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

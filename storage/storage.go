package storage

import (
	"encoding/json"
	"github.com/ottojo/blnk/protocol"
	"github.com/ottojo/blnk/neoPixel"
	"github.com/ottojo/blnk/vector"
	"io/ioutil"
)

type ClientFile []struct {
	ID             string      `json:"id"`
	Address        string      `json:"address"`
	StartPosition  vector.Vec3 `json:"startPosition"`
	EndPosition    vector.Vec3 `json:"endPosition"`
	PixelsPerMeter float64     `json:"pixelsPerMeter"`
}

// TODO: assign host+port automatically if running with simulator

/*
func ReadSimClientsFromFile(filename string) ([]protocol.NeoClient, error) {
	cls, err := ReadClientsFromFile(filename)
	if err != nil {
		return nil, err
	}
	for i, _ := range cls {
		cls[i].Address="localhost:"+port
	}
}*/

func ReadClientsFromFile(filename string) ([]*protocol.NeoClient, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c, err := DecodeClients(f)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func DecodeClients(data []byte) ([]*protocol.NeoClient, error) {
	var c ClientFile
	err := json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	var clients []*protocol.NeoClient
	for _, cl := range c {
		newClient := protocol.NeoClient{ID: cl.ID, Address: cl.Address,
			Strip: neoPixel.NewNeoPixelStrip(cl.StartPosition, cl.EndPosition, cl.PixelsPerMeter)}
		clients = append(clients, &newClient)
	}
	return clients, nil
}

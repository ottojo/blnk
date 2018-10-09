package main

import (
	"encoding/json"
	"github.com/ottojo/blnkServer/blnkProtocol"
	"github.com/ottojo/blnkServer/neoPixel"
	"github.com/ottojo/blnkServer/vector"
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
func ReadSimClientsFromFile(filename string) ([]blnkProtocol.NeoClient, error) {
	cls, err := ReadClientsFromFile(filename)
	if err != nil {
		return nil, err
	}
	for i, _ := range cls {
		cls[i].Address="localhost:"+port
	}
}*/

func ReadClientsFromFile(filename string) ([]blnkProtocol.NeoClient, error) {
	var c ClientFile
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(f, &c)
	if err != nil {
		return nil, err
	}
	var clients []blnkProtocol.NeoClient
	for _, cl := range c {
		newClient := blnkProtocol.NeoClient{ID: cl.ID, Address: cl.Address,
			Strip: neoPixel.NewNeoPixelStrip(cl.StartPosition, cl.EndPosition, cl.PixelsPerMeter)}
		clients = append(clients, newClient)
	}
	return clients, nil
}

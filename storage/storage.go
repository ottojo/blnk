package storage

import (
	"encoding/json"
	"github.com/ottojo/blnk2/client"
	"github.com/ottojo/blnk2/vector"
	"io/ioutil"
)

type ClientFile []struct {
	ID     string `json:"id"`
	Strips []struct {
		StartPosition  vector.Vec3 `json:"startPosition"`
		EndPosition    vector.Vec3 `json:"endPosition"`
		PixelsPerMeter int         `json:"pixelsPerMeter"`
	} `json:"strips"`
}

func ReadClientsFromFile(filename string) ([]client.Client, [][]client.Led, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	c, leds, err := DecodeClients(f)
	if err != nil {
		return nil, nil, err
	}

	return c, leds, nil
}

func DecodeClients(data []byte) ([]client.Client, [][]client.Led, error) {
	var c ClientFile
	err := json.Unmarshal(data, &c)
	if err != nil {
		return nil, nil, err
	}
	var clients []client.Client
	var leds [][]client.Led
	for _, cl := range c {
		var clientLeds []client.Led
		var clientLength = 0
		for _, s := range cl.Strips {
			var meterPerPixel float64 = 1 / float64(s.PixelsPerMeter)
			sLength := int(s.EndPosition.Minus(s.StartPosition).Length() * float64(s.PixelsPerMeter))
			clientLength += sLength
			for i := 0; i < sLength; i++ {
				clientLeds = append(clientLeds, client.Led{
					Position: s.StartPosition.Plus(
						s.EndPosition.Minus(s.StartPosition).Normalize().Times(meterPerPixel * float64(i)))})
			}
		}
		newClient := client.Client{Length: clientLength, Id: cl.ID}
		leds = append(leds, clientLeds)
		clients = append(clients, newClient)
	}
	return clients, leds, nil
}

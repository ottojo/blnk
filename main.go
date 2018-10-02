package main

import (
	"github.com/ottojo/blnkServer/color"
	"log"
)

func main() {

	clients, err := ReadClientsFromFile("/home/jonas/sketchbook/blnkClientSim/clients.json")
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range clients {
		c.SetLeds(color.FillColor(c.Strip.Length(), color.Color8bit{230, 79, 45}))
	}

}

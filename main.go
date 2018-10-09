package main

import (
	"github.com/ottojo/blnkServer/blnkProtocol"
	"github.com/ottojo/blnkServer/color"
	"github.com/ottojo/blnkServer/vector"
	"log"
	"math"
)

func main() {
	clients, err := ReadClientsFromFile("C:\\Users\\Jonas Otto\\Documents\\toolbox\\blnkClientSim\\clients.json")
	if err != nil {
		log.Fatal(err)
	}
	/*for _, c := range clients {
		h := 0
		step := 360.0 / len(c.Strip.NeoPixels)
		for i := range c.Strip.NeoPixels {
			col := colorful.Hsv(float64(h), 1, 1)
			r, g, b := col.LinearRgb()
			h += step
			c.Strip.NeoPixels[i].SetColor(
				color.Color8bit{byte(r * 255), byte(g * 255), byte(b * 255)})
		}
		c.Commit()
	}*/

	// G B
	// R T
	RenderBitmap(
		[][]color.Color8bit{
			{color.Color8bit{0, 255, 0}, color.Color8bit{255, 0, 0}},
			{color.Color8bit{0, 0, 255}, color.Color8bit{0, 255, 255}},
		},
		clients,
		vector.Vec3{0, 0.2, 0}, vector.Vec3{1.5, 1, 0.3},
		0.3*math.Pi)
	for _, c := range clients {
		c.Commit()
	}
}

// Right handed coordinate system, Z=UP
// TODO: Find out why image is mirrored
func RenderBitmap(bmp [][]color.Color8bit, clients []blnkProtocol.NeoClient, pStart vector.Vec3, pDir vector.Vec3, horizontalAngle float64) {
	w := 2 * math.Tan(horizontalAngle/2)
	h := w * (float64(len(bmp[0])) / float64(len(bmp)))

	for clientIndex := 0; clientIndex < len(clients); clientIndex++ {
		for pixelIndex := 0; pixelIndex < len(clients[clientIndex].Strip.NeoPixels); pixelIndex++ {

			pixelPos := clients[clientIndex].Strip.NeoPixels[pixelIndex].GetPosition()

			pLoc := pixelPos.Minus(pStart)

			phi := pLoc.Phi() - pDir.Phi()

			if phi < 0 {
				phi += 2 * math.Pi
			}

			plTheta := pLoc.Theta()
			pdTheta := pDir.Theta()
			theta := plTheta - pdTheta
			if theta < 0 {
				theta += math.Pi
			}

			x := (w / 2) - math.Tan(phi)
			x = float64(len(bmp)) * (x / w)

			y := (h / 2) - math.Tan(theta)
			y = float64(len(bmp[0])) * (y / h)

			if x < 0 || y < 0 || int(x) >= len(bmp) || int(y) >= len(bmp[0]) {
				continue
			}
			clients[clientIndex].Strip.NeoPixels[pixelIndex].SetColor(bmp[int(x)][int(y)])
		}
	}
}

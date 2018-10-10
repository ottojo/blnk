package main

import (
	"fmt"
	"github.com/ottojo/blnkServer/blnkProtocol"
	"github.com/ottojo/blnkServer/color"
	"github.com/ottojo/blnkServer/vector"
	"image"
	"image/gif"
	"log"
	"math"
	"os"
	"time"
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

	f, err := os.Open("g.gif")

	g, err := gif.DecodeAll(f)

	for true {
		for i, im := range g.Image {
			fmt.Println(i)
			go RenderBitmap(im, clients, vector.Vec3{-5, -5, 1.8},
				vector.Vec3{1, 1, 0.1},
				0.25*math.Pi)

			time.Sleep(time.Duration(g.Delay[i]*10 ) * time.Millisecond)
		}
	} /*

	// G B
	// R T
	RenderBitmap(
		[][]color.Color8bit{
			{color.Color8bit{0, 255, 0}, color.Color8bit{255, 0, 0}},
			{color.Color8bit{0, 0, 255}, color.Color8bit{0, 255, 255}},
		},
		clients,
		vector.Vec3{-5, -5, 1.8},
		vector.Vec3{1, 1, 0.1},
		0.25*math.Pi)
	for _, c := range clients {
		c.Commit()
	}*/
}

// Right handed coordinate system, Z=UP
// TODO: Find out why image is mirrored
func RenderBitmap(bmp image.Image, clients []blnkProtocol.NeoClient, pStart vector.Vec3, pDir vector.Vec3, horizontalAngle float64) {

	bmp.Bounds()

	imageWidth := bmp.Bounds().Max.X - bmp.Bounds().Min.X
	imageHeight := bmp.Bounds().Max.Y - bmp.Bounds().Min.Y

	w := 2 * math.Tan(horizontalAngle/2)
	h := w * (float64(imageHeight) / float64(imageWidth))

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
			x = float64(imageWidth) * (x / w)

			y := (h / 2) - math.Tan(theta)
			y = float64(imageHeight) * (y / h)

			if x < 0 || y < 0 || int(x) >= imageWidth || int(y) >= imageHeight {
				continue
			}
			r, g, b, _ := bmp.At(int(x), int(y)).RGBA()
			rB := byte((float64(r) / float64(0xffff)) * 255)
			gB := byte((float64(g) / float64(0xffff)) * 255)
			bB := byte((float64(b) / float64(0xffff)) * 255)

			clients[clientIndex].Strip.NeoPixels[pixelIndex].SetColor(color.Color8bit{R: rB, G: gB, B: bB})
		}
		clients[clientIndex].Commit()
	}
}

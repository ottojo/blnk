package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/ottojo/blnkServer/color"
	"log"
)

func main() {
	clients, err := ReadClientsFromFile("/home/jonas/Projects/ProcessingTest/clients.json")
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range clients {
		h := 0
		step := 360.0 / len(c.Strip.NeoPixels)
		for i, _ := range c.Strip.NeoPixels {
			col := colorful.Hsv(float64(h), 1, 1)
			r, g, b := col.LinearRgb()
			h += step
			c.Strip.NeoPixels[i].SetColor(
				color.Color8bit{byte(r * 255), byte(g * 255), byte(b * 255)})
		}
		c.Commit()
	}
}
/*
//Right handed coordinate system, Z=UP
func RenderBitmap(bmp [][]color.Color8bit, clients []blnkProtocol.NeoClient, projectionStart vector.Vec3, projectionDirectionCenter vector.Vec3, horizontalAngle float64) {

	var anglePerPixel = horizontalAngle / float64(len(bmp))

	for _, client := range clients {
		for _, pixel := range client.Strip.NeoPixels {

			// Translate everything to origin

			pixelLocation := pixel.GetPosition().Minus(projectionStart)

			// Rotate projectionDirection to Ex^

			// Rotate by phi around z axis, so y=0

			phi := -math.Atan2(projectionDirectionCenter.Y, projectionDirectionCenter.X)

			// Rotate by theta around y axis, so z=0

			theta := (math.Pi / 2) - math.Acos(projectionDirectionCenter.Z/projectionDirectionCenter.Length())

			// So the same to pixelLocation.

			pixelLocation = pixelLocation.RotateZ(phi).RotateY(theta)

		}
	}

}
*/
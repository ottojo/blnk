package neoPixel

import (
	"github.com/ottojo/blnkServer/color"
	"github.com/ottojo/blnkServer/vector"
)

type NeoPixelStrip struct {
	StartPosition      vector.Vec3
	EndPosition        vector.Vec3
	NeoPixels          []NeoPixel
	interPixelDistance float64
}

func NewNeoPixelStrip(start, end vector.Vec3, ppm float64) NeoPixelStrip {
	n := NeoPixelStrip{
		interPixelDistance: 1 / ppm, StartPosition: start, EndPosition: end}

	n.NeoPixels = make([]NeoPixel, int(end.Minus(start).Length()*ppm/100))

	return n
}

func (s NeoPixelStrip) Length() int {
	return len(s.NeoPixels)
}

type NeoPixel struct {
	i     int
	color color.Color8bit
	s     *NeoPixelStrip
}

func (n *NeoPixel) SetColor(c color.Color8bit) {
	n.color = c
}

func (n NeoPixel) Color() color.Color8bit {
	return n.color
}

func (n NeoPixel) GetPosition() vector.Vec3 {
	return n.s.StartPosition.Plus(
		n.s.EndPosition.Minus(n.s.StartPosition).
			ScaleToLength(n.s.interPixelDistance).
			Times(float64(n.i)))
}

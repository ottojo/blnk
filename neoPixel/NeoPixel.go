package neoPixel

import (
	"github.com/ottojo/blnk/color"
	"github.com/ottojo/blnk/vector"
)

type Strip struct {
	StartPosition      vector.Vec3
	EndPosition        vector.Vec3
	NeoPixels          []NeoPixel
	interPixelDistance float64
}

func NewNeoPixelStrip(start, end vector.Vec3, ppm float64) Strip {
	n := Strip{
		interPixelDistance: 1 / ppm, StartPosition: start, EndPosition: end}

	n.NeoPixels = make([]NeoPixel, int(end.Minus(start).Length()*ppm))
	for i := 0; i < len(n.NeoPixels); i++ {
		n.NeoPixels[i].s = &n
		n.NeoPixels[i].i = i
	}

	return n
}

func (s Strip) Length() int {
	return len(s.NeoPixels)
}

type NeoPixel struct {
	color color.Color8bit
	next  *NeoPixel
}

type NeoPixelList struct {
	first  *NeoPixel
	length int
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

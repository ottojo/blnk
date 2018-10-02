package neoPixel
import "github.com/ottojo/blnkServer/vector"
type NeoPixelStrip struct {
	StartPosition  vector.Vec3
	EndPosition    vector.Vec3
	PixelsPerMeter float64
}

func (s NeoPixelStrip) Length() int {
	return int((s.EndPosition.Minus(s.StartPosition).Length() / 100) * s.PixelsPerMeter)
}


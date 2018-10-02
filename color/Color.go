package color

type Color8bit struct {
	R byte
	G byte
	B byte
}
func FillColor(n int, c Color8bit) []Color8bit {
	var r = make([]Color8bit, n, n)
	for i := 0; i < n; i++ {
		r[i] = c
	}
	return r
}

package main

import "golang.org/x/tour/pic"

// Pic generates a pic.Show compliant [][]unit8 slice.
func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dx)

	for y := range p {
		p[y] = make([]uint8, dy)

		for x := range p[y] {
			p[y][x] = uint8(x^y)
		}
	}

	return p
}

func main() {
	pic.Show(Pic)

	// Result can be viewed at https://codebeautify.org/base64-to-image-converter (minus 'IMAGE:' at the start)
}
package main

import (
	"gba/mode4"
	"image/color"
	"math/rand"
)

var display = Mode4.Display
var palette = Mode4.Palette

func main() {
	display.Configure()

	// fill the palette buffer with random colors
	for i := 1; i < 256; i++ {
		r := uint8(rand.Intn(32))
		g := uint8(rand.Intn(32))
		b := uint8(rand.Intn(32))
		palette.SetColor(uint8(i), color.RGBA{r, g, b, 255})
	}

	page := true
	color := uint8(1)

	for {
		// fill page with current palette index
		for y := 0; y < Mode4.Height; y++ {
			for x := 0; x < Mode4.Width; x++ {
				display.SetPixel(page, x, y, color)
			}
		}
		for !display.VBlank() {
			continue
		} // wait for VBlank
		display.DisplayPage(page) // show page
		page = !page              // flip page
		color += 1                // next color
	}
}

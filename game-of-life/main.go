package main

import (
	"gba/mode4"
	"image/color"
	"math/rand"
)

type Universe struct {
	display       Mode4.DisplayBuffer
	width, height int
	page          bool
}

func (u *Universe) Populate() {
	for i := 0; i < (u.width * u.height / 8); i++ {
		u.display.SetPixel(u.page, rand.Intn(u.width), rand.Intn(u.height), 1)
	}
}

func (u *Universe) Alive(x, y int) bool {
	x += u.width
	x %= u.width
	y += u.height
	y %= u.height
	return u.display.GetPixel(!u.page, x, y) != 0
}

func (u *Universe) Next(x, y int) bool {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && u.Alive(x+i, y+j) {
				alive++
			}
		}
	}
	// game rules:
	//   if 3 alive neighbors: set alive
	//   if 2 alive neighbors: no change
	//   else: set dead
	return alive == 3 || alive == 2 && u.Alive(x, y)
}

func (u *Universe) Step() {
	for x := 0; x < u.width; x++ {
		for y := 0; y < u.height; y++ {
			var color uint8
			if u.Next(x, y) { // get state from active page
				color = 1 // green
			} else {
				color = 0 // black
			}
			u.display.SetPixel(u.page, x, y, color) // set non-active page
		}
	}
}

func (u *Universe) Show() {
	u.display.DisplayPage(u.page)
	u.page = !u.page // swap active page
}

func main() {
	universe := Universe{
		display: Mode4.Display,
		width:   Mode4.Width,
		height:  Mode4.Height,
		page:    true}

	universe.display.Configure()

	// palette index 0 is black and 1 is green
	Mode4.Palette.SetColor(1, color.RGBA{0, 255, 0, 255})

	universe.Populate()
	universe.Show()

	for {
		universe.Step()
		universe.Show()
	}
}

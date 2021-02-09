package main

import (
	"gba/mode5"
	"gba/registers"
	"image/color"
	"math/rand"
)

func randColor() color.RGBA {
	r := uint8(rand.Intn(0x20))
	g := uint8(rand.Intn(0x20))
	b := uint8(rand.Intn(0x20))
	return color.RGBA{R: r, G: g, B: b}
}

type Universe struct {
	display       *Mode5.DisplayBuffer
	width, height int
	page          bool
}

func (u *Universe) Populate() {
	for i := 0; i < (u.width * u.height / 4); i++ {
		u.display.SetPixel(u.page, rand.Intn(u.width), rand.Intn(u.height), randColor())
	}
}

func (u *Universe) Alive(x, y int) bool {
	x += u.width
	x %= u.width
	y += u.height
	y %= u.height
	return u.display.GetPixel(u.page, x, y) != color.RGBA{} // black
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
			// get state from active page and set non-active page
			if u.Next(x, y) {
				u.display.SetPixel(!u.page, x, y, randColor())
			} else {
				u.display.SetPixel(!u.page, x, y, color.RGBA{})
			}
		}
	}
}

func (u *Universe) Show() {
	u.page = !u.page // swap active page
	u.display.DisplayPage(u.page)
}

func main() {
	universe := Universe{
		display: &Mode5.Display,
		width:   60,
		height:  40,
		page:    false}

	// scale top left corner 4x
	scale := uint16(1 << 6)
	Registers.BG2PA.Set(scale)
	Registers.BG2PD.Set(scale)

	universe.display.Configure()
	universe.Populate()

	for {
		universe.Step()
		universe.Show()
	}
}

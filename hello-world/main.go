package main

import (
	"image/color"
	"machine"
	"tinygo.org/x/tinyfont"
)

var display = machine.Display
var colors = []color.RGBA{
	{255, 0, 0, 255},
	{255, 255, 0, 255},
	{0, 255, 0, 255},
	{0, 255, 255, 255},
	{0, 0, 255, 255},
	{255, 0, 255, 255},
	{255, 255, 255, 255},
}

func main() {
	display.Configure()

	tinyfont.WriteLine(&display, &tinyfont.TomThumb, 18, 12, "Hello", colors[0])
	tinyfont.WriteLineColors(&display, &tinyfont.Org01, 12, 28, "World!", colors[1:])
}

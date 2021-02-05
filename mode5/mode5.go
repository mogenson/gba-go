package Mode5

import (
	"gba/registers"

	"image/color"
	"runtime/volatile"
	"unsafe"
)

const Width = 160
const Height = 128

var page0 = (*[Height][Width]volatile.Register16)(unsafe.Pointer(uintptr(0x06000000)))
var page1 = (*[Height][Width]volatile.Register16)(unsafe.Pointer(uintptr(0x0600A000)))

var Display = DisplayBuffer{page0, page1}

type DisplayBuffer struct {
	page0 *[Height][Width]volatile.Register16
	page1 *[Height][Width]volatile.Register16
}

func (d DisplayBuffer) Configure() {
	Registers.DISPCNT.Set(0x0405) // Mode5 + BG2
}

func (d DisplayBuffer) Size() (x, y int16) {
	return Width, Height
}

func (d DisplayBuffer) SetPixel(p bool, x, y int, c color.RGBA) {
	r := uint16(c.R) & 0x1f
	g := uint16(c.G) & 0x1f
	b := uint16(c.B) & 0x1f
	val := b<<10 | g<<5 | r

	if p {
		d.page1[y][x].Set(val)
	} else {
		d.page0[y][x].Set(val)
	}
}

func (d DisplayBuffer) GetPixel(p bool, x, y int) color.RGBA {
	var val uint16

	if p {
		val = d.page1[y][x].Get()
	} else {
		val = d.page0[y][x].Get()
	}

	r := uint8(val) & 0x1f
	g := uint8(val>>5) & 0x1f
	b := uint8(val>>10) & 0x1f

	return color.RGBA{R: r, G: g, B: b}
}

func (d DisplayBuffer) DisplayPage(p bool) {
	mask := uint16(0x0010)
	if p {
		Registers.DISPCNT.SetBits(mask) // page 1
	} else {
		Registers.DISPCNT.ClearBits(mask) // page 0
	}
}

func (d DisplayBuffer) VBlank() bool {
	return Registers.VCOUNT.Get() >= Height // in VBlank period
}

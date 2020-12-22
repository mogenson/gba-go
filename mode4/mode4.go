package Mode4

import (
	"gba/registers"

	"image/color"
	"runtime/volatile"
	"unsafe"
)

const Width = 240
const Height = 160

var pram = (*[256]volatile.Register16)(unsafe.Pointer(uintptr(0x05000000)))
var page0 = (*[Height][Width]volatile.Register8)(unsafe.Pointer(uintptr(0x06000000)))
var page1 = (*[Height][Width]volatile.Register8)(unsafe.Pointer(uintptr(0x0600A000)))

var Display = DisplayBuffer{page0, page1}
var Palette = PaletteBuffer{pram}

type PaletteBuffer struct {
	port *[256]volatile.Register16
}

func (p PaletteBuffer) SetColor(i uint8, c color.RGBA) {
	r := uint16(c.R) & 0x1f
	g := uint16(c.G) & 0x1f
	b := uint16(c.B) & 0x1f
	p.port[i].Set(b<<10 | g<<5 | r)
}

type DisplayBuffer struct {
	page0 *[Height][Width]volatile.Register8
	page1 *[Height][Width]volatile.Register8
}

func (d DisplayBuffer) Configure() {
	Registers.DISPCNT.Set(0x0404) // Mode4 + BG2
}

func (d DisplayBuffer) Size() (x, y int16) {
	return Width, Height
}

func (d DisplayBuffer) SetPixel(p bool, x, y int, c uint8) {
	if p {
		d.page1[y][x].Set(c)
	} else {
		d.page0[y][x].Set(c)
	}
}

func (d DisplayBuffer) GetPixel(p bool, x, y int) uint8 {
	if p {
		return d.page1[y][x].Get()
	} else {
		return d.page0[y][x].Get()
	}
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

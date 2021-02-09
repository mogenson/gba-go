package main

import (
	"image/color"
	"machine"
	"runtime/interrupt"
	"runtime/volatile"
	"unsafe"
)

var display = machine.Display
var DISPSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004)))

func main() {
	display.Configure()
	display.SetPixel(80, 80, color.RGBA{R: 255})
	DISPSTAT.SetBits(1 << 3) // interrupt on vblank
	interrupt.New(machine.IRQ_VBLANK, vblankIrq).Enable()

	for {
		display.SetPixel(100, 80, color.RGBA{G: 255})
	}
}

func vblankIrq(interrupt.Interrupt) {
	display.SetPixel(120, 80, color.RGBA{B: 255})
}

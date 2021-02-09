package main

import (
	"image/color"
	"machine"
	"runtime/interrupt"
	"runtime/volatile"
	"unsafe"
)

var display = machine.Display
var TM0CNT_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000102)))

func main() {
	display.Configure()
	display.SetPixel(80, 80, color.RGBA{R: 255})
	TM0CNT_H.SetBits(1 << 6 | 1 << 7) // start timer, interrupt on overflow
	interrupt.New(machine.IRQ_TIMER0, timerIrq).Enable()

	for {
		display.SetPixel(100, 80, color.RGBA{G: 255})
	}
}

func timerIrq(interrupt.Interrupt) {
	display.SetPixel(120, 80, color.RGBA{B: 255})
}

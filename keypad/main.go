package main

import (
	"image/color"
	"machine"
	"runtime/interrupt"
	"runtime/volatile"
	"unsafe"
)

var KEYPAD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000130)))
var KEYCNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000132)))
var display = machine.Display

func main() {
	display.Configure()
	display.SetPixel(120, 80, color.RGBA{R: 255})
	KEYCNT.SetBits(1 | 1 << 1 | 1 << 14) // generate interrupts for A and B buttons
	interrupt.New(machine.IRQ_KEYPAD, keypadIrq).Enable()

	for {
		//readKeypad()
		display.SetPixel(0,0, color.RGBA{})
	}
}

func readKeypad() {
	if KEYPAD.Get()&1 == 0 {
		// A is pushed when bit 0 is cleared
		display.SetPixel(120, 80, color.RGBA{G: 255})
	} else if KEYPAD.Get()&2 == 0 {
		// B is pushed when bit 1 is cleared
		display.SetPixel(120, 80, color.RGBA{B: 255})
	}
}

func keypadIrq(interrupt.Interrupt) {
	readKeypad()
}

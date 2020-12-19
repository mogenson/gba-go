package Registers

import (
	"runtime/volatile"
	"unsafe"
)

// display
var DISPCNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000000)))
var DISPSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000004)))
var VCOUNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000006)))

// background
var BG0CNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000008)))
var BG1CNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400000A)))
var BG2CNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400000C)))
var BG3CNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400000E)))
var BG0HOFS = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000010)))
var BG0VOFS = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000012)))
var BG1HOFS = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000014)))
var BG1VOFS = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000016)))
var BG2HOFS = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000018)))
var BG2VOFS = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400001A)))
var BG3HOFS = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400001C)))
var BG3VOFS = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400001E)))
var BG2PA = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000020)))
var BG2PB = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000022)))
var BG2PC = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000024)))
var BG2PD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000026)))
var BG2X = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000028)))
var BG2Y = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400002C)))
var BG3PA = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000030)))
var BG3PB = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000032)))
var BG3PC = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000034)))
var BG3PD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000036)))
var B32X = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000038)))
var B32Y = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400003C)))

// window
var WIN0H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000040)))
var WIN1H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000042)))
var WIN0V = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000044)))
var WIN1V = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000046)))
var WININ = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000048)))
var WINOUT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400004A)))

// effects
var MOSAIC = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400004C)))
var BLDMOD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000050)))
var COLEV = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000052)))
var COLEY = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000054)))

// sound
var SOUND1CNT_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000060)))
var SOUND1CNT_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000062)))
var SOUND1CNT_X = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000064)))
var SOUND2CNT_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000068)))
var SOUND2CNT_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400006C)))
var SOUND3CNT_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000070)))
var SOUND3CNT_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000072)))
var SOUND3CNT_X = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000074)))
var SOUND4CNT_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000078)))
var SOUND4CNT_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400007C)))
var SOUNDCNT_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000080)))
var SOUNDCNT_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000082)))
var SOUNDCNT_X = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000084)))
var SOUNDBIAS = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000088)))
var WAVE_RAM0_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000090)))
var WAVE_RAM0_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000092)))
var WAVE_RAM1_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000094)))
var WAVE_RAM1_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000096)))
var WAVE_RAM2_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000098)))
var WAVE_RAM2_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400009A)))
var WAVE_RAM3_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400009C)))
var WAVE_RAM3_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400009E)))
var FIFO_A_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000A0)))
var FIFO_A_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000A2)))
var FIFO_B_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000A4)))
var FIFO_B_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000A6)))

// dma
var DMA0SAD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000B0)))
var DMA0DAD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000B4)))
var DMA0CNT_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000B8)))
var DMA0CNT_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000BA)))
var DMA1SAD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000BC)))
var DMA1DAD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000C0)))
var DMA1CNT_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000C4)))
var DMA1CNT_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000C6)))
var DMA2SAD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000C8)))
var DMA2DAD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000CC)))
var DMA2CNT_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000D0)))
var DMA2CNT_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000D2)))
var DMA3SAD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000D4)))
var DMA3DAD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000D8)))
var DMA3CNT_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000DC)))
var DMA3CNT_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x040000DE)))

// timers
var TM0D = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000100)))
var TM0CNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000102)))
var TM1D = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000104)))
var TM1CNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000106)))
var TM2D = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000108)))
var TM2CNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400010A)))
var TM3D = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400010C)))
var TM3CNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400010E)))

// serial communication
var SCD0 = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000120)))
var SCD1 = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000122)))
var SCD2 = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000124)))
var SCD3 = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000126)))
var SCCNT_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000128)))
var SCCNT_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x0400012A)))

// keypad
var KEY = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000130)))
var P1CNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000132)))

// gpio
var RCNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000134)))

// joybus
var JOYCNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000140)))
var JOYRE_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000150)))
var JOYRE_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000152)))
var JOYTR_L = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000154)))
var JOYTR_H = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000156)))
var JOYSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000158)))

// interrupts
var IE = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000200)))
var IF = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000202)))
var WSCNT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000204)))
var IME = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000208)))

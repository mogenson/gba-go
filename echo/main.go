package main

import (
	"gba/uart"
)

var uart = UART.UART0

func main() {
	uart.Configure(UART.UARTConfig{BaudRate: 115200})
	uart.Write([]byte("Hello World\n"))

	for {
		if uart.Buffered() > 0 {
			data, _ := uart.ReadByte()
			uart.WriteByte(data)
		}
	}
}

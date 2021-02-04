package UART

import (
	"errors"
	"gba/registers"
	"machine"
	"runtime/interrupt"
)

type UART struct {
	Buffer *machine.RingBuffer
}

type UARTConfig struct {
	BaudRate uint32
}

var UART0 = UART{Buffer: machine.NewRingBuffer()}

func (uart UART) Configure(config UARTConfig) {
	sccnt := uint16(0)

	switch config.BaudRate {
	case 9600:
		sccnt |= 0
	case 38400:
		sccnt |= 1
	case 57600:
		sccnt |= 2
	case 115200:
		sccnt |= 3
	}

	sccnt |= 1 << 7  // 8-bit data length
	sccnt |= 1 << 10 // tx enable
	sccnt |= 1 << 11 // rx enable
	sccnt |= 3 << 12 // uart mode
	sccnt |= 1 << 14 // irq enable

	Registers.RCNT.Set(0)        // io mode disabled
	Registers.SCCNT_L.Set(sccnt) // 8n1, no flow control, no parity

	interrupt.New(machine.IRQ_COM, UART0.handleInterrupt).Enable()
}

func (uart *UART) handleInterrupt(interrupt.Interrupt) {
	if (Registers.SCCNT_L.Get() & (1 << 5)) == 0 { // rx empty flag is clear
		data := byte(Registers.SCCNT_H.Get())
		uart.Receive(data)
	}
}

func (uart UART) WriteByte(c byte) {
	for Registers.SCCNT_L.HasBits(1 << 4) { // tx full flag is set
		continue
	}

	Registers.SCCNT_H.Set(uint16(c))
}

// the following is copied from tinygo/src/machine/uart.go

var errUARTBufferEmpty = errors.New("UART buffer empty")

// Read from the RX buffer.
func (uart UART) Read(data []byte) (n int, err error) {
	// check if RX buffer is empty
	size := uart.Buffered()
	if size == 0 {
		return 0, nil
	}

	// Make sure we do not read more from buffer than the data slice can hold.
	if len(data) < size {
		size = len(data)
	}

	// only read number of bytes used from buffer
	for i := 0; i < size; i++ {
		v, _ := uart.ReadByte()
		data[i] = v
	}

	return size, nil
}

// Write data to the UART.
func (uart UART) Write(data []byte) (n int, err error) {
	for _, v := range data {
		uart.WriteByte(v)
	}
	return len(data), nil
}

// ReadByte reads a single byte from the RX buffer.
// If there is no data in the buffer, returns an error.
func (uart UART) ReadByte() (byte, error) {
	// check if RX buffer is empty
	buf, ok := uart.Buffer.Get()
	if !ok {
		return 0, errUARTBufferEmpty
	}
	return buf, nil
}

// Buffered returns the number of bytes currently stored in the RX buffer.
func (uart UART) Buffered() int {
	return int(uart.Buffer.Used())
}

// Receive handles adding data to the UART's data buffer.
// Usually called by the IRQ handler for a machine.
func (uart UART) Receive(data byte) {
	uart.Buffer.Put(data)
}

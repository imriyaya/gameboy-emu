package cpu

import "github.com/imriyaya/gameboy-emu/internal/memory"

type CPU struct {
	registers *Registers
	memory    *memory.Memory
}

func (cpu *CPU) execute(opcode uint8) {
	switch opcode {
	// NOP
	case 0b00000000:
		break

	// STOP n8

	// LD r8 r8
	case 0b01000000:
		cpu.LD8(&cpu.registers.B, cpu.registers.B)
	case 0b01000001:
		cpu.LD8(&cpu.registers.B, cpu.registers.C)
	case 0b01000010:
		cpu.LD8(&cpu.registers.B, cpu.registers.D)
	case 0b01000011:
		cpu.LD8(&cpu.registers.B, cpu.registers.E)
	case 0b01000100:
		cpu.LD8(&cpu.registers.B, cpu.registers.H)
	case 0b01000101:
		cpu.LD8(&cpu.registers.B, cpu.registers.L)
	case 0b01000110:
		cpu.LD8(&cpu.registers.B, cpu.memory.GetValue(cpu.registers.GetHL()))
	case 0b01000111:
		cpu.LD8(&cpu.registers.B, cpu.registers.A)
	case 0b01001000:
		cpu.LD8(&cpu.registers.C, cpu.registers.B)
	case 0b01001001:
		cpu.LD8(&cpu.registers.C, cpu.registers.C)
	case 0b01001010:
		cpu.LD8(&cpu.registers.C, cpu.registers.D)
	case 0b01001011:
		cpu.LD8(&cpu.registers.C, cpu.registers.E)
	case 0b01001100:
		cpu.LD8(&cpu.registers.C, cpu.registers.H)
	case 0b01001101:
		cpu.LD8(&cpu.registers.C, cpu.registers.L)
	case 0b01001110:
		cpu.LD8(&cpu.registers.C, cpu.memory.GetValue(cpu.registers.GetHL()))
	case 0b01001111:
		cpu.LD8(&cpu.registers.C, cpu.registers.A)
	case 0b01010000:
		cpu.LD8(&cpu.registers.D, cpu.registers.B)
	case 0b01010001:
		cpu.LD8(&cpu.registers.D, cpu.registers.C)
	case 0b01010010:
		cpu.LD8(&cpu.registers.D, cpu.registers.D)
	case 0b01010011:
		cpu.LD8(&cpu.registers.D, cpu.registers.E)
	case 0b01010100:
		cpu.LD8(&cpu.registers.D, cpu.registers.H)
	case 0b01010101:
		cpu.LD8(&cpu.registers.D, cpu.registers.L)
	case 0b01010110:
		cpu.LD8(&cpu.registers.D, cpu.memory.GetValue(cpu.registers.GetHL()))
	case 0b01010111:
		cpu.LD8(&cpu.registers.D, cpu.registers.A)
	case 0b01011000:
		cpu.LD8(&cpu.registers.E, cpu.registers.B)
	case 0b01011001:
		cpu.LD8(&cpu.registers.E, cpu.registers.C)
	case 0b01011010:
		cpu.LD8(&cpu.registers.E, cpu.registers.D)
	case 0b01011011:
		cpu.LD8(&cpu.registers.E, cpu.registers.E)
	case 0b01011100:
		cpu.LD8(&cpu.registers.E, cpu.registers.H)
	case 0b01011101:
		cpu.LD8(&cpu.registers.E, cpu.registers.L)
	case 0b01011110:
		cpu.LD8(&cpu.registers.E, cpu.memory.GetValue(cpu.registers.GetHL()))
	case 0b01011111:
		cpu.LD8(&cpu.registers.E, cpu.registers.A)
	case 0b01100000:
		cpu.LD8(&cpu.registers.H, cpu.registers.B)
	case 0b01100001:
		cpu.LD8(&cpu.registers.H, cpu.registers.C)
	case 0b01100010:
		cpu.LD8(&cpu.registers.H, cpu.registers.D)
	case 0b01100011:
		cpu.LD8(&cpu.registers.H, cpu.registers.E)
	case 0b01100100:
		cpu.LD8(&cpu.registers.H, cpu.registers.H)
	case 0b01100101:
		cpu.LD8(&cpu.registers.H, cpu.registers.L)
	case 0b01100110:
		cpu.LD8(&cpu.registers.H, cpu.memory.GetValue(cpu.registers.GetHL()))
	case 0b01100111:
		cpu.LD8(&cpu.registers.H, cpu.registers.A)
	case 0b01101000:
		cpu.LD8(&cpu.registers.L, cpu.registers.B)
	case 0b01101001:
		cpu.LD8(&cpu.registers.L, cpu.registers.C)
	case 0b01101010:
		cpu.LD8(&cpu.registers.L, cpu.registers.D)
	case 0b01101011:
		cpu.LD8(&cpu.registers.L, cpu.registers.E)
	case 0b01101100:
		cpu.LD8(&cpu.registers.L, cpu.registers.H)
	case 0b01101101:
		cpu.LD8(&cpu.registers.L, cpu.registers.L)
	case 0b01101110:
		cpu.LD8(&cpu.registers.L, cpu.memory.GetValue(cpu.registers.GetHL()))
	case 0b01101111:
		cpu.LD8(&cpu.registers.L, cpu.registers.A)
	case 0b01110000:
		cpu.LD8(cpu.memory.GetPointer(cpu.registers.GetHL()), cpu.registers.B)
	case 0b01110001:
		cpu.LD8(cpu.memory.GetPointer(cpu.registers.GetHL()), cpu.registers.C)
	case 0b01110010:
		cpu.LD8(cpu.memory.GetPointer(cpu.registers.GetHL()), cpu.registers.D)
	case 0b01110011:
		cpu.LD8(cpu.memory.GetPointer(cpu.registers.GetHL()), cpu.registers.E)
	case 0b01110100:
		cpu.LD8(cpu.memory.GetPointer(cpu.registers.GetHL()), cpu.registers.H)
	case 0b01110101:
		cpu.LD8(cpu.memory.GetPointer(cpu.registers.GetHL()), cpu.registers.L)
	case 0b01110110:
		cpu.LD8(cpu.memory.GetPointer(cpu.registers.GetHL()), cpu.memory.GetValue(cpu.registers.GetHL()))
	case 0b01110111:
		cpu.LD8(cpu.memory.GetPointer(cpu.registers.GetHL()), cpu.registers.A)
	case 0b01111000:
		cpu.LD8(&cpu.registers.A, cpu.registers.B)
	case 0b01111001:
		cpu.LD8(&cpu.registers.A, cpu.registers.C)
	case 0b01111010:
		cpu.LD8(&cpu.registers.A, cpu.registers.D)
	case 0b01111011:
		cpu.LD8(&cpu.registers.A, cpu.registers.E)
	case 0b01111100:
		cpu.LD8(&cpu.registers.A, cpu.registers.H)
	case 0b01111101:
		cpu.LD8(&cpu.registers.A, cpu.registers.L)
	case 0b01111110:
		cpu.LD8(&cpu.registers.A, cpu.memory.GetValue(cpu.registers.GetHL()))
	case 0b01111111:
		cpu.LD8(&cpu.registers.A, cpu.registers.A)

	}
}

func (cpu *CPU) ADD8(destination *uint8, source uint8) {
	total := uint(source) + uint(*destination)
	*destination = uint8(total)

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry((source&0xF)+(*destination&0xF) > 0xF)
	cpu.registers.SetCarry(total > 0xFF)
}

func (cpu *CPU) ADC8(destination *uint8, source uint8) {
	carry := uint(0)
	if cpu.registers.GetCarry() {
		carry = uint(1)
	}

	total := uint(source) + uint(*destination) + carry
	*destination = uint8(total)

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry((source&0xF)+(*destination&0xF) > 0xF)
	cpu.registers.SetCarry(total > 0xFF)
}

func (cpu *CPU) SUB8(destination *uint8, source uint8) {
	total := uint(*destination) - uint(source)
	*destination = uint8(total)

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(true)
	cpu.registers.SetHalfCarry(int16(*destination&0x0f)-int16(source&0xF) < 0)
	cpu.registers.SetCarry(total < 0)
}

func (cpu *CPU) SBC8(destination *uint8, source uint8) {
	carry := uint(0)
	if cpu.registers.GetCarry() {
		carry = uint(1)
	}
	total := uint(*destination) - uint(source) - carry
	*destination = uint8(total)

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(true)
	cpu.registers.SetHalfCarry(int16(*destination&0x0f)-int16(source&0xF)-int16(carry) < 0)
	cpu.registers.SetCarry(total < 0)
}

func (cpu *CPU) AND8(destination *uint8, source uint8) {
	*destination = *destination & source

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry(true)
	cpu.registers.SetCarry(false)
}

func (cpu *CPU) CP8(destination *uint8, source uint8) {
	result := *destination - source

	cpu.registers.SetZero(result == 0)
	cpu.registers.SetSubtract(true)
	cpu.registers.SetHalfCarry((source & 0x0f) > (*destination & 0x0f))
	cpu.registers.SetCarry(source > *destination)
}

func (cpu *CPU) OR8(destination *uint8, source uint8) {
	*destination = *destination | source

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry(false)
	cpu.registers.SetCarry(false)
}

func (cpu *CPU) XOR8(destination *uint8, source uint8) {
	*destination = *destination ^ source

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry(false)
	cpu.registers.SetCarry(false)
}

func (cpu *CPU) INC8(destination *uint8) {
	*destination = *destination + 1

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry((*destination&0xF)+(1&0xF) > 0xF)
}

func (cpu *CPU) DEC8(destination *uint8) {
	*destination = *destination - 1

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(true)
	cpu.registers.SetHalfCarry(*destination&0x0F == 0)
}

func (cpu *CPU) ADD16(destination *uint16, source uint16) {
	total := uint32(source) + uint32(*destination)
	*destination = uint16(total)

	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry(uint32(*destination&0xFFF) > (total & 0xFFF))
	cpu.registers.SetCarry(total > 0xFFFF)
}

func (cpu *CPU) INC16(destination *uint16) {
	*destination = *destination + 1
}

func (cpu *CPU) DEC16(destination *uint16) {
	*destination = *destination - 1
}

func (cpu *CPU) BIT(source uint8, bitIndex uint8) {
	cpu.registers.SetZero((source>>bitIndex)&1 == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry(true)
}

func (cpu *CPU) SET(destination *uint8, bitIndex uint8) {
	*destination = *destination | (1 << bitIndex)
}

func (cpu *CPU) RES(destination *uint8, bitIndex uint8) {
	*destination = *destination & ^(1 << bitIndex)
}

func (cpu *CPU) SWAP(destination *uint8) {
	swapped := *destination<<4&240 | *destination>>4
	*destination = swapped

	cpu.registers.SetZero(swapped == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry(false)
	cpu.registers.SetCarry(false)
}

func (cpu *CPU) RL(destination *uint8) {
	newCarry := *destination >> 7
	oldCarry := uint8(0)
	if cpu.registers.GetCarry() {
		oldCarry = uint8(1)
	}
	*destination = (*destination<<1)&0xFF | oldCarry

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry(false)
	cpu.registers.SetCarry(newCarry == 1)
}

func (cpu *CPU) RLC(destination *uint8) {
	carry := *destination >> 7
	*destination = (*destination<<1)&0xFF | carry

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry(false)
	cpu.registers.SetCarry(carry == 1)
}

func (cpu *CPU) RR(destination *uint8) {
	newCarry := *destination & 1
	oldCarry := uint8(0)
	if cpu.registers.GetCarry() {
		oldCarry = uint8(1)
	}
	*destination = (*destination >> 1) | (oldCarry << 7)

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry(false)
	cpu.registers.SetCarry(newCarry == 1)
}

func (cpu *CPU) RRC(destination *uint8) {
	carry := *destination & 1
	*destination = (*destination >> 1) | (carry << 7)

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry(false)
	cpu.registers.SetCarry(carry == 1)
}

func (cpu *CPU) SLA(destination *uint8) {
	carry := *destination >> 7
	*destination = *destination << 1 & 0xFF

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry(false)
	cpu.registers.SetCarry(carry == 1)
}

func (cpu *CPU) SRA(destination *uint8) {
	*destination = (*destination & 128) | (*destination >> 1)

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry(false)
	cpu.registers.SetCarry(*destination&1 == 1)
}

func (cpu *CPU) SRL(destination *uint8) {
	carry := *destination & 1
	*destination = *destination >> 1

	cpu.registers.SetZero(*destination == 0)
	cpu.registers.SetSubtract(false)
	cpu.registers.SetHalfCarry(false)
	cpu.registers.SetCarry(carry == 1)
}

func (cpu *CPU) LD8(destination *uint8, source uint8) {
	*destination = source
}

func (cpu *CPU) LD16(destination *uint16, source uint16) {
	*destination = source
}

func (cpu *CPU) LDH8(destination *uint8, source uint8) {

}

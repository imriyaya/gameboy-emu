package pkg

type CPU struct {
	registers *Registers
}

func (gameboy *Gameboy) execute(opcode uint8) {
	// The following instruction codes from https://gbdev.io/pandocs/CPU_Instruction_Set.html

	switch opcode {
	// NOP
	case 0b00000000:
		break

	// LD r16 imm16
	case 0b00000001:
		gameboy.cpu.registers.SetBC(gameboy.POP_PC16())
	case 0b00010001:
		gameboy.cpu.registers.SetDE(gameboy.POP_PC16())
	case 0b00100001:
		gameboy.cpu.registers.SetHL(gameboy.POP_PC16())
	case 0b00110001:
		gameboy.cpu.registers.SetSP(gameboy.POP_PC16())

	// LD [r16mem] a
	case 0b00000010:
		gameboy.LD8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetBC()), gameboy.cpu.registers.A)
	case 0b00010010:
		gameboy.LD8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetDE()), gameboy.cpu.registers.A)
	case 0b00100010:
		gameboy.LD8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()), gameboy.cpu.registers.A)
	case 0b00110010:
		gameboy.LD8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetSP()), gameboy.cpu.registers.A)

	// LD a [r16mem]
	case 0b00001010:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.cpu.registers.GetBC()))
	case 0b00011010:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.cpu.registers.GetDE()))
	case 0b00101010:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b00111010:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.cpu.registers.GetSP()))

	// LD imm16 sp
	case 0b00001000:
		address := gameboy.cpu.registers.GetSP()
		gameboy.memory.SetValue(address, gameboy.cpu.registers.P)
		gameboy.memory.SetValue(address+1, gameboy.cpu.registers.S)

	// INC r16
	case 0b00000011:
		r16 := gameboy.cpu.registers.GetBC()
		gameboy.INC16(&r16)
		gameboy.cpu.registers.SetBC(r16)
	case 0b00010011:
		r16 := gameboy.cpu.registers.GetDE()
		gameboy.INC16(&r16)
		gameboy.cpu.registers.SetDE(r16)
	case 0b00100011:
		r16 := gameboy.cpu.registers.GetHL()
		gameboy.INC16(&r16)
		gameboy.cpu.registers.SetHL(r16)
	case 0b00110011:
		r16 := gameboy.cpu.registers.GetSP()
		gameboy.INC16(&r16)
		gameboy.cpu.registers.SetSP(r16)

	// DEC r16
	case 0b00001011:
		r16 := gameboy.cpu.registers.GetBC()
		gameboy.DEC16(&r16)
		gameboy.cpu.registers.SetBC(r16)
	case 0b00011011:
		r16 := gameboy.cpu.registers.GetDE()
		gameboy.DEC16(&r16)
		gameboy.cpu.registers.SetDE(r16)
	case 0b00101011:
		r16 := gameboy.cpu.registers.GetHL()
		gameboy.DEC16(&r16)
		gameboy.cpu.registers.SetHL(r16)
	case 0b00111011:
		r16 := gameboy.cpu.registers.GetSP()
		gameboy.DEC16(&r16)
		gameboy.cpu.registers.SetSP(r16)

	// ADD hl r16
	case 0b00001001:
		hl := gameboy.cpu.registers.GetHL()
		gameboy.ADD16(&hl, gameboy.cpu.registers.GetBC())
		gameboy.cpu.registers.SetHL(hl)
	case 0b00011001:
		hl := gameboy.cpu.registers.GetHL()
		gameboy.ADD16(&hl, gameboy.cpu.registers.GetDE())
		gameboy.cpu.registers.SetHL(hl)
	case 0b00101001:
		hl := gameboy.cpu.registers.GetHL()
		gameboy.ADD16(&hl, gameboy.cpu.registers.GetHL())
		gameboy.cpu.registers.SetHL(hl)
	case 0b00111001:
		hl := gameboy.cpu.registers.GetHL()
		gameboy.ADD16(&hl, gameboy.cpu.registers.GetAF())
		gameboy.cpu.registers.SetHL(hl)

	// INC r8
	case 0b00000100:
		gameboy.INC8(&gameboy.cpu.registers.B)
	case 0b00001100:
		gameboy.INC8(&gameboy.cpu.registers.C)
	case 0b00010100:
		gameboy.INC8(&gameboy.cpu.registers.D)
	case 0b00011100:
		gameboy.INC8(&gameboy.cpu.registers.E)
	case 0b00100100:
		gameboy.INC8(&gameboy.cpu.registers.H)
	case 0b00101100:
		gameboy.INC8(&gameboy.cpu.registers.L)
	case 0b00110100:
		gameboy.INC8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
	case 0b00111100:
		gameboy.INC8(&gameboy.cpu.registers.A)

	// DEC r8
	case 0b00000101:
		gameboy.DEC8(&gameboy.cpu.registers.B)
	case 0b00001101:
		gameboy.DEC8(&gameboy.cpu.registers.C)
	case 0b00010101:
		gameboy.DEC8(&gameboy.cpu.registers.D)
	case 0b00011101:
		gameboy.DEC8(&gameboy.cpu.registers.E)
	case 0b00100101:
		gameboy.DEC8(&gameboy.cpu.registers.H)
	case 0b00101101:
		gameboy.DEC8(&gameboy.cpu.registers.L)
	case 0b00110101:
		gameboy.DEC8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
	case 0b00111101:
		gameboy.DEC8(&gameboy.cpu.registers.A)

	// LD r8 imm8
	case 0b00000110:
		gameboy.LD8(&gameboy.cpu.registers.B, gameboy.POP_PC8())
	case 0b00001110:
		gameboy.LD8(&gameboy.cpu.registers.C, gameboy.POP_PC8())
	case 0b00010110:
		gameboy.LD8(&gameboy.cpu.registers.D, gameboy.POP_PC8())
	case 0b00011110:
		gameboy.LD8(&gameboy.cpu.registers.E, gameboy.POP_PC8())
	case 0b00100110:
		gameboy.LD8(&gameboy.cpu.registers.H, gameboy.POP_PC8())
	case 0b00101110:
		gameboy.LD8(&gameboy.cpu.registers.L, gameboy.POP_PC8())
	case 0b00110110:
		gameboy.LD8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()), gameboy.POP_PC8())
	case 0b00111110:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.POP_PC8())

	// RLCA
	case 0b00000111:
		gameboy.RLC(&gameboy.cpu.registers.A)

	// RRCA
	case 0b00001111:
		gameboy.RRC(&gameboy.cpu.registers.A)

	// RLA
	case 0b00010111:
		gameboy.RL(&gameboy.cpu.registers.A)

	// RRA
	case 0b00011111:
		gameboy.RR(&gameboy.cpu.registers.A)

	// DAA
	case 0b00100111:
		gameboy.DAA()

	// CPL
	case 0b00101111:
		gameboy.CPL()

	// SCF
	case 0b00110111:
		gameboy.SCF()

	// CCF
	case 0b00111111:
		gameboy.CCF()

	// JR imm8
	case 0b00011000:
		// TODO
		break

	// JR cond imm8
	case 0b00100000:
		// TODO
		break

	case 0b00101000:
		// TODO
		break

	case 0b00110000:
		// TODO
		break

	case 0b00111000:
		// TODO
		break

	// STOP
	case 0b00010000:
		// TODO
		break

	// LD r8 r8
	case 0b01000000:
		gameboy.LD8(&gameboy.cpu.registers.B, gameboy.cpu.registers.B)
	case 0b01000001:
		gameboy.LD8(&gameboy.cpu.registers.B, gameboy.cpu.registers.C)
	case 0b01000010:
		gameboy.LD8(&gameboy.cpu.registers.B, gameboy.cpu.registers.D)
	case 0b01000011:
		gameboy.LD8(&gameboy.cpu.registers.B, gameboy.cpu.registers.E)
	case 0b01000100:
		gameboy.LD8(&gameboy.cpu.registers.B, gameboy.cpu.registers.H)
	case 0b01000101:
		gameboy.LD8(&gameboy.cpu.registers.B, gameboy.cpu.registers.L)
	case 0b01000110:
		gameboy.LD8(&gameboy.cpu.registers.B, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b01000111:
		gameboy.LD8(&gameboy.cpu.registers.B, gameboy.cpu.registers.A)
	case 0b01001000:
		gameboy.LD8(&gameboy.cpu.registers.C, gameboy.cpu.registers.B)
	case 0b01001001:
		gameboy.LD8(&gameboy.cpu.registers.C, gameboy.cpu.registers.C)
	case 0b01001010:
		gameboy.LD8(&gameboy.cpu.registers.C, gameboy.cpu.registers.D)
	case 0b01001011:
		gameboy.LD8(&gameboy.cpu.registers.C, gameboy.cpu.registers.E)
	case 0b01001100:
		gameboy.LD8(&gameboy.cpu.registers.C, gameboy.cpu.registers.H)
	case 0b01001101:
		gameboy.LD8(&gameboy.cpu.registers.C, gameboy.cpu.registers.L)
	case 0b01001110:
		gameboy.LD8(&gameboy.cpu.registers.C, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b01001111:
		gameboy.LD8(&gameboy.cpu.registers.C, gameboy.cpu.registers.A)
	case 0b01010000:
		gameboy.LD8(&gameboy.cpu.registers.D, gameboy.cpu.registers.B)
	case 0b01010001:
		gameboy.LD8(&gameboy.cpu.registers.D, gameboy.cpu.registers.C)
	case 0b01010010:
		gameboy.LD8(&gameboy.cpu.registers.D, gameboy.cpu.registers.D)
	case 0b01010011:
		gameboy.LD8(&gameboy.cpu.registers.D, gameboy.cpu.registers.E)
	case 0b01010100:
		gameboy.LD8(&gameboy.cpu.registers.D, gameboy.cpu.registers.H)
	case 0b01010101:
		gameboy.LD8(&gameboy.cpu.registers.D, gameboy.cpu.registers.L)
	case 0b01010110:
		gameboy.LD8(&gameboy.cpu.registers.D, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b01010111:
		gameboy.LD8(&gameboy.cpu.registers.D, gameboy.cpu.registers.A)
	case 0b01011000:
		gameboy.LD8(&gameboy.cpu.registers.E, gameboy.cpu.registers.B)
	case 0b01011001:
		gameboy.LD8(&gameboy.cpu.registers.E, gameboy.cpu.registers.C)
	case 0b01011010:
		gameboy.LD8(&gameboy.cpu.registers.E, gameboy.cpu.registers.D)
	case 0b01011011:
		gameboy.LD8(&gameboy.cpu.registers.E, gameboy.cpu.registers.E)
	case 0b01011100:
		gameboy.LD8(&gameboy.cpu.registers.E, gameboy.cpu.registers.H)
	case 0b01011101:
		gameboy.LD8(&gameboy.cpu.registers.E, gameboy.cpu.registers.L)
	case 0b01011110:
		gameboy.LD8(&gameboy.cpu.registers.E, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b01011111:
		gameboy.LD8(&gameboy.cpu.registers.E, gameboy.cpu.registers.A)
	case 0b01100000:
		gameboy.LD8(&gameboy.cpu.registers.H, gameboy.cpu.registers.B)
	case 0b01100001:
		gameboy.LD8(&gameboy.cpu.registers.H, gameboy.cpu.registers.C)
	case 0b01100010:
		gameboy.LD8(&gameboy.cpu.registers.H, gameboy.cpu.registers.D)
	case 0b01100011:
		gameboy.LD8(&gameboy.cpu.registers.H, gameboy.cpu.registers.E)
	case 0b01100100:
		gameboy.LD8(&gameboy.cpu.registers.H, gameboy.cpu.registers.H)
	case 0b01100101:
		gameboy.LD8(&gameboy.cpu.registers.H, gameboy.cpu.registers.L)
	case 0b01100110:
		gameboy.LD8(&gameboy.cpu.registers.H, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b01100111:
		gameboy.LD8(&gameboy.cpu.registers.H, gameboy.cpu.registers.A)
	case 0b01101000:
		gameboy.LD8(&gameboy.cpu.registers.L, gameboy.cpu.registers.B)
	case 0b01101001:
		gameboy.LD8(&gameboy.cpu.registers.L, gameboy.cpu.registers.C)
	case 0b01101010:
		gameboy.LD8(&gameboy.cpu.registers.L, gameboy.cpu.registers.D)
	case 0b01101011:
		gameboy.LD8(&gameboy.cpu.registers.L, gameboy.cpu.registers.E)
	case 0b01101100:
		gameboy.LD8(&gameboy.cpu.registers.L, gameboy.cpu.registers.H)
	case 0b01101101:
		gameboy.LD8(&gameboy.cpu.registers.L, gameboy.cpu.registers.L)
	case 0b01101110:
		gameboy.LD8(&gameboy.cpu.registers.L, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b01101111:
		gameboy.LD8(&gameboy.cpu.registers.L, gameboy.cpu.registers.A)
	case 0b01110000:
		gameboy.LD8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()), gameboy.cpu.registers.B)
	case 0b01110001:
		gameboy.LD8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()), gameboy.cpu.registers.C)
	case 0b01110010:
		gameboy.LD8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()), gameboy.cpu.registers.D)
	case 0b01110011:
		gameboy.LD8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()), gameboy.cpu.registers.E)
	case 0b01110100:
		gameboy.LD8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()), gameboy.cpu.registers.H)
	case 0b01110101:
		gameboy.LD8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()), gameboy.cpu.registers.L)
	// This is exception. instead HALT
	// case 0b01110110:
	// gameboy.LD8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()), gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b01110111:
		gameboy.LD8(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()), gameboy.cpu.registers.A)
	case 0b01111000:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.B)
	case 0b01111001:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.C)
	case 0b01111010:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.D)
	case 0b01111011:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.E)
	case 0b01111100:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.H)
	case 0b01111101:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.L)
	case 0b01111110:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b01111111:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.A)

	// HALT
	case 0b01110110:
		// TODO
		break

	// ADD a r8
	case 0b10000000:
		gameboy.ADD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.B)
	case 0b10000001:
		gameboy.ADD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.C)
	case 0b10000010:
		gameboy.ADD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.D)
	case 0b10000011:
		gameboy.ADD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.E)
	case 0b10000100:
		gameboy.ADD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.H)
	case 0b10000101:
		gameboy.ADD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.L)
	case 0b10000110:
		gameboy.ADD8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b10000111:
		gameboy.ADD8(&gameboy.cpu.registers.A, gameboy.cpu.registers.A)

	// ADC a r8
	case 0b10001000:
		gameboy.ADC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.B)
	case 0b10001001:
		gameboy.ADC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.C)
	case 0b10001010:
		gameboy.ADC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.D)
	case 0b10001011:
		gameboy.ADC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.E)
	case 0b10001100:
		gameboy.ADC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.H)
	case 0b10001101:
		gameboy.ADC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.L)
	case 0b10001110:
		gameboy.ADC8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b10001111:
		gameboy.ADC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.A)

	// SUB a r8
	case 0b10010000:
		gameboy.SUB8(&gameboy.cpu.registers.A, gameboy.cpu.registers.B)
	case 0b10010001:
		gameboy.SUB8(&gameboy.cpu.registers.A, gameboy.cpu.registers.C)
	case 0b10010010:
		gameboy.SUB8(&gameboy.cpu.registers.A, gameboy.cpu.registers.D)
	case 0b10010011:
		gameboy.SUB8(&gameboy.cpu.registers.A, gameboy.cpu.registers.E)
	case 0b10010100:
		gameboy.SUB8(&gameboy.cpu.registers.A, gameboy.cpu.registers.H)
	case 0b10010101:
		gameboy.SUB8(&gameboy.cpu.registers.A, gameboy.cpu.registers.L)
	case 0b10010110:
		gameboy.SUB8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b10010111:
		gameboy.SUB8(&gameboy.cpu.registers.A, gameboy.cpu.registers.A)

	// SBC a r8
	case 0b10011000:
		gameboy.SBC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.B)
	case 0b10011001:
		gameboy.SBC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.C)
	case 0b10011010:
		gameboy.SBC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.D)
	case 0b10011011:
		gameboy.SBC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.E)
	case 0b10011100:
		gameboy.SBC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.H)
	case 0b10011101:
		gameboy.SBC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.L)
	case 0b10011110:
		gameboy.SBC8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b10011111:
		gameboy.SBC8(&gameboy.cpu.registers.A, gameboy.cpu.registers.A)

	// AND a r8
	case 0b10100000:
		gameboy.AND8(&gameboy.cpu.registers.A, gameboy.cpu.registers.B)
	case 0b10100001:
		gameboy.AND8(&gameboy.cpu.registers.A, gameboy.cpu.registers.C)
	case 0b10100010:
		gameboy.AND8(&gameboy.cpu.registers.A, gameboy.cpu.registers.D)
	case 0b10100011:
		gameboy.AND8(&gameboy.cpu.registers.A, gameboy.cpu.registers.E)
	case 0b10100100:
		gameboy.AND8(&gameboy.cpu.registers.A, gameboy.cpu.registers.H)
	case 0b10100101:
		gameboy.AND8(&gameboy.cpu.registers.A, gameboy.cpu.registers.L)
	case 0b10100110:
		gameboy.AND8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b10100111:
		gameboy.AND8(&gameboy.cpu.registers.A, gameboy.cpu.registers.A)

	// XOR a r8
	case 0b10101000:
		gameboy.XOR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.B)
	case 0b10101001:
		gameboy.XOR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.C)
	case 0b10101010:
		gameboy.XOR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.D)
	case 0b10101011:
		gameboy.XOR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.E)
	case 0b10101100:
		gameboy.XOR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.H)
	case 0b10101101:
		gameboy.XOR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.L)
	case 0b10101110:
		gameboy.XOR8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b10101111:
		gameboy.XOR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.A)

	// OR a r8
	case 0b10110000:
		gameboy.OR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.B)
	case 0b10110001:
		gameboy.OR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.C)
	case 0b10110010:
		gameboy.OR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.D)
	case 0b10110011:
		gameboy.OR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.E)
	case 0b10110100:
		gameboy.OR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.H)
	case 0b10110101:
		gameboy.OR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.L)
	case 0b10110110:
		gameboy.OR8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b10110111:
		gameboy.OR8(&gameboy.cpu.registers.A, gameboy.cpu.registers.A)

	// CP a r8
	case 0b10111000:
		gameboy.CP8(&gameboy.cpu.registers.A, gameboy.cpu.registers.B)
	case 0b10111001:
		gameboy.CP8(&gameboy.cpu.registers.A, gameboy.cpu.registers.C)
	case 0b10111010:
		gameboy.CP8(&gameboy.cpu.registers.A, gameboy.cpu.registers.D)
	case 0b10111011:
		gameboy.CP8(&gameboy.cpu.registers.A, gameboy.cpu.registers.E)
	case 0b10111100:
		gameboy.CP8(&gameboy.cpu.registers.A, gameboy.cpu.registers.H)
	case 0b10111101:
		gameboy.CP8(&gameboy.cpu.registers.A, gameboy.cpu.registers.L)
	case 0b10111110:
		gameboy.CP8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
	case 0b10111111:
		gameboy.CP8(&gameboy.cpu.registers.A, gameboy.cpu.registers.A)

	// ADD a imm8
	case 0b11000110:
		gameboy.ADD8(&gameboy.cpu.registers.A, gameboy.POP_PC8())

	// ADC a imm8
	case 0b11001110:
		gameboy.ADC8(&gameboy.cpu.registers.A, gameboy.POP_PC8())

	// SUB a imm8
	case 0b11010110:
		gameboy.SUB8(&gameboy.cpu.registers.A, gameboy.POP_PC8())

	// SBC a imm8
	case 0b11011110:
		gameboy.SBC8(&gameboy.cpu.registers.A, gameboy.POP_PC8())

	// AND a imm8
	case 0b11100110:
		gameboy.AND8(&gameboy.cpu.registers.A, gameboy.POP_PC8())

	// XOR a imm8
	case 0b11101110:
		gameboy.XOR8(&gameboy.cpu.registers.A, gameboy.POP_PC8())

	// OR a imm8
	case 0b11110110:
		gameboy.OR8(&gameboy.cpu.registers.A, gameboy.POP_PC8())

	// CP a imm8
	case 0b11111110:
		gameboy.CP8(&gameboy.cpu.registers.A, gameboy.POP_PC8())

	// RET NZ
	case 0b11000000:
		if !gameboy.cpu.registers.GetZero() {
			gameboy.RET()
			// TODO add some cycles 12
		}

	// RET Z
	case 0b11001000:
		if gameboy.cpu.registers.GetZero() {
			gameboy.RET()
			// TODO add some cycles 12
		}

	// RET NC
	case 0b11010000:
		if !gameboy.cpu.registers.GetCarry() {
			gameboy.RET()
			// TODO add some cycles 12
		}

	// RET C
	case 0b11011000:
		if gameboy.cpu.registers.GetCarry() {
			gameboy.RET()
			// TODO add some cycles 12
		}

	// RET
	case 0b11001001:
		gameboy.RET()

	// RETI
	case 0b11011001:
		gameboy.RET()
		// TODO interruptsEnabling True

	// JP NZ
	case 0b11000010:
		next := gameboy.POP_PC16()
		if !gameboy.cpu.registers.GetZero() {
			gameboy.JUMP(next)
			// TODO add some cycles 4
		}

	// JP Z
	case 0b11001010:
		next := gameboy.POP_PC16()
		if gameboy.cpu.registers.GetZero() {
			gameboy.JUMP(next)
			// TODO add some cycles 4
		}

	// JP NC
	case 0b11010010:
		next := gameboy.POP_PC16()
		if !gameboy.cpu.registers.GetCarry() {
			gameboy.JUMP(next)
			// TODO add some cycles 4
		}

	// JP C
	case 0b11011010:
		next := gameboy.POP_PC16()
		if gameboy.cpu.registers.GetCarry() {
			gameboy.JUMP(next)
			// TODO add some cycles 4
		}

	// JP imm16
	case 0b11000011:
		next := gameboy.POP_PC16()
		gameboy.JUMP(next)

	// JP hl
	case 0b11101001:
		gameboy.JUMP(gameboy.cpu.registers.GetHL())

	// CALL NZ imm16
	case 0b11000100:
		next := gameboy.POP_PC16()
		if !gameboy.cpu.registers.GetZero() {
			gameboy.CALL(next)
			// TODO add some cycles 12
		}

	// CALL Z imm16
	case 0b11001100:
		next := gameboy.POP_PC16()
		if gameboy.cpu.registers.GetZero() {
			gameboy.CALL(next)
			// TODO add some cycles 12
		}

	// CALL NC imm16
	case 0b11010100:
		next := gameboy.POP_PC16()
		if !gameboy.cpu.registers.GetCarry() {
			gameboy.CALL(next)
			// TODO add some cycles 12
		}

	// CALL C imm16
	case 0b11011100:
		next := gameboy.POP_PC16()
		if gameboy.cpu.registers.GetCarry() {
			gameboy.CALL(next)
			// TODO add some cycles 12
		}

	// CALL imm16
	case 0b11001101:
		gameboy.CALL(gameboy.POP_PC16())
		// TODO add some cycles 12

	// RST tgt3
	case 0b11000111:
		gameboy.CALL(0x0000)
	case 0b11001111:
		gameboy.CALL(0x0008)
	case 0b11010111:
		gameboy.CALL(0x0010)
	case 0b11011111:
		gameboy.CALL(0x0018)
	case 0b11100111:
		gameboy.CALL(0x0020)
	case 0b11101111:
		gameboy.CALL(0x0028)
	case 0b11110111:
		gameboy.CALL(0x0030)
	case 0b11111111:
		gameboy.CALL(0x0038)

	// POP r16stk
	case 0b11000001:
		gameboy.cpu.registers.SetBC(gameboy.POPStack())
	case 0b11010001:
		gameboy.cpu.registers.SetDE(gameboy.POPStack())
	case 0b11100001:
		gameboy.cpu.registers.SetHL(gameboy.POPStack())
	case 0b11110001:
		gameboy.cpu.registers.SetAF(gameboy.POPStack())

	// PUSH r16stk
	case 0b11000101:
		gameboy.PUSH(gameboy.cpu.registers.GetBC())
	case 0b11010101:
		gameboy.PUSH(gameboy.cpu.registers.GetDE())
	case 0b11100101:
		gameboy.PUSH(gameboy.cpu.registers.GetHL())
	case 0b11110101:
		gameboy.PUSH(gameboy.cpu.registers.GetAF())

	// Prefix
	case 0b11001011:
		next := gameboy.POP_PC8()
		switch next {
		// RLC r8
		case 0b00000000:
			gameboy.RLC(&gameboy.cpu.registers.B)
		case 0b00000001:
			gameboy.RLC(&gameboy.cpu.registers.C)
		case 0b00000010:
			gameboy.RLC(&gameboy.cpu.registers.D)
		case 0b00000011:
			gameboy.RLC(&gameboy.cpu.registers.E)
		case 0b00000100:
			gameboy.RLC(&gameboy.cpu.registers.H)
		case 0b00000101:
			gameboy.RLC(&gameboy.cpu.registers.L)
		case 0b00000110:
			gameboy.RLC(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b00000111:
			gameboy.RLC(&gameboy.cpu.registers.A)

		// RRC r8
		case 0b00001000:
			gameboy.RRC(&gameboy.cpu.registers.B)
		case 0b00001001:
			gameboy.RRC(&gameboy.cpu.registers.C)
		case 0b00001010:
			gameboy.RRC(&gameboy.cpu.registers.D)
		case 0b00001011:
			gameboy.RRC(&gameboy.cpu.registers.E)
		case 0b00001100:
			gameboy.RRC(&gameboy.cpu.registers.H)
		case 0b00001101:
			gameboy.RRC(&gameboy.cpu.registers.L)
		case 0b00001110:
			gameboy.RRC(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b00001111:
			gameboy.RRC(&gameboy.cpu.registers.A)

		// RL r8
		case 0b00010000:
			gameboy.RL(&gameboy.cpu.registers.B)
		case 0b00010001:
			gameboy.RL(&gameboy.cpu.registers.C)
		case 0b00010010:
			gameboy.RL(&gameboy.cpu.registers.D)
		case 0b00010011:
			gameboy.RL(&gameboy.cpu.registers.E)
		case 0b00010100:
			gameboy.RL(&gameboy.cpu.registers.H)
		case 0b00010101:
			gameboy.RL(&gameboy.cpu.registers.L)
		case 0b00010110:
			gameboy.RL(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b00010111:
			gameboy.RL(&gameboy.cpu.registers.A)

		// RR r8
		case 0b00011000:
			gameboy.RR(&gameboy.cpu.registers.B)
		case 0b00011001:
			gameboy.RR(&gameboy.cpu.registers.C)
		case 0b00011010:
			gameboy.RR(&gameboy.cpu.registers.D)
		case 0b00011011:
			gameboy.RR(&gameboy.cpu.registers.E)
		case 0b00011100:
			gameboy.RR(&gameboy.cpu.registers.H)
		case 0b00011101:
			gameboy.RR(&gameboy.cpu.registers.L)
		case 0b00011110:
			gameboy.RR(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b00011111:
			gameboy.RR(&gameboy.cpu.registers.A)

		// SLA r8
		case 0b00100000:
			gameboy.SLA(&gameboy.cpu.registers.B)
		case 0b00100001:
			gameboy.SLA(&gameboy.cpu.registers.C)
		case 0b00100010:
			gameboy.SLA(&gameboy.cpu.registers.D)
		case 0b00100011:
			gameboy.SLA(&gameboy.cpu.registers.E)
		case 0b00100100:
			gameboy.SLA(&gameboy.cpu.registers.H)
		case 0b00100101:
			gameboy.SLA(&gameboy.cpu.registers.L)
		case 0b00100110:
			gameboy.SLA(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b00100111:
			gameboy.SLA(&gameboy.cpu.registers.A)

		// SRA r8
		case 0b00101000:
			gameboy.SRA(&gameboy.cpu.registers.B)
		case 0b00101001:
			gameboy.SRA(&gameboy.cpu.registers.C)
		case 0b00101010:
			gameboy.SRA(&gameboy.cpu.registers.D)
		case 0b00101011:
			gameboy.SRA(&gameboy.cpu.registers.E)
		case 0b00101100:
			gameboy.SRA(&gameboy.cpu.registers.H)
		case 0b00101101:
			gameboy.SRA(&gameboy.cpu.registers.L)
		case 0b00101110:
			gameboy.SRA(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b00101111:
			gameboy.SRA(&gameboy.cpu.registers.A)

		// SWAP r8
		case 0b00110000:
			gameboy.SWAP(&gameboy.cpu.registers.B)
		case 0b00110001:
			gameboy.SWAP(&gameboy.cpu.registers.C)
		case 0b00110010:
			gameboy.SWAP(&gameboy.cpu.registers.D)
		case 0b00110011:
			gameboy.SWAP(&gameboy.cpu.registers.E)
		case 0b00110100:
			gameboy.SWAP(&gameboy.cpu.registers.H)
		case 0b00110101:
			gameboy.SWAP(&gameboy.cpu.registers.L)
		case 0b00110110:
			gameboy.SWAP(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b00110111:
			gameboy.SWAP(&gameboy.cpu.registers.A)

		// SRL r8
		case 0b00111000:
			gameboy.SRL(&gameboy.cpu.registers.B)
		case 0b00111001:
			gameboy.SRL(&gameboy.cpu.registers.C)
		case 0b00111010:
			gameboy.SRL(&gameboy.cpu.registers.D)
		case 0b00111011:
			gameboy.SRL(&gameboy.cpu.registers.E)
		case 0b00111100:
			gameboy.SRL(&gameboy.cpu.registers.H)
		case 0b00111101:
			gameboy.SRL(&gameboy.cpu.registers.L)
		case 0b00111110:
			gameboy.SRL(gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b00111111:
			gameboy.SRL(&gameboy.cpu.registers.A)

		// bit b3 r8
		case 0b01000000:
			gameboy.BIT(0, gameboy.cpu.registers.B)
		case 0b01000001:
			gameboy.BIT(0, gameboy.cpu.registers.C)
		case 0b01000010:
			gameboy.BIT(0, gameboy.cpu.registers.D)
		case 0b01000011:
			gameboy.BIT(0, gameboy.cpu.registers.E)
		case 0b01000100:
			gameboy.BIT(0, gameboy.cpu.registers.H)
		case 0b01000101:
			gameboy.BIT(0, gameboy.cpu.registers.L)
		case 0b01000110:
			gameboy.BIT(0, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
		case 0b01000111:
			gameboy.BIT(0, gameboy.cpu.registers.A)
		case 0b01001000:
			gameboy.BIT(1, gameboy.cpu.registers.B)
		case 0b01001001:
			gameboy.BIT(1, gameboy.cpu.registers.C)
		case 0b01001010:
			gameboy.BIT(1, gameboy.cpu.registers.D)
		case 0b01001011:
			gameboy.BIT(1, gameboy.cpu.registers.E)
		case 0b01001100:
			gameboy.BIT(1, gameboy.cpu.registers.H)
		case 0b01001101:
			gameboy.BIT(1, gameboy.cpu.registers.L)
		case 0b01001110:
			gameboy.BIT(1, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
		case 0b01001111:
			gameboy.BIT(1, gameboy.cpu.registers.A)
		case 0b01010000:
			gameboy.BIT(2, gameboy.cpu.registers.B)
		case 0b01010001:
			gameboy.BIT(2, gameboy.cpu.registers.C)
		case 0b01010010:
			gameboy.BIT(2, gameboy.cpu.registers.D)
		case 0b01010011:
			gameboy.BIT(2, gameboy.cpu.registers.E)
		case 0b01010100:
			gameboy.BIT(2, gameboy.cpu.registers.H)
		case 0b01010101:
			gameboy.BIT(2, gameboy.cpu.registers.L)
		case 0b01010110:
			gameboy.BIT(2, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
		case 0b01010111:
			gameboy.BIT(2, gameboy.cpu.registers.A)
		case 0b01011000:
			gameboy.BIT(3, gameboy.cpu.registers.B)
		case 0b01011001:
			gameboy.BIT(3, gameboy.cpu.registers.C)
		case 0b01011010:
			gameboy.BIT(3, gameboy.cpu.registers.D)
		case 0b01011011:
			gameboy.BIT(3, gameboy.cpu.registers.E)
		case 0b01011100:
			gameboy.BIT(3, gameboy.cpu.registers.H)
		case 0b01011101:
			gameboy.BIT(3, gameboy.cpu.registers.L)
		case 0b01011110:
			gameboy.BIT(3, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
		case 0b01011111:
			gameboy.BIT(3, gameboy.cpu.registers.A)
		case 0b01100000:
			gameboy.BIT(4, gameboy.cpu.registers.B)
		case 0b01100001:
			gameboy.BIT(4, gameboy.cpu.registers.C)
		case 0b01100010:
			gameboy.BIT(4, gameboy.cpu.registers.D)
		case 0b01100011:
			gameboy.BIT(4, gameboy.cpu.registers.E)
		case 0b01100100:
			gameboy.BIT(4, gameboy.cpu.registers.H)
		case 0b01100101:
			gameboy.BIT(4, gameboy.cpu.registers.L)
		case 0b01100110:
			gameboy.BIT(4, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
		case 0b01100111:
			gameboy.BIT(4, gameboy.cpu.registers.A)
		case 0b01101000:
			gameboy.BIT(5, gameboy.cpu.registers.B)
		case 0b01101001:
			gameboy.BIT(5, gameboy.cpu.registers.C)
		case 0b01101010:
			gameboy.BIT(5, gameboy.cpu.registers.D)
		case 0b01101011:
			gameboy.BIT(5, gameboy.cpu.registers.E)
		case 0b01101100:
			gameboy.BIT(5, gameboy.cpu.registers.H)
		case 0b01101101:
			gameboy.BIT(5, gameboy.cpu.registers.L)
		case 0b01101110:
			gameboy.BIT(5, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
		case 0b01101111:
			gameboy.BIT(5, gameboy.cpu.registers.A)
		case 0b01110000:
			gameboy.BIT(6, gameboy.cpu.registers.B)
		case 0b01110001:
			gameboy.BIT(6, gameboy.cpu.registers.C)
		case 0b01110010:
			gameboy.BIT(6, gameboy.cpu.registers.D)
		case 0b01110011:
			gameboy.BIT(6, gameboy.cpu.registers.E)
		case 0b01110100:
			gameboy.BIT(6, gameboy.cpu.registers.H)
		case 0b01110101:
			gameboy.BIT(6, gameboy.cpu.registers.L)
		case 0b01110110:
			gameboy.BIT(6, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
		case 0b01110111:
			gameboy.BIT(6, gameboy.cpu.registers.A)
		case 0b01111000:
			gameboy.BIT(7, gameboy.cpu.registers.B)
		case 0b01111001:
			gameboy.BIT(7, gameboy.cpu.registers.C)
		case 0b01111010:
			gameboy.BIT(7, gameboy.cpu.registers.D)
		case 0b01111011:
			gameboy.BIT(7, gameboy.cpu.registers.E)
		case 0b01111100:
			gameboy.BIT(7, gameboy.cpu.registers.H)
		case 0b01111101:
			gameboy.BIT(7, gameboy.cpu.registers.L)
		case 0b01111110:
			gameboy.BIT(7, gameboy.memory.GetValue(gameboy.cpu.registers.GetHL()))
		case 0b01111111:
			gameboy.BIT(7, gameboy.cpu.registers.A)

		// RES b3 r8
		case 0b10000000:
			gameboy.RES(0, &gameboy.cpu.registers.B)
		case 0b10000001:
			gameboy.RES(0, &gameboy.cpu.registers.C)
		case 0b10000010:
			gameboy.RES(0, &gameboy.cpu.registers.D)
		case 0b10000011:
			gameboy.RES(0, &gameboy.cpu.registers.E)
		case 0b10000100:
			gameboy.RES(0, &gameboy.cpu.registers.H)
		case 0b10000101:
			gameboy.RES(0, &gameboy.cpu.registers.L)
		case 0b10000110:
			gameboy.RES(0, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b10000111:
			gameboy.RES(0, &gameboy.cpu.registers.A)
		case 0b10001000:
			gameboy.RES(1, &gameboy.cpu.registers.B)
		case 0b10001001:
			gameboy.RES(1, &gameboy.cpu.registers.C)
		case 0b10001010:
			gameboy.RES(1, &gameboy.cpu.registers.D)
		case 0b10001011:
			gameboy.RES(1, &gameboy.cpu.registers.E)
		case 0b10001100:
			gameboy.RES(1, &gameboy.cpu.registers.H)
		case 0b10001101:
			gameboy.RES(1, &gameboy.cpu.registers.L)
		case 0b10001110:
			gameboy.RES(1, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b10001111:
			gameboy.RES(1, &gameboy.cpu.registers.A)
		case 0b10010000:
			gameboy.RES(2, &gameboy.cpu.registers.B)
		case 0b10010001:
			gameboy.RES(2, &gameboy.cpu.registers.C)
		case 0b10010010:
			gameboy.RES(2, &gameboy.cpu.registers.D)
		case 0b10010011:
			gameboy.RES(2, &gameboy.cpu.registers.E)
		case 0b10010100:
			gameboy.RES(2, &gameboy.cpu.registers.H)
		case 0b10010101:
			gameboy.RES(2, &gameboy.cpu.registers.L)
		case 0b10010110:
			gameboy.RES(2, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b10010111:
			gameboy.RES(2, &gameboy.cpu.registers.A)
		case 0b10011000:
			gameboy.RES(3, &gameboy.cpu.registers.B)
		case 0b10011001:
			gameboy.RES(3, &gameboy.cpu.registers.C)
		case 0b10011010:
			gameboy.RES(3, &gameboy.cpu.registers.D)
		case 0b10011011:
			gameboy.RES(3, &gameboy.cpu.registers.E)
		case 0b10011100:
			gameboy.RES(3, &gameboy.cpu.registers.H)
		case 0b10011101:
			gameboy.RES(3, &gameboy.cpu.registers.L)
		case 0b10011110:
			gameboy.RES(3, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b10011111:
			gameboy.RES(3, &gameboy.cpu.registers.A)
		case 0b10100000:
			gameboy.RES(4, &gameboy.cpu.registers.B)
		case 0b10100001:
			gameboy.RES(4, &gameboy.cpu.registers.C)
		case 0b10100010:
			gameboy.RES(4, &gameboy.cpu.registers.D)
		case 0b10100011:
			gameboy.RES(4, &gameboy.cpu.registers.E)
		case 0b10100100:
			gameboy.RES(4, &gameboy.cpu.registers.H)
		case 0b10100101:
			gameboy.RES(4, &gameboy.cpu.registers.L)
		case 0b10100110:
			gameboy.RES(4, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b10100111:
			gameboy.RES(4, &gameboy.cpu.registers.A)
		case 0b10101000:
			gameboy.RES(5, &gameboy.cpu.registers.B)
		case 0b10101001:
			gameboy.RES(5, &gameboy.cpu.registers.C)
		case 0b10101010:
			gameboy.RES(5, &gameboy.cpu.registers.D)
		case 0b10101011:
			gameboy.RES(5, &gameboy.cpu.registers.E)
		case 0b10101100:
			gameboy.RES(5, &gameboy.cpu.registers.H)
		case 0b10101101:
			gameboy.RES(5, &gameboy.cpu.registers.L)
		case 0b10101110:
			gameboy.RES(5, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b10101111:
			gameboy.RES(5, &gameboy.cpu.registers.A)
		case 0b10110000:
			gameboy.RES(6, &gameboy.cpu.registers.B)
		case 0b10110001:
			gameboy.RES(6, &gameboy.cpu.registers.C)
		case 0b10110010:
			gameboy.RES(6, &gameboy.cpu.registers.D)
		case 0b10110011:
			gameboy.RES(6, &gameboy.cpu.registers.E)
		case 0b10110100:
			gameboy.RES(6, &gameboy.cpu.registers.H)
		case 0b10110101:
			gameboy.RES(6, &gameboy.cpu.registers.L)
		case 0b10110110:
			gameboy.RES(6, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b10110111:
			gameboy.RES(6, &gameboy.cpu.registers.A)
		case 0b10111000:
			gameboy.RES(7, &gameboy.cpu.registers.B)
		case 0b10111001:
			gameboy.RES(7, &gameboy.cpu.registers.C)
		case 0b10111010:
			gameboy.RES(7, &gameboy.cpu.registers.D)
		case 0b10111011:
			gameboy.RES(7, &gameboy.cpu.registers.E)
		case 0b10111100:
			gameboy.RES(7, &gameboy.cpu.registers.H)
		case 0b10111101:
			gameboy.RES(7, &gameboy.cpu.registers.L)
		case 0b10111110:
			gameboy.RES(7, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b10111111:
			gameboy.RES(7, &gameboy.cpu.registers.A)

		// SET b3 r8
		case 0b11000000:
			gameboy.SET(0, &gameboy.cpu.registers.B)
		case 0b11000001:
			gameboy.SET(0, &gameboy.cpu.registers.C)
		case 0b11000010:
			gameboy.SET(0, &gameboy.cpu.registers.D)
		case 0b11000011:
			gameboy.SET(0, &gameboy.cpu.registers.E)
		case 0b11000100:
			gameboy.SET(0, &gameboy.cpu.registers.H)
		case 0b11000101:
			gameboy.SET(0, &gameboy.cpu.registers.L)
		case 0b11000110:
			gameboy.SET(0, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b11000111:
			gameboy.SET(0, &gameboy.cpu.registers.A)
		case 0b11001000:
			gameboy.SET(1, &gameboy.cpu.registers.B)
		case 0b11001001:
			gameboy.SET(1, &gameboy.cpu.registers.C)
		case 0b11001010:
			gameboy.SET(1, &gameboy.cpu.registers.D)
		case 0b11001011:
			gameboy.SET(1, &gameboy.cpu.registers.E)
		case 0b11001100:
			gameboy.SET(1, &gameboy.cpu.registers.H)
		case 0b11001101:
			gameboy.SET(1, &gameboy.cpu.registers.L)
		case 0b11001110:
			gameboy.SET(1, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b11001111:
			gameboy.SET(1, &gameboy.cpu.registers.A)
		case 0b11010000:
			gameboy.SET(2, &gameboy.cpu.registers.B)
		case 0b11010001:
			gameboy.SET(2, &gameboy.cpu.registers.C)
		case 0b11010010:
			gameboy.SET(2, &gameboy.cpu.registers.D)
		case 0b11010011:
			gameboy.SET(2, &gameboy.cpu.registers.E)
		case 0b11010100:
			gameboy.SET(2, &gameboy.cpu.registers.H)
		case 0b11010101:
			gameboy.SET(2, &gameboy.cpu.registers.L)
		case 0b11010110:
			gameboy.SET(2, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b11010111:
			gameboy.SET(2, &gameboy.cpu.registers.A)
		case 0b11011000:
			gameboy.SET(3, &gameboy.cpu.registers.B)
		case 0b11011001:
			gameboy.SET(3, &gameboy.cpu.registers.C)
		case 0b11011010:
			gameboy.SET(3, &gameboy.cpu.registers.D)
		case 0b11011011:
			gameboy.SET(3, &gameboy.cpu.registers.E)
		case 0b11011100:
			gameboy.SET(3, &gameboy.cpu.registers.H)
		case 0b11011101:
			gameboy.SET(3, &gameboy.cpu.registers.L)
		case 0b11011110:
			gameboy.SET(3, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b11011111:
			gameboy.SET(3, &gameboy.cpu.registers.A)
		case 0b11100000:
			gameboy.SET(4, &gameboy.cpu.registers.B)
		case 0b11100001:
			gameboy.SET(4, &gameboy.cpu.registers.C)
		case 0b11100010:
			gameboy.SET(4, &gameboy.cpu.registers.D)
		case 0b11100011:
			gameboy.SET(4, &gameboy.cpu.registers.E)
		case 0b11100100:
			gameboy.SET(4, &gameboy.cpu.registers.H)
		case 0b11100101:
			gameboy.SET(4, &gameboy.cpu.registers.L)
		case 0b11100110:
			gameboy.SET(4, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b11100111:
			gameboy.SET(4, &gameboy.cpu.registers.A)
		case 0b11101000:
			gameboy.SET(5, &gameboy.cpu.registers.B)
		case 0b11101001:
			gameboy.SET(5, &gameboy.cpu.registers.C)
		case 0b11101010:
			gameboy.SET(5, &gameboy.cpu.registers.D)
		case 0b11101011:
			gameboy.SET(5, &gameboy.cpu.registers.E)
		case 0b11101100:
			gameboy.SET(5, &gameboy.cpu.registers.H)
		case 0b11101101:
			gameboy.SET(5, &gameboy.cpu.registers.L)
		case 0b11101110:
			gameboy.SET(5, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b11101111:
			gameboy.SET(5, &gameboy.cpu.registers.A)
		case 0b11110000:
			gameboy.SET(6, &gameboy.cpu.registers.B)
		case 0b11110001:
			gameboy.SET(6, &gameboy.cpu.registers.C)
		case 0b11110010:
			gameboy.SET(6, &gameboy.cpu.registers.D)
		case 0b11110011:
			gameboy.SET(6, &gameboy.cpu.registers.E)
		case 0b11110100:
			gameboy.SET(6, &gameboy.cpu.registers.H)
		case 0b11110101:
			gameboy.SET(6, &gameboy.cpu.registers.L)
		case 0b11110110:
			gameboy.SET(6, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		case 0b11110111:
			gameboy.SET(6, &gameboy.cpu.registers.A)
		case 0b11111000:
			gameboy.SET(7, &gameboy.cpu.registers.B)
		case 0b11111001:
			gameboy.SET(7, &gameboy.cpu.registers.C)
		case 0b11111010:
			gameboy.SET(7, &gameboy.cpu.registers.D)
		case 0b11111011:
			gameboy.SET(7, &gameboy.cpu.registers.E)
		case 0b11111100:
			gameboy.SET(7, &gameboy.cpu.registers.H)
		case 0b11111101:
			gameboy.SET(7, &gameboy.cpu.registers.L)
		case 0b11111111:
			gameboy.SET(7, &gameboy.cpu.registers.A)
		case 0b11111110:
			gameboy.SET(7, gameboy.memory.GetPointer(gameboy.cpu.registers.GetHL()))
		}

	// LDH [c] a
	case 0b11100010:
		gameboy.LD8(gameboy.memory.GetPointer(uint16(0xFF00)+uint16(gameboy.cpu.registers.C)), gameboy.cpu.registers.A)

	// LDH [imm8] a
	case 0b11100000:
		gameboy.LD8(gameboy.memory.GetPointer(uint16(0xFF00)+uint16(gameboy.POP_PC8())), gameboy.cpu.registers.A)

	// LD [imm16] a
	case 0b11101010:
		gameboy.LD8(gameboy.memory.GetPointer(gameboy.POP_PC16()), gameboy.cpu.registers.A)

	// LDH a [c]
	case 0b11110010:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(uint16(0xFF00)+uint16(gameboy.cpu.registers.B)))

	// LDH a [imm8]
	case 0b11110000:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(uint16(0xFF00)+uint16(gameboy.POP_PC8())))

	// LD a [imm16]
	case 0b11111010:
		gameboy.LD8(&gameboy.cpu.registers.A, gameboy.memory.GetValue(gameboy.POP_PC16()))

	// ADD sp imm8
	case 0b11101000:
		sp := gameboy.cpu.registers.GetSP()
		gameboy.ADD16Signed(&sp, int8(gameboy.POP_PC8()))
		gameboy.cpu.registers.SetSP(sp)

	// LD hl sp + imm8
	case 0b11111000:
		tmpValue := gameboy.cpu.registers.GetSP()
		gameboy.ADD16Signed(&tmpValue, int8(gameboy.POP_PC8()))
		gameboy.cpu.registers.SetHL(tmpValue)

	// LD sp hl
	case 0b11111001:
		gameboy.cpu.registers.SetSP(gameboy.cpu.registers.GetHL())

	// DI
	case 0b11110011:
		// TODO interruptsOn = false
		break

	// EI
	case 0b11111011:
		// TODO interruptsEnabling = true
		break
	}
}

func (gameboy *Gameboy) ADD8(destination *uint8, source uint8) {
	total := uint(source) + uint(*destination)
	*destination = uint8(total)

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry((source&0xF)+(*destination&0xF) > 0xF)
	gameboy.cpu.registers.SetCarry(total > 0xFF)
}

func (gameboy *Gameboy) ADC8(destination *uint8, source uint8) {
	carry := uint(0)
	if gameboy.cpu.registers.GetCarry() {
		carry = uint(1)
	}

	total := uint(source) + uint(*destination) + carry
	*destination = uint8(total)

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry((source&0xF)+(*destination&0xF) > 0xF)
	gameboy.cpu.registers.SetCarry(total > 0xFF)
}

func (gameboy *Gameboy) SUB8(destination *uint8, source uint8) {
	total := uint(*destination) - uint(source)
	*destination = uint8(total)

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(true)
	gameboy.cpu.registers.SetHalfCarry(int16(*destination&0x0f)-int16(source&0xF) < 0)
	gameboy.cpu.registers.SetCarry(total < 0)
}

func (gameboy *Gameboy) SBC8(destination *uint8, source uint8) {
	carry := uint(0)
	if gameboy.cpu.registers.GetCarry() {
		carry = uint(1)
	}
	total := uint(*destination) - uint(source) - carry
	*destination = uint8(total)

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(true)
	gameboy.cpu.registers.SetHalfCarry(int16(*destination&0x0f)-int16(source&0xF)-int16(carry) < 0)
	gameboy.cpu.registers.SetCarry(total < 0)
}

func (gameboy *Gameboy) AND8(destination *uint8, source uint8) {
	*destination = *destination & source

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(true)
	gameboy.cpu.registers.SetCarry(false)
}

func (gameboy *Gameboy) CP8(destination *uint8, source uint8) {
	result := *destination - source

	gameboy.cpu.registers.SetZero(result == 0)
	gameboy.cpu.registers.SetSubtract(true)
	gameboy.cpu.registers.SetHalfCarry((source & 0x0f) > (*destination & 0x0f))
	gameboy.cpu.registers.SetCarry(source > *destination)
}

func (gameboy *Gameboy) OR8(destination *uint8, source uint8) {
	*destination = *destination | source

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(false)
	gameboy.cpu.registers.SetCarry(false)
}

func (gameboy *Gameboy) XOR8(destination *uint8, source uint8) {
	*destination = *destination ^ source

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(false)
	gameboy.cpu.registers.SetCarry(false)
}

func (gameboy *Gameboy) INC8(destination *uint8) {
	*destination = *destination + 1

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry((*destination&0xF)+(1&0xF) > 0xF)
}

func (gameboy *Gameboy) DEC8(destination *uint8) {
	*destination = *destination - 1

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(true)
	gameboy.cpu.registers.SetHalfCarry(*destination&0x0F == 0)
}

func (gameboy *Gameboy) ADD16(destination *uint16, source uint16) {
	total := uint32(source) + uint32(*destination)
	*destination = uint16(total)

	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(uint32(*destination&0xFFF) > (total & 0xFFF))
	gameboy.cpu.registers.SetCarry(total > 0xFFFF)
}

func (gameboy *Gameboy) ADD16Signed(destination *uint16, source int8) {
	total := uint16(int32(*destination) + int32(source))
	*destination = total
	tmpVal := *destination ^ uint16(source) ^ total
	gameboy.cpu.registers.SetZero(false)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry((tmpVal & 0x10) == 0x10)
	gameboy.cpu.registers.SetCarry((tmpVal & 0x100) == 0x100)
}

func (gameboy *Gameboy) INC16(destination *uint16) {
	*destination = *destination + 1
}

func (gameboy *Gameboy) DEC16(destination *uint16) {
	*destination = *destination - 1
}

func (gameboy *Gameboy) BIT(bitIndex uint8, source uint8) {
	gameboy.cpu.registers.SetZero((source>>bitIndex)&1 == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(true)
}

func (gameboy *Gameboy) SET(bitIndex uint8, destination *uint8) {
	*destination = *destination | (1 << bitIndex)
}

func (gameboy *Gameboy) RES(bitIndex uint8, destination *uint8) {
	*destination = *destination & ^(1 << bitIndex)
}

func (gameboy *Gameboy) SWAP(destination *uint8) {
	swapped := *destination<<4&240 | *destination>>4
	*destination = swapped

	gameboy.cpu.registers.SetZero(swapped == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(false)
	gameboy.cpu.registers.SetCarry(false)
}

func (gameboy *Gameboy) RL(destination *uint8) {
	newCarry := *destination >> 7
	oldCarry := uint8(0)
	if gameboy.cpu.registers.GetCarry() {
		oldCarry = uint8(1)
	}
	*destination = (*destination<<1)&0xFF | oldCarry

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(false)
	gameboy.cpu.registers.SetCarry(newCarry == 1)
}

func (gameboy *Gameboy) RLC(destination *uint8) {
	carry := *destination >> 7
	*destination = (*destination<<1)&0xFF | carry

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(false)
	gameboy.cpu.registers.SetCarry(carry == 1)
}

func (gameboy *Gameboy) RR(destination *uint8) {
	newCarry := *destination & 1
	oldCarry := uint8(0)
	if gameboy.cpu.registers.GetCarry() {
		oldCarry = uint8(1)
	}
	*destination = (*destination >> 1) | (oldCarry << 7)

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(false)
	gameboy.cpu.registers.SetCarry(newCarry == 1)
}

func (gameboy *Gameboy) RRC(destination *uint8) {
	carry := *destination & 1
	*destination = (*destination >> 1) | (carry << 7)

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(false)
	gameboy.cpu.registers.SetCarry(carry == 1)
}

func (gameboy *Gameboy) SLA(destination *uint8) {
	carry := *destination >> 7
	*destination = *destination << 1 & 0xFF

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(false)
	gameboy.cpu.registers.SetCarry(carry == 1)
}

func (gameboy *Gameboy) SRA(destination *uint8) {
	*destination = (*destination & 128) | (*destination >> 1)

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(false)
	gameboy.cpu.registers.SetCarry(*destination&1 == 1)
}

func (gameboy *Gameboy) SRL(destination *uint8) {
	carry := *destination & 1
	*destination = *destination >> 1

	gameboy.cpu.registers.SetZero(*destination == 0)
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(false)
	gameboy.cpu.registers.SetCarry(carry == 1)
}

func (gameboy *Gameboy) LD8(destination *uint8, source uint8) {
	*destination = source
}

func (gameboy *Gameboy) LD16(destination *uint16, source uint16) {
	*destination = source
}

func (gameboy *Gameboy) POP_PC8() uint8 {
	opcode := gameboy.memory.GetValue(gameboy.cpu.registers.PC)
	gameboy.cpu.registers.PC++
	return opcode
}

func (gameboy *Gameboy) POP_PC16() uint16 {
	b1 := uint16(gameboy.POP_PC8())
	b2 := uint16(gameboy.POP_PC8())
	return b2<<8 | b1
}

func (gameboy *Gameboy) POPStack() uint16 {
	byte1 := uint16(gameboy.memory.GetValue(gameboy.cpu.registers.GetSP()))
	byte2 := uint16(gameboy.memory.GetValue(gameboy.cpu.registers.GetSP()+1)) << 8
	gameboy.cpu.registers.SetSP(gameboy.cpu.registers.GetSP() + 2)
	return byte1 | byte2
}

func (gameboy *Gameboy) PUSH(address uint16) {
	gameboy.memory.SetValue(gameboy.cpu.registers.GetSP()-1, uint8(uint16(address&0xFF00)>>8))
	gameboy.memory.SetValue(gameboy.cpu.registers.GetSP()-2, uint8(address&0xFF))
	gameboy.cpu.registers.PC = gameboy.cpu.registers.GetSP() - 2
}

func (gameboy *Gameboy) CALL(next uint16) {
	gameboy.PUSH(gameboy.cpu.registers.PC)
	gameboy.cpu.registers.PC = next
}

func (gameboy *Gameboy) JUMP(next uint16) {
	gameboy.cpu.registers.PC = next
}

func (gameboy *Gameboy) RET() {
	gameboy.cpu.registers.PC = gameboy.POPStack()
}

func (gameboy *Gameboy) SCF() {
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(false)
	gameboy.cpu.registers.SetCarry(true)
}

func (gameboy *Gameboy) CCF() {
	gameboy.cpu.registers.SetSubtract(false)
	gameboy.cpu.registers.SetHalfCarry(false)
	gameboy.cpu.registers.SetCarry(!gameboy.cpu.registers.GetCarry())
}

func (gameboy *Gameboy) CP(destination uint8, source uint8) {
	total := destination - source
	gameboy.cpu.registers.SetZero(total == 0)
	gameboy.cpu.registers.SetSubtract(true)
	gameboy.cpu.registers.SetHalfCarry((source & 0x0f) > (destination & 0x0f))
	gameboy.cpu.registers.SetCarry(source > destination)
}

func (gameboy *Gameboy) CPL() {
	gameboy.cpu.registers.A = 0xFF ^ gameboy.cpu.registers.A
	gameboy.cpu.registers.SetSubtract(true)
	gameboy.cpu.registers.SetHalfCarry(true)
}

func (gameboy *Gameboy) DAA() {
	if !gameboy.cpu.registers.GetSubtract() {
		if gameboy.cpu.registers.GetCarry() || gameboy.cpu.registers.A > 0x99 {
			gameboy.cpu.registers.A = gameboy.cpu.registers.A + 0x60
			gameboy.cpu.registers.SetCarry(true)
		}
		if gameboy.cpu.registers.GetHalfCarry() || gameboy.cpu.registers.A&0xF > 0x9 {
			gameboy.cpu.registers.A = gameboy.cpu.registers.A + 0x06
			gameboy.cpu.registers.SetHalfCarry(false)
		}
	} else if gameboy.cpu.registers.GetCarry() && gameboy.cpu.registers.GetHalfCarry() {
		gameboy.cpu.registers.A = gameboy.cpu.registers.A + 0x9A
		gameboy.cpu.registers.SetHalfCarry(false)
	} else if gameboy.cpu.registers.GetCarry() {
		gameboy.cpu.registers.A = gameboy.cpu.registers.A + 0xA0
	} else if gameboy.cpu.registers.GetHalfCarry() {
		gameboy.cpu.registers.A = gameboy.cpu.registers.A + 0xFA
		gameboy.cpu.registers.SetHalfCarry(false)
	}
	gameboy.cpu.registers.SetZero(gameboy.cpu.registers.A == 0)
}

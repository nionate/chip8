package chip8

type Instruction func(cpu *CPU, params uint16) error

/**

**/
func SE(cpu *CPU, params uint16) error {
	x := params >> 8 & 0xF
	kk := params & 0xFF
	if cpu.V[x] == uint8(kk) {
		cpu.PC += InstructionSize
	}
	return nil
}

func SNE(cpu *CPU, params uint16) error {
	x := params >> 8 & 0xF
	kk := params & 0xFF
	if cpu.V[x] != uint8(kk) {
		cpu.PC += InstructionSize
	}
	return nil
}

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

	//Skip next instruction if Vx = Vy.
	x = params >> 8 & 0xF
	y := params >> 4 & 0xF
	if cpu.V[x] == cpu.V[y] {
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

func LD(cpu *CPU, params uint16) error {
	x := params >> 8 & 0xF
	kk := params & 0xFF
	cpu.V[x] = uint8(kk)

	/*

		res1 := fmt.Sprintf("%x", kk)
		res2 := fmt.Sprintf("%x", cpu.V[x])

		¿Por que solo imprime en consola cuando el test falla?
		fmt.Printf("kk '%s'\n", res1)
		fmt.Printf("Vx '%s'\n", res2)*/
	return nil
}

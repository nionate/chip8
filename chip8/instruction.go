package chip8

import "fmt"

type Instruction func(cpu *CPU, params uint16) error

/**
3xkk - SE Vx, byte
Skip next instruction if Vx = kk.

The interpreter compares register Vx to kk,
and if they are equal,
increments the program counter by 2.

5xy0 - SE Vx, Vy
Skip next instruction if Vx = Vy.

The interpreter compares register Vx to register Vy,
and if they are equal,
increments the program counter by 2.
**/
func SE(cpu *CPU, params uint16) error {
	x := params >> 8 & 0xF
	kk := params & 0xFF
	if cpu.V[x] == uint8(kk) {
		cpu.PC += InstructionSize
	}
	fmt.Printf("%04X -> SE V%d, %02X\n", params, x, kk)
	return nil
}

func SE5(cpu *CPU, params uint16) error {
	//Skip next instruction if Vx = Vy.
	x := params >> 8 & 0xF
	y := params >> 4 & 0xF
	if cpu.V[x] == cpu.V[y] {
		cpu.PC += InstructionSize
	}
	fmt.Printf("%04X -> SE5 V%d, V%d\n", params, x, y)

	return nil
}

/*
4xkk - SNE Vx, byte
Skip next instruction if Vx != kk.

The interpreter compares register Vx to kk,
and if they are not equal,
increments the program counter by 2.
*/
func SNE(cpu *CPU, params uint16) error {
	x := params >> 8 & 0xF
	kk := params & 0xFF
	if cpu.V[x] != uint8(kk) {
		cpu.PC += InstructionSize
	}
	fmt.Printf("%04X -> SNE V%d, %02X\n", params, x, kk)
	return nil
}

/*
6xkk - LD Vx, byte
Set Vx = kk.

The interpreter puts the value kk into register Vx.
*/
func LD(cpu *CPU, params uint16) error {
	x := params >> 8 & 0xF
	kk := params & 0xFF
	cpu.V[x] = uint8(kk)

	fmt.Printf("%04X -> LD V%d, %02X\n", params, x, kk)
	/*

		res1 := fmt.Sprintf("%x", kk)
		res2 := fmt.Sprintf("%x", cpu.V[x])

		¿Por que solo imprime en consola cuando el test falla?
		fmt.Printf("kk '%s'\n", res1)
		fmt.Printf("Vx '%s'\n", res2)*/
	return nil
}

/*
7xkk - ADD Vx, byte
Set Vx = Vx + kk.

Adds the value kk to the value of register Vx,
then stores the result in Vx.
*/
func ADD(cpu *CPU, params uint16) error {
	x := params >> 8 & 0xF
	kk := params & 0xFF
	cpu.V[x] = uint8(cpu.V[x]) + uint8(kk)

	fmt.Printf("%04X -> ADD V%d, %02X\n", params, x, kk)
	return nil
}

func LD8(cpu *CPU, params uint16) error {
	x := params >> 8 & 0xF
	y := params >> 4 & 0xF

	//cpu.V[x] = cpu.V[y]

	res1 := fmt.Sprintf("%x", x)
	res2 := fmt.Sprintf("%x", cpu.V[y])

	fmt.Printf("x '%s'\n", res1)
	fmt.Printf("Vy '%s'\n", res2)

	return nil
}

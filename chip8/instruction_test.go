package chip8_test

import (
	"testing"

	"github.com/nionate/chip8/chip8"
)

func TestSETrue(t *testing.T) {
	var params uint16 = 0x34CC
	cpu := chip8.NewCPU()
	cpu.V[0x4] = 0xCC
	err := chip8.SE(cpu, params)
	if err != nil {
		t.Error(err)
	}

	if cpu.PC != chip8.MemoryStart+chip8.InstructionSize {
		t.Error("Expected to skip an instruction")
	}
}

func TestSEFalse(t *testing.T) {
	var params uint16 = 0x3D23
	cpu := chip8.NewCPU()
	cpu.V[0xD] = 0xBB
	err := chip8.SE(cpu, params)
	if err != nil {
		t.Error(err)
	}

	if cpu.PC == chip8.MemoryStart+chip8.InstructionSize {
		t.Error("Expected to skip an instruction")
	}
}

func TestSNETrue(t *testing.T) {
	var params uint16 = 0x44CC
	cpu := chip8.NewCPU()
	cpu.V[0x4] = 0xDB
	err := chip8.SNE(cpu, params)
	if err != nil {
		t.Error(err)
	}

	if cpu.PC != chip8.MemoryStart+chip8.InstructionSize {
		t.Error("Expected to skip an instruction")
	}
}

func TestSNEFalse(t *testing.T) {
	var params uint16 = 0x4D23
	cpu := chip8.NewCPU()
	cpu.V[0xD] = 0x23
	err := chip8.SNE(cpu, params)
	if err != nil {
		t.Error(err)
	}

	if cpu.PC == chip8.MemoryStart+chip8.InstructionSize {
		t.Error("Expected to skip an instruction")
	}
}

func TestSE5True(t *testing.T) {
	var params uint16 = 0x5D40
	cpu := chip8.NewCPU()
	cpu.V[0xD] = 0xF
	cpu.V[0x4] = 0xF
	err := chip8.SE5(cpu, params)
	if err != nil {
		t.Error(err)
	}

	if cpu.PC != chip8.MemoryStart+chip8.InstructionSize {
		//No se salto la instruccion
		t.Error("Expected to skip an instruction")
	}
}

func TestSE5False(t *testing.T) {
	var params uint16 = 0x5D40
	cpu := chip8.NewCPU()
	cpu.V[0xD] = 0xF
	cpu.V[0x4] = 0xD
	err := chip8.SE5(cpu, params)
	if err != nil {
		t.Error(err)
	}

	if cpu.PC == chip8.MemoryStart+chip8.InstructionSize {
		//No se salto la instruccion
		t.Error("Expected to skip an instruction")
	}
}

func TestLDTrue(t *testing.T) {
	var params uint16 = 0x6D40
	cpu := chip8.NewCPU()
	err := chip8.LD(cpu, params)
	if err != nil {
		t.Error(err)
	}

	if cpu.V[0xD] != 0x40 {
		t.Error("Expected Vx to be equals to kk")
	}
}

/*func TestLDFalse(t *testing.T) {
	var params uint16 = 0x6D40
	cpu := chip8.NewCPU()
	err := chip8.LD(cpu, params)
	if err != nil {
		t.Error(err)
	}

	if cpu.V[0xD] == 0x40 {
		t.Error("Expected Vx to be not equals to kk")
	}
}*/

func TestADDTrue(t *testing.T) {
	var params uint16 = 0x7D40
	cpu := chip8.NewCPU()
	cpu.V[0xD] = uint8(0x5)
	err := chip8.ADD(cpu, params)

	if err != nil {
		t.Error(err)
	}

	if cpu.V[0xD] != (uint8(0x5) + uint8(0x40)) {
		t.Error("Expected Vx to be equals to Vx + kk")
	}
}

func TestLD8True(t *testing.T) {
	var params uint16 = 0x8D40
	cpu := chip8.NewCPU()
	err := chip8.LD8(cpu, params)
	cpu.V[0x4] = 0xDD

	if err != nil {
		t.Error(err)
	}

	if cpu.V[0xD] != cpu.V[0x4] {
		t.Error("Expected Vx to be equals to Vy")
	}
}

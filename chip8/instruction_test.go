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
	var params uint16 = 0x34CC
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
	var params uint16 = 0x3D23
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

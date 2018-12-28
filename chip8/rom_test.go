package chip8_test

import (
	"bytes"
	"testing"

	"github.com/nionate/chip8/chip8"
)

func TestShouldNotWriteOnROM(t *testing.T) {
	rom := chip8.NewROM()
	err := rom.Write(5, []byte{1, 2, 3, 4})

	if err == nil {
		t.Error(err)
	}
	//fmt.Println(ram)
}

func TestShouldNotReadFromROMWhenPositionIsLargerThanMemory(t *testing.T) {
	rom := chip8.NewROM()
	rom[62] = 4
	rom[63] = 2

	_, err := rom.Read(61, 4)
	if err == nil {
		t.Error(err)
	}
}

func TestShouldReadFromROM(t *testing.T) {
	rom := chip8.NewROM()
	rom[12] = 9
	rom[13] = 8

	value, err := rom.Read(11, 4)
	if err != nil {
		t.Error(err)
	}

	expected := []byte{0, 9, 8, 0}
	if !bytes.Equal(expected, value) {
		t.Errorf("%x should be %x", value, expected)
	}
}

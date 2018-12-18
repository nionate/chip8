package chip8

import (
	"errors"
	"io/ioutil"
)

type Cartridge struct {
	RAM
	Filename string
}

func (c Cartridge) Write(pos int16, data []byte) error {
	return errors.New("unsupported operation. You can't write to a cartridge")
}

func LoadCartridge(filename string) (*Cartridge, error) {	
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	cart := &Cartridge {
		Filename: filename,
		RAM: data,
	}

	return cart, nil
}

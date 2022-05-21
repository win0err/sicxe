package utils

import "github.com/win0err/sicxe/vm"

func BytesToWord(b []vm.Byte) vm.Word {
	return vm.Word(int(b[2]) | int(b[1])<<8 | int(b[0])<<16)
}

func WordToBytes(w vm.Word) []vm.Byte {
	data := []vm.Byte{0, 0, 0}
	data[0] = vm.Byte((w >> 16) & 0xFF)
	data[1] = vm.Byte((w >> 8) & 0xFF)
	data[2] = vm.Byte(w & 0xFF)

	return data
}

package main

import (
	"os"

	"github.com/win0err/sicxe/reader"
	"github.com/win0err/sicxe/utils"
	"github.com/win0err/sicxe/vm"
)

func main() {
	//objFile, err := os.Open("print.obj")
	objFile, err := os.Open("drawtext.obj")
	if err != nil {
		panic(err)
	}

	rd := reader.NewReader(objFile)

	registers := vm.NewRegisters()
	memory := vm.NewMemory()

	obj, err := reader.ReadAsObject(rd)
	if err != nil {
		panic(err)
	}

	for _, tr := range obj.TextRecords {
		copy(memory[tr.StartingAddress:], tr.Code)
	}

	utils.DumpMemory(memory, 0, 80)
	utils.DumpRegisters(*registers)
}

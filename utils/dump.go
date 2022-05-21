package utils

import (
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/win0err/sicxe/vm"
)

func DumpRegisters(r vm.Registers) {
	v := reflect.ValueOf(r)
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf(
			"%s:\t%08b\n",
			v.Type().Field(i).Name,
			WordToBytes(int(v.Field(i).Int())),
		)
	}
}

func DumpFullMemory(m vm.Memory) {
	fmt.Println(hex.Dump(m))
}

func DumpMemory(m vm.Memory, from, to int) {
	fmt.Println(hex.Dump(m[from:to]))
}

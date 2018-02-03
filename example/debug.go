package main


//#include <memory.h>
import "C"

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var a []uint8
	for i := 0; i <= math.MaxUint8; i++{
		a = append(a, uint8(i))
	}
	var ptr = unsafe.Pointer(&a[0])

	offset := 16
	data := []uint8{0xDE, 0xAD, 0xBE, 0xEF}
	C.memcpy(unsafe.Pointer(uintptr(ptr) + uintptr(offset)) , unsafe.Pointer(&data[0]), C.size_t(len(data)))
	for i := 0; i < 16; i++{
		fmt.Printf("%-4d : ", i * 16)
		for j:= 0;j < 16; j++{
			fmt.Printf("%02x ", a[16 * i + j])
		}
		fmt.Println()

	}

}
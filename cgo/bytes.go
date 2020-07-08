package main

/*
#include <stdint.h>

static void fill_255(char* buf, int32_t len) {
	int32_t i;
	for (i = 0; i < len; i++) {
		buf[i] = 255;
	}
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	b := make([]byte, 5)
	fmt.Println(b) // [0 0 0 0 0]

	var cptr = unsafe.Pointer(&b[0])
	C.fill_255((*C.char)(cptr), C.int32_t(len(b)))
	fmt.Println(b) // [255 255 255 255 255]
}

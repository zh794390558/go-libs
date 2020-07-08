package main

/*
#include <stdio.h>

void FloatPrint(void* ptr, int len) {
   float *data = (float*)(ptr);
   printf("len: %d\n", len);
   for (int i = 0; i < len; i++){
      printf("%f\t", data[i]);
   }
   printf("\n");
}

void IntPrint(void* ptr, int len) {
   int *data = (int*)(ptr);
   printf("len: %d\n", len);
   for (int i = 0; i < len; i++){
      printf("%d\t", data[i]);
   }
   printf("\n");
}

void CharPrint(void* ptr, int len) {
   char *data = (char*)(ptr);
   printf("len: %d\n", len);
   for (int i = 0; i < len; i++){
      printf("%d\t", data[i]);
   }
   printf("\n");
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	data := []float32{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(data)
	C.FloatPrint(unsafe.Pointer(&data[0]), C.int(len(data)))

	fmt.Printf("v%\n", data)
	fmt.Println()
	fmt.Println()
	fmt.Println(fmt.Sprintf("v%", data))
	fmt.Println()

	byteKey := []byte(fmt.Sprintf("%v", data))
	fmt.Println(byteKey)
	fmt.Printf("v%\n", byteKey)
	C.CharPrint(unsafe.Pointer(&byteKey[0]), C.int(len(byteKey)))

	fmt.Println()
	fmt.Println()

	data1 := []int{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(data1)
	C.IntPrint(unsafe.Pointer(&data1[0]), C.int(len(data1)))

	fmt.Printf("v%\n", data1)
	fmt.Println()
	fmt.Println()
	fmt.Println(fmt.Sprintf("v%", data1))
	fmt.Println()

	byteKey = []byte(fmt.Sprintf("%v", data1))
	fmt.Println(byteKey)
	fmt.Printf("v%\n", byteKey)
	C.CharPrint(unsafe.Pointer(&byteKey[0]), C.int(len(byteKey)))

	byteKey = []byte(data1)
	fmt.Println(byteKey)

}

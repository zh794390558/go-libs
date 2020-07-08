package main

/*
#include <stdio.h>
#include <stdlib.h>
void IntPrint(void* ptr, int len) {
   int *data = (int*)(ptr);
   printf("len: %d\n", len);
   for (int i = 0; i < len; i++){
      printf("%d\t", data[i]);
   }
   printf("\n");
}
*/
import "C"

import (
    "bytes"
    "encoding/binary"
    "unsafe"
    "fmt"
    "encoding/gob"
)

func main () {
    aaa := [12]int32{10, 3, 43, 434, 5}
    fmt.Println(aaa)
    bytesBuffer := bytes.NewBuffer([]byte{})
    for i:=0; i < len(aaa); i++ {
        err := binary.Write(bytesBuffer, binary.LittleEndian, aaa[i])
        if err != nil {
            fmt.Println("err:", err)
        }
    }
    by := bytesBuffer.Bytes()
    fmt.Println(by)
    C.IntPrint(unsafe.Pointer(&by[0]), C.int(len(by)/4))

    type A struct {
            // should be exported member when read back from buffer
                One int32
                    Two int32
    }
    var a A
    a.One = int32(1)
    a.Two = int32(2)
    buf := new(bytes.Buffer)
    fmt.Println("a's size is ",binary.Size(a))
    binary.Write(buf,binary.LittleEndian,a)
    fmt.Println("after write ，buf is:",buf.Bytes())

    fmt.Println("-------gob-------")
    bbb := []int{10, 3, 43, 434, 5}
    fmt.Println("raw: ", bbb)
    buf = bytes.NewBuffer([]byte{})
    enc := gob.NewEncoder(buf)
    dec := gob.NewDecoder(buf)

    enc.Encode(bbb)
    by = bytesBuffer.Bytes()
    C.IntPrint(unsafe.Pointer(&by[0]), C.int(len(by)/int(unsafe.Sizeof(bbb[0]))))

    var m2 []int
    dec.Decode(&m2)
    fmt.Println("gob decode:", m2)


    fmt.Println("-------gob interface-------")
    raw := []int{10, 3, 43, 434, 5}
    var bbbn []interface{} = make([]interface{}, len(raw))
    for i, d := range raw {
       bbb[i] = d
    }
    fmt.Println("raw: ", bbbn)
    buf = bytes.NewBuffer([]byte{})
    enc = gob.NewEncoder(buf)
    dec = gob.NewDecoder(buf)

    enc.Encode(bbbn)
    by = bytesBuffer.Bytes()
    C.IntPrint(unsafe.Pointer(&by[0]), C.int(len(by)/int(unsafe.Sizeof(bbbn[0].(int)))))

    var m3 []int
    dec.Decode(&m3)
    fmt.Println("gob decode:", m3)

}

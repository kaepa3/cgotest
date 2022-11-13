package main

import (
	"fmt"
	"unsafe"
)

/*
#cgo LDFLAGS: -L. -lhello
#include <stdlib.h>
#include "../include/hello.h"
*/
import "C"

func main() {
	fmt.Println("hoge")
	C.hello()
	C.helloint(2)

	//
	str := "world"
	p := C.CString(str)
	defer C.free(unsafe.Pointer(p))
	C.hellostr(p)
}

package main

import (
	"fmt"
)

/*
#cgo LDFLAGS: -L. -lhello
#include "../include/hello.h"
*/
import "C"

func main() {
	fmt.Println("hoge")
	C.hello()
}

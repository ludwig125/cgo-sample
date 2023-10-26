package main

//#cgo LDFLAGS: -L. -ladd
//#include "libadd.h"
import "C"
import "fmt"

func main() {
	x := C.Add(1, 2)
	fmt.Println(x)
}

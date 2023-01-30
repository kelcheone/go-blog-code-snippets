package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

// use SwapPointer

func main() {
	var i int64 = 1
	var p unsafe.Pointer = unsafe.Pointer(&i)
	fmt.Println(atomic.AddPointer(&p, 1))
	fmt.Println(*(*int64)(p))
}

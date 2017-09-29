package main

import (
	"fmt"
	"unsafe"
)

type Persion struct {
	Name string
	Age  int
}

func main() {
	v := Persion{Name: "panlin", Age: 26}
	p := &Persion{Name: "panlin", Age: 26}

	fmt.Printf("v : sizeof(%d), member: %#v\noffset.Name: %d, offset.Age: %d\nalign: %d\n", unsafe.Sizeof(v), v, unsafe.Offsetof(v.Name), unsafe.Offsetof(v.Age), unsafe.Alignof(v))
	fmt.Printf("p : sizeof(%d), member: %#v\naddr: 0x%x\n", unsafe.Sizeof(p), p, uintptr(unsafe.Pointer(p)))

	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&v)) + unsafe.Offsetof(v.Age)))
	*age = 56

	fmt.Printf("v.Name: %d\n", v.Age)
}

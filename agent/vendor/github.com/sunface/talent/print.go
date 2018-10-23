package talent

import (
	"fmt"
	"unsafe"
)

//打印浮点数的特定表现格式
func Float64Bits(f float64, d int) {
	b := *(*uint64)(unsafe.Pointer(&f))

	switch d {
	case 16:
		fmt.Printf("浮点数%.1f的16进制表示是%016x\n", f, b)
	case 2:
		fmt.Printf("浮点数%.1f的2进制表示是%02b\n", f, b)
	default:
		fmt.Println("error decimal: ", d)
	}
}

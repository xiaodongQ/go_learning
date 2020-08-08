package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestSliceReflect(t *testing.T) {
	var a []int
	b := []int{}

	// &reflect.SliceHeader{Data:0x0, Len:0, Cap:0}
	// 指向一个空地址，nil
	fmt.Printf("a: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))
	// &reflect.SliceHeader{Data:0x6e97b8, Len:0, Cap:0}
	// 指向一个具体地址，已经分配了内存，进行了初始化操作
	fmt.Printf("b: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)))
}

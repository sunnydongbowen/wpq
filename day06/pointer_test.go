package day06

import (
	"fmt"
	"testing"
	"unsafe"
)

// @program:   GoPrimary
// @file:      pointer_test.go
// @author:    bowen
// @time:      2024-11-18 20:23
// @description:

func TestPointer(t *testing.T) {
	v1 := "bowen"
	v2 := &v1
	fmt.Println(v1, v2, *v2)

	v1 = "hhh"
	fmt.Println(v1, v2, *v2)
}

func TestPointer2(t *testing.T) {
	// 数组指针
	// 注意如果是切片的话，打印的地址就不会一样了！
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	//fmt.Println(&a, &a[0])
	fmt.Printf("%p, %p\n", &a, &a[0])
}

func TestPointer3(t *testing.T) {
	var name string = "bowen"
	var pointer *string = &name
	var age int = 18
	var x1 *int = &age
	fmt.Println(pointer, x1)
}

func TestPointer4(t *testing.T) {
	v1 := "www"
	v2 := v1
	v1 = "hhh"
	fmt.Println(v1, v2)
}

func TestPointer5(t *testing.T) {
	v1 := "www"
	v2 := &v1
	v1 = "hhh"
	fmt.Println(v1, *v2)
}

func changedata(data string) {
	data = "bowen"
}

func changedata2(data *string) {
	*data = "bowen"
}

func TestPointer6(t *testing.T) {
	var data string = "www"
	// not changed
	//changedata(data)
	// 注意传的指针
	changedata2(&data)
	fmt.Println(data)
}

func TestPointer7(t *testing.T) {
	datalist := [3]int8{11, 22, 33}
	fmt.Printf("%p, %p,%T,%T\n", &datalist, &datalist[0], &datalist, &datalist[0])

	// get first data pointer address
	var firstDataPtr *int8 = &datalist[0]

	// convert to unsafe.Pointer
	ptr := unsafe.Pointer(firstDataPtr)

	// get target address，convert uintf.Pointer to uintptr
	// 相当于转换成整型+1
	targetAddress := uintptr(ptr) + 1

	// convert to unsafe.Pointer
	// 相当于转换成指针类型
	newPtr := unsafe.Pointer(targetAddress)

	// convert to int8 pointer
	value := (*int8)(newPtr)

	fmt.Println(*value)

}

func TestPointer8(t *testing.T) {
	name := "bowne"
	var p1 *string = &name
	var p2 **string = &p1
	var p3 ***string = &p2
	*p1 = "张三"
	**p2 = "李四"
	***p3 = "王五"
	fmt.Println(name, &name)
	fmt.Println(p1, &p1)
	fmt.Println(p2, &p2)
	fmt.Println(p3, &p3)
}

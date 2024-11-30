package day06

import (
	"fmt"
	"testing"
)

// @program:   GoPrimary
// @file:      slice_test.go
// @author:    bowen
// @time:      2024-11-18 15:20
// @description:

func TestSlice(t *testing.T) {
	var users = make([]string, 0, 10)
	var v1 = new([]int)
	var v2 *[]int
	fmt.Println(users, v1, v2)

	var v3 = make([]int, 3)
	// cap and len both are 3
	fmt.Println(cap(v3), len(v3))

	v3 = append(v3, 1, 2, 3)
	fmt.Println(cap(v3), len(v3))
	// cap is 6, len is 6，and the 0~2 index elem
	fmt.Println(v3)
}

func TestSlice2(t *testing.T) {
	var s1 = []int{1, 2, 3, 4, 5}
	deleteindex := 2
	s1 = append(s1[:deleteindex], s1[deleteindex+1:]...)
	fmt.Println(s1)
}

// slice的插入
func TestSliceInsert(t *testing.T) {
	var s1 = []int{1, 2, 3, 4, 5}
	insertindex := 2
	s1 = append(s1[:insertindex], append([]int{100}, s1[insertindex:]...)...)
	fmt.Println(s1)
}

func TestSlice11(t *testing.T) {
	var b1 bool = false
	b2 := b1
	fmt.Printf("%p, %p\n", &b1, &b2)

	var s1 string = "hello"
	s2 := s1
	fmt.Printf("%p, %p\n", &s1, &s2)

	var s3 []int = []int{1, 2, 3, 4, 5}
	s4 := s3
	fmt.Printf("%p, %p\n", &s3, &s4)
	s4[0] = 100
	fmt.Println(s3, s4)
	fmt.Printf("%p, %p\n", &s3, &s4)
	s3 = append(s3, 1399)
	fmt.Printf("%p, %p\n", &s3, &s4)

	var a1 = [3]int{1, 2, 3}
	a2 := a1
	a2[1] = 200
	fmt.Println(a1, a2)
}

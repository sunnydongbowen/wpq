package day04

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

// @program:   GoPrimary
// @file:      type1_test.go
// @author:    bowen
// @time:      2024-11-16 20:14
// @description:

func TestType1(t *testing.T) {
	var v1 int16 = 10
	v2 := int8(v1)
	fmt.Println(reflect.TypeOf(v1), reflect.TypeOf(v2))
}

func TestType2(t *testing.T) {
	var v1 int = 10
	s := strconv.Itoa(v1)
	fmt.Println(reflect.TypeOf(v1), reflect.TypeOf(s))
}

func TestType3(t *testing.T) {
	var s1 = "66"
	v1, _ := strconv.Atoi(s1)
	fmt.Println(reflect.TypeOf(v1), reflect.TypeOf(s1))
}

// 进制转换
func TestType4(t *testing.T) {
	var v1 = 5
	// int->binary
	r1 := strconv.FormatInt(int64(v1), 2)
	fmt.Println(r1, reflect.TypeOf(r1))

	data := "10001"
	// binary->int64
	v2, _ := strconv.ParseInt(data, 2, 64)
	fmt.Println(v2, reflect.TypeOf(v2))

	// formatInt must 传int64
	v3 := strconv.FormatInt(14, 16)
	fmt.Println(v3, reflect.TypeOf(v3))

}

func TestType5(t *testing.T) {
	res, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res, reflect.TypeOf(res))
}

func TestType6(t *testing.T) {
	res := strconv.FormatBool(true)
	fmt.Println(res, reflect.TypeOf(res))

}

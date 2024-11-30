package day04

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

// @program:   GoPrimary
// @file:      string_test.go
// @author:    bowen
// @time:      2024-11-22 21:09
// @description:

func TestString(t *testing.T) {
	//
	msg := "我爱你" + "bovan"
	fmt.Println(msg)

	// 字符串拼接
	dataList := []string{"a", "b", "c"}
	res := strings.Join(dataList, "-")
	fmt.Println(res)

	// 不常用
	var buffer bytes.Buffer
	buffer.WriteString("a")
	buffer.WriteString("b")
	buffer.WriteString("c")
	data := buffer.String()
	fmt.Println(data)

	// more efficent
	var builder strings.Builder
	builder.WriteString("hhhh")
	builder.WriteString("ddd")
	v := builder.String()
	fmt.Println(v)

}

func TestString2(t *testing.T) {

	//fmt.Println(string(65))

	v3, size := utf8.DecodeRuneInString("A")
	fmt.Println(v3, size)

	v4, size := utf8.DecodeRuneInString("文")
	fmt.Println(v4, size)

}

func TestString3(t *testing.T) {
	fmt.Println("a" == "a")
	var name string = "你好"
	//
	v1 := name[0] //228
	fmt.Println(v1)

	v2 := name[0:3]
	fmt.Println(v2) //你

	// get all byte
	for i := 0; i < len(name); i++ {
		fmt.Println(i, name[i])
	}

	// for range get all char
	for index, item := range name {
		fmt.Println(index, item, string(item))
	}

	//convert to rune
	dataList := []rune(name)
	fmt.Println(dataList[0], string(dataList[0]))

}

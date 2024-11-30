package funcc

import (
	"fmt"
	"testing"
)

// @program:   GoPrimary
// @file:      funcc4_test.go
// @author:    bowen
// @time:      2024-11-25 12:19
// @description:

func F1(n1 int, n2 int) func(int) string {
	// 不定义，直接用匿名函数返回
	return func(n1 int) string {
		return "success"
	}
}

func TestFuncc(t *testing.T) {
	v1 := func(n1 int, n2 int) int {
		return 123
	}
	data := v1(11, 22)
	fmt.Println(data)

	v2 := func(n1 int, n2 int) int {
		return 1
	}(11, 22)
	fmt.Println(v2)
}

package funcc

import (
	"fmt"
	"testing"
)

// @program:   GoPrimary
// @file:      funcc3_test.go
// @author:    bowen
// @time:      2024-11-25 12:06
// @description:

func exec(num1 int, num2 int) string {
	fmt.Println("执行函数")
	return "success"
}

type f2 func(int, int) string

func getFunc() f2 {
	return exec
}

func TestGetFunc(t *testing.T) {
	function := getFunc()
	res := function(10, 20)
	fmt.Println(res)
}

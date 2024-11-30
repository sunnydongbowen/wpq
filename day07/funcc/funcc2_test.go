package funcc

import (
	"fmt"
	"testing"
)

// @program:   GoPrimary
// @file:      funcc2_test.go
// @author:    bowen
// @time:      2024-11-25 11:13
// @description:

type f1 func(arg int) (int, bool)

func proxy2(data int, exec f1) int {
	data, flag := exec(data)
	if flag {
		return data
	} else {
		return 9999
	}
}

func TestProxy2(t *testing.T) {
	res := proxy2(10, add100)
	fmt.Println(res)
}

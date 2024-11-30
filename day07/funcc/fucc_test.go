package funcc

import (
	"fmt"
	"testing"
)

// @program:   GoPrimary
// @file:      fucc_test.go
// @author:    bowen
// @time:      2024-11-25 11:09
// @description:

func add100(arg int) (int, bool) {
	return arg + 100, true
}

func proxy(data int, exec func(int) (int, bool)) int {
	data, flag := exec(data)
	if flag {
		return data
	} else {
		return 9999
	}
}

func TestProxy(t *testing.T) {
	//传的时候传函数名就可以了
	res := proxy(10, add100)
	fmt.Println(res)

}

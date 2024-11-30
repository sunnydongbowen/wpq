package funcc

import (
	"fmt"
	"testing"
)

// @program:   GoPrimary
// @file:      funcc5_test.go
// @author:    bowen
// @time:      2024-11-25 15:21
// @description:

func TestBibao(t *testing.T) {
	var functionList []func()
	for i := 0; i < 5; i++ {
		function := func() {
			fmt.Println(i)
		}
		functionList = append(functionList, function)
	}
	// 运行的时候内部代码才会执行，在执行函数前，for循环已经执行完了。此时i已经=5
	//
	for _, function := range functionList {
		function()
	}
	functionList[0]()
	functionList[1]()
	functionList[2]()
}

func TestBibao2(t *testing.T) {
	var functionList []func()
	for i := 0; i < 5; i++ {
		function := func(arg int) func() {
			//fmt.Println(i)
			//函数内部用函数外部的变量
			return func() {
				fmt.Println(arg)
			}
		}(i)
		functionList = append(functionList, function)
	}
	functionList[0]()
	functionList[1]()
	functionList[2]()
}

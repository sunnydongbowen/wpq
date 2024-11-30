package main

import "fmt"

// @program:   GoPrimary
// @file:      main.go
// @author:    bowen
// @time:      2024-11-15 19:05
// @description:

func main() {
	var a int
	fmt.Print("请输入数字:")
	fmt.Scanln(&a)
	if a%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}

package main

import (
	"fmt"
)

// @program:   GoPrimary
// @file:      main.go
// @author:    bowen
// @time:      2024-11-15 20:49
// @description:guess number game

func main() {
	data := 66
	for {
		// num must declare inner for
		var num int
		fmt.Print("请输入数字:")
		fmt.Scan(&num)
		if num < data {
			fmt.Println("小了，请继续猜")
		} else if num > data {
			fmt.Println("大了，请继续猜")
		} else {
			fmt.Println("猜对了")
			break
		}
	}
}

package main

import "fmt"

// @program:   GoPrimary
// @file:      main.go
// @authotr:    bowen
// @time:      2024-11-16 11:40
// @description:

func main() {
	var name string
	fmt.Print("请输入你的名字：")
	fmt.Scanln(&name)
	if name == "bowensip" {
		goto svip
	} else if name == "bowenvip" {
		goto vip
	}
	fmt.Println("普通用户请耐心等待")
vip:
	fmt.Println("欢迎来到VIP界面")
svip:
	fmt.Println("欢迎来到SVIP界面")
}

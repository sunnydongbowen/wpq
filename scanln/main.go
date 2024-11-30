package main

import "fmt"

// @program:   GoPrimary
// @file:      main.go
// @author:    bowen
// @time:      2024-11-15 19:10
// @description:

func main() {
	var name string
	var passwd string
	fmt.Print("请输入用户名：")
	fmt.Scanln(&name)
	fmt.Print("请输入密码：")
	fmt.Scanln(&passwd)

	if name == "admin" && passwd == "123456" {
		fmt.Println("登录成功")
	} else {
		fmt.Println("登录失败")
	}
}

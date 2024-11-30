package main

import "fmt"

// @program:   GoPrimary
// @file:      main.go
// @author:    bowen
// @time:      2024-11-24 21:44
// @description:

type School struct {
	band string
	city string
}

type Class struct {
	title  string
	count  int16
	school *School
}

func main() {
	sch := &School{"路飞学城", "北京"}

	// 循环创建班级
	var classList []Class

	for {
		var cls Class
		fmt.Print("请输入班级和人数")
		//var title  string
		//var count int16
		fmt.Scanln(&cls.title, &cls.count)
		if cls.title == "Q" {
			break
		}
		cls.school = sch
		classList = append(classList, cls)
	}
	fmt.Println(classList)
	//c1 := Class{"python1班", 20, sch}
	//c2 := Class{"Python2班", 20, sch}
	//c1.school.band = "老男孩"
	//fmt.Println(c1.school)
	//fmt.Println(c2.school)

	for _, item := range classList {
		message := fmt.Sprintf("班级：%s，人数：%d，学校：%s", item.title, item.count, item.school.band)
		fmt.Println(message)
	}
}

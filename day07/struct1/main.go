package struct1

import "fmt"

// @program:   GoPrimary
// @file:      main.go
// @author:    bowen
// @time:      2024-11-24 21:35
// @description:

type School struct {
	band string
	city string
}

func main() {
	var s []School
	for {
		fmt.Print("请输入学校名称和城市名称")
		var band, city string
		fmt.Scanln(&band, &city)
		if band == "Q" {
			break
		}
		sch := School{
			band,
			city,
		}
		s = append(s, sch)
	}
	fmt.Println(s)

}

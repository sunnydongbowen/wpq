package main

import (
	"fmt"
)

// @program:   GoPrimary
// @file:      main.go
// @author:    bowen
// @time:      2024-11-23 20:20
// @description:

func main() {
	var v1 float32
	v1 = 3.14
	v2 := 99.9
	v3 := float64(v1) + v2
	fmt.Println(v1, v2, v3)

	v4 := 0.1
	v5 := 0.2
	res := v4 + v5
	fmt.Println(res)
	fmt.Println(0.1 + 0.2)
	fmt.Println(0.3 + 0.2)

}

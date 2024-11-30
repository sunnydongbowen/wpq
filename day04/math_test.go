package day04

import (
	"fmt"
	"math"
	"testing"
)

// @program:   GoPrimary
// @file:      math_test.go
// @author:    bowen
// @time:      2024-11-17 10:07
// @description:

func TestMath(t *testing.T) {
	fmt.Println(math.Abs(-19))     // 绝对值
	fmt.Println(math.Floor(19.99)) // 向下取整
	fmt.Println(math.Ceil(19.99))  // 向上取整
	fmt.Println(math.Round(19.49)) // 四舍五入
	fmt.Println(math.Round(19.5))  // 四舍五入
	fmt.Println(math.Round(-19.5)) //四舍五入
	fmt.Println(math.Pow(2, 3))    // 2的3次方
	fmt.Println(math.Pow10(2))     // 10的2次方
	fmt.Println(math.Max(1, 2))    // 最大值
	fmt.Println(math.Min(1, 2))    // 最小值
}

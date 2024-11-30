package main

import (
	"fmt"
	"testing"
	"time"
)

// @program:   GoPrimary
// @file:      forr_test.go
// @author:    bowen
// @time:      2024-11-15 20:37
// @description:

func TestForr(t *testing.T) {
	fmt.Println("start")
	for {
		fmt.Println("红鲤鱼与绿鲤鱼与驴")
		time.Sleep(time.Second * 1)
	}
	fmt.Println("over")
}

func TestForr1(t *testing.T) {
	fmt.Println("start")
	for 2 > 1 {
		fmt.Println("红鲤鱼与绿鲤鱼与驴")
		time.Sleep(time.Second * 1)
	}
	fmt.Println("over")
}

// 输出1~100
func TestForr2(t *testing.T) {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}

// 不输出7
func TestForr3(t *testing.T) {
	//不输出7
	for i := 1; i <= 100; i++ {
		if i == 7 {
			continue
		}
		fmt.Println(i)
	}
}

// 求1~100的和
func TestForr4(t *testing.T) {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 输出10~1
func TestForr5(t *testing.T) {
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}
}

// continue
func TestForr6(t *testing.T) {
	for {
		fmt.Println("红鲤鱼与绿鲤鱼与驴")
		continue
		// never can reach here because of continue
		fmt.Println("end")
	}
}

// continue
func TestForr7(t *testing.T) {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			if j == 3 {
				continue
			}

			fmt.Printf("%d-%d\n", i, j)
		}
	}
}

func TestForr8(t *testing.T) {
f1:
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 5; j++ {
			if j == 3 {
				//
				continue f1
			}
			fmt.Printf("%d-%d\n", i, j)
		}

	}
}

// 冒泡排序
func TestForr9(t *testing.T) {
	nums := []int{1, 5, 3, 2, 4}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
		}
		fmt.Println(nums)
	}
}

// 求斐波那契数列
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func TestForr10(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}

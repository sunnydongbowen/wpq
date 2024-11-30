package day04

import (
	"fmt"
	"testing"
)

// @program:   GoPrimary
// @file:      floatt_test.go
// @author:    bowen
// @time:      2024-11-23 20:17
// @description:

func TestFloat1(t *testing.T) {
	var v1 float32
	v1 = 3.14
	v2 := 99.9
	v3 := float64(v1) + v2
	fmt.Println(v1, v2, v3)
}

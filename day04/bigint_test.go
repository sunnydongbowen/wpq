package day04

import (
	"fmt"
	"math/big"
	"testing"
)

// @program:   GoPrimary
// @file:      bigint_test.go
// @author:    bowen
// @time:      2024-11-17 10:35
// @description:

func TestBigint(t *testing.T) {
	var v1 big.Int
	//var v2 *big.Int
	//v3 := new(big.Int)

	v1.SetInt64(1993)
	fmt.Println(v1)
	// 把我们输入的字符串以10进制的方式写入内存
	// 超大整型，超过了int64代表的最大范围。这个用的比较多
	v1.SetString("1993", 10)
	fmt.Println(v1)
}

func TestBigintAdd(t *testing.T) {
	var v1 big.Int
	v1.SetString("1993", 10)
	fmt.Println(v1.Add(&v1, big.NewInt(1)))

	var result big.Int
	result.Add(&v1, big.NewInt(1))
	fmt.Println(result)
}

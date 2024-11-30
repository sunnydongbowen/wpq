package day07

// @program:   GoPrimary
// @file:      factory_test.go
// @author:    bowen
// @time:      2024-11-26 10:43
// @description:

import (
	"fmt"
	"testing"
	"wpq/day07/util"
)

func TestNewFile(t *testing.T) {
	//强制用户去使用这种方式创建
	obj := util.NewFile(10, "test.txt")
	fmt.Println(obj)
}

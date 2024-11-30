package day07

import (
	"fmt"
	"reflect"
	"testing"
)

// @program:   GoPrimary
// @file:      structt_test.go
// @author:    bowen
// @time:      2024-11-24 21:24
// @description:

type Person struct {
	name string "姓名"
	age  int    "年龄"
}

func Test(t *testing.T) {
	p := Person{
		"bowen",
		33,
	}

	p2 := &Person{
		"bowen2",
		13,
	}

	fmt.Println(p.name, p.age)
	fmt.Println(p2.name, p2.age)

	p3 := new(Person)
	p3.name = "bowen3"
	fmt.Println(p3.name, p3.age)

	pType := reflect.TypeOf(p)
	filed1 := pType.Field(0)
	fmt.Println(filed1.Name, filed1.Tag)

	// numberField 字段数
	// 根据索引获取每个字段对象
	// 然后再根据根据字段对象获得name和tag
	for i := 0; i < pType.NumField(); i++ {
		filed := pType.Field(i)
		fmt.Println(filed.Name, filed.Tag)
	}
}

var P Person = Person{"bowen12", 19}

func doSth() *Person {
	return &P
}

func TestPer(t *testing.T) {
	data := doSth()
	P.name = "eee"
	fmt.Println(data)
	fmt.Println(P)
}

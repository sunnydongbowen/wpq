package day06

import (
	"fmt"
	"testing"
)

// @program:   GoPrimary
// @file:      map_test.go
// @author:    bowen
// @time:      2024-11-20 14:34
// @description:

func TestMap1(t *testing.T) {
	var map1 map[string]interface{} = map[string]interface{}{
		"name": "bowen",
		"age":  18,
	}
	map1["email"] = "bo@126.com"
	fmt.Println(map1)

	//map2 := new(map[string]int)
	//map2["name"] = 18
	//fmt.Println(map2)

	// 这种数据结构类似坐标系，很少用
	v1 := make(map[[2]int]float32)
	v1[[2]int{1, 1}] = 1.0
	v1[[2]int{1, 2}] = 2.0
	fmt.Println(v1)

	v2 := make(map[[2]int][3]string)
	v2[[2]int{1, 1}] = [3]string{"a", "b", "c"}
	fmt.Println(v2)

	fmt.Println(len(map1))
}

func TestMap2(t *testing.T) {
	info := make(map[int]string, 10)
	info[1] = "a"
	info[2] = "b"
	fmt.Println(len(info))
}

func TestMap3(t *testing.T) {
	//v1 := make(map[string]int)
	//v2 := make(map[string]string)
	//v3 := make(map[string]interface{})
	//v4 := make(map[string]map[string]int)
	v5 := make(map[string][]int)
	v5["aa"] = []int{1, 2, 3}
	v5["bb"] = []int{4, 5, 6}
	fmt.Println(v5)

	v6 := make(map[string]map[int]int)
	v6["a"] = map[int]int{1: 1, 2: 2}
	fmt.Println(v6)

	v7 := make(map[string][2]map[string]string)
	v7["a"] = [2]map[string]string{{"a": "a"}, {"b": "b"}}
	v7["b"] = [2]map[string]string{{"c": "c"}, {"d": "d"}}
	fmt.Println(v7)
}

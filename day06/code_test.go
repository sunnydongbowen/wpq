package day06

import (
	"fmt"
	"strconv"
	"testing"
	"unicode/utf8"
)

// @program:   GoPrimary
// @file:      code_test.go
// @author:    bowen
// @time:      2024-11-22 12:08
// @description:

func TestCode(t *testing.T) {
	name := "大灰狼"
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2)) //
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2))
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2))

	fmt.Println(name[3])
	fmt.Println(name[4])
	fmt.Println(name[5])

	fmt.Println(name[6])

	// str 转成byte
	byteSet := []byte(name)
	fmt.Println(byteSet)

	// byte 转成str
	byteList := []byte{229, 164, 167, 231, 129, 176, 231, 139, 188}
	targetStr := string(byteList)
	fmt.Println(targetStr)

	//str unicode 字符编码的集合
	tempSet := []rune(name)
	fmt.Println(tempSet)
	fmt.Println(tempSet[0], strconv.FormatInt(int64(tempSet[0]), 16))
	fmt.Println(tempSet[1], strconv.FormatInt(int64(tempSet[1]), 16))
	fmt.Println(tempSet[2], strconv.FormatInt(int64(tempSet[2]), 16))

	runeList := []rune{0x5927, 0x7070, 0x72fc}
	fmt.Println(string(runeList))

	runelen := utf8.RuneCountInString(name)
	fmt.Println(runelen, len(name))
}

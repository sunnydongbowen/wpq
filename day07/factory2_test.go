package day07

import (
	"fmt"
	"testing"
)

// @program:   GoPrimary
// @file:      factory2_test.go
// @author:    bowen
// @time:      2024-11-26 10:48
// @description:

type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	return &File{fd, name}
}

func TestNewFile1(t *testing.T) {
	f := NewFile(1, "test.txt")
	fmt.Println(f)
}

func TestNewFile2(t *testing.T) {
	f := File{
		10,
		"test.txt",
	}
	fmt.Println(f)
}

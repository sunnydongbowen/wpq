package util

// @program:   GoPrimary
// @file:      ct.go
// @author:    bowen
// @time:      2024-11-26 10:45
// @description:

type file struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *file {
	return &file{fd: fd, name: name}
}

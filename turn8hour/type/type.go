package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//var a string
	//// pair<statictype:string, value:"aceld">
	//a = "aceld"
	//
	//var allType interface{}
	//// pair<type:string, value:"aceld">
	//allType = a
	//
	//str, _ := allType.(string)
	//
	//println(str)

	// tty: pair<type: *os.File, value: "/dev/tty" 文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	// r: pair<type: , value: >
	var r io.Reader
	// r: pair<type: *os.File, value: "/dev/tty" 文件描述符>
	r = tty

	// w: pair<type: , value: >
	var w io.Writer
	// w: pair<type: *os.File, value: "/dev/tty" 文件描述符>
	w = r.(io.Writer) // 强转

	w.Write([]byte("HELLO THIS IS A TEST!!\n"))

}

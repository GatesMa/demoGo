package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[PS]")
	log.Println("这是一条很普通的日志。")

	/// ----------
	logger := log.New(os.Stdout, "<PS>", log.Lshortfile|log.Ldate|log.Ltime)

	logFile1, err := os.OpenFile("./xx1.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	logger.SetOutput(logFile1)
	logger.Println("这是自定义的logger记录的日志。")
}

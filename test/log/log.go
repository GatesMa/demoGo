package main

import "log"

func main() {
	log.SetFlags(log.Lmicroseconds | log.Ldate)
	log.Println("这是一条优雅的日志。")
	v := "优雅的"
	log.Printf("这是一个%s日志\n", v)
	//fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。
	log.Fatalln("这是一天会触发fatal的日志")
	log.Panicln("这是一个会触发panic的日志。") //执行后会自动触发一个异常
}

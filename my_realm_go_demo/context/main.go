package main

import (
	"context"
	"fmt"
	"time"
)

func contextTest(ctx context.Context) {

	cancelCtx, cancelFunc := context.WithCancel(ctx)

	go func() {
		select {
		case <-cancelCtx.Done():
			fmt.Println("cancelCtx 关闭了")
		}
	}()

	time.Sleep(time.Second * 3)
	cancelFunc()

}

func main() {
	todo := context.TODO()
	ctx, cancelFunc := context.WithCancel(todo)
	go contextTest(ctx)

	// 让 contextTest 协程运行3秒钟，然后调用取消函数
	time.Sleep(time.Second * 3)

	fmt.Println("call cancel func")
	cancelFunc()

	// 等待 3 秒钟，让 contextTest 协程优雅结束。
	time.Sleep(3 * time.Second)
	s := <-ctx.Done()
	fmt.Printf("receive %v, %T\n", s, s)

	intChan := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		println("intChan <- 1")
		intChan <- 1
	}()

	for {
		select {
		case <-intChan:
			fmt.Println("ctx done")
		default:
			fmt.Println("default")
		}
		time.Sleep(time.Second * 1)
	}

}

package main

import (
	"os"
	"os/exec"
)

func main() {
	// 删除操作系统中的冲突环境变量
	os.Unsetenv("GOPRIVATE")

	// 设置 GOPRIVATE 环境变量
	cmd := exec.Command("go", "env", "-w", "GOPRIVATE=''")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

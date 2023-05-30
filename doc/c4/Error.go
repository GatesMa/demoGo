package main

import "fmt"

type ParseError struct {
	FileName string // 文件名
	Line     int    // 行号
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.FileName, e.Line)
}

func NewParseError(fileName string, line int) error {
	return &ParseError{fileName, line}
}

func main() {

	e := NewParseError("error.go", 1)
	fmt.Println(e.Error())

	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("Filename: %s Line: %d\n", detail.FileName, detail.Line)
	default: // 其他类型的错误
		fmt.Println("other error")
	}

}

package main

import (
	"fmt"
	"os"
)

func main() {
	fd, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Failed to openfile", err)
		return
	}
	n, err := fd.WriteString("Hello world!\n")
	if err != nil {
		fmt.Println("Failed to writestring", err)
		return
	}
	fmt.Printf("write sucess %d bytes\n", n)
	//调整文件读写位置
	fd.Seek(0, os.SEEK_SET)
	//构建缓冲区
	buf := make([]byte, 20)
	n, err = fd.Read(buf)
	os.Stdout.Write(buf)
	fd.Close()
}

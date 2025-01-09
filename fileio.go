package main

import (
	"fmt"
	"os"
)

// 複製檔案
func copyFile(src, dst string) {
	_, err := os.Stat(src)
	if err != nil {
		fmt.Println("错误：无法读取文件:", src)
		return
	}

	in, err := os.Open(src)
	if err != nil {
		fmt.Println("错误：无法打开文件:", src)
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		fmt.Println("错误：无法创建文件:", dst)
		return
	}
	defer out.Close()

	buf := make([]byte, 4096)
	for {
		n, err := in.Read(buf)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("错误：读取文件时出错:", src)
			return
		}
		if n == 0 {
			break
		}
		out.Write(buf[:n])
		currentSize += int64(n)
	}

	if verbose {
		fmt.Println("复制:", src, "->", dst)
	}

	showProgress()
}

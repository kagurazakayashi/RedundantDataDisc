package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	filesDir    string
	isoName     string
	tempDir     string
	isoSizeStr  string
	verbose     bool
	isoSize     int64
	currentSize int64
)

func init() {
	flag.StringVar(&filesDir, "d", "files", "存放源文件的目录 (默认: files)")
	flag.StringVar(&isoName, "o", "secure_backup.iso", "ISO 文件名 (默认: secure_backup.iso)")
	flag.StringVar(&tempDir, "t", "iso_temp", "临时目录 (默认: iso_temp)")
	flag.StringVar(&isoSizeStr, "s", "4.7G", "ISO 目标大小 (默认: 4.7G，支持 B/K/M/G)")
	flag.BoolVar(&verbose, "v", false, "启用详细模式 (显示 mkdir 和 cp 详细信息)")
	flag.Parse()
	// 解析大小引數
	isoSize = parseSize(isoSizeStr)
}

func main() {
	// 檢查原始檔目錄
	if _, err := os.Stat(filesDir); os.IsNotExist(err) {
		fmt.Println("错误：源文件目录不存在:", filesDir)
		os.Exit(1)
	}

	// 刪除並建立臨時目錄
	os.RemoveAll(tempDir)
	os.Mkdir(tempDir, 0755)

	// 獲取原始檔案列表
	files, err := filepath.Glob(filepath.Join(filesDir, "*"))
	if err != nil || len(files) == 0 {
		fmt.Println("错误：没有找到源文件。")
		os.Exit(1)
	}

	// 複製檔案並填充空間
	fmt.Println("开始填充数据...")
	folderIndex := 1
	for currentSize < isoSize {
		folderName := filepath.Join(tempDir, fmt.Sprintf("%04d", folderIndex))
		if verbose {
			fmt.Println("创建目录:", folderName)
		}
		os.Mkdir(folderName, 0755)

		for _, file := range files {
			if currentSize >= isoSize {
				break
			}
			destFile := filepath.Join(folderName, filepath.Base(file))
			copyFile(file, destFile)
		}

		folderIndex++
		showProgress()
	}

	fmt.Println("\n填充完成，创建 ISO 文件...")

	// 生成 ISO
	createISO()

	// 刪除臨時目錄
	os.RemoveAll(tempDir)

	fmt.Println("ISO 文件已生成:", isoName)

}

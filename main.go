package main

import (
	"flag"
)

var (
	keysDir     string
	isoName     string
	tempDir     string
	isoSizeStr  string
	verbose     bool
	isoSize     int64
	currentSize int64
)

func init() {
	flag.StringVar(&keysDir, "d", "keys", "存放私钥的目录 (默认: keys)")
	flag.StringVar(&isoName, "o", "secure_backup.iso", "ISO 文件名 (默认: secure_backup.iso)")
	flag.StringVar(&tempDir, "t", "iso_temp", "临时目录 (默认: iso_temp)")
	flag.StringVar(&isoSizeStr, "s", "4.7G", "ISO 目标大小 (默认: 4.7G，支持 B/K/M/G)")
	flag.BoolVar(&verbose, "v", false, "启用详细模式 (显示 mkdir 和 cp 详细信息)")
	flag.Parse()

	// 解析大小参数
	isoSize = parseSize(isoSizeStr)
}

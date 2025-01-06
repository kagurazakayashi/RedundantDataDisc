package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseSize(sizeStr string) int64 {
	re := regexp.MustCompile(`(?i)^(\d+(\.\d+)?)([BKMG]?)$`)
	matches := re.FindStringSubmatch(sizeStr)
	if matches == nil {
		fmt.Println("错误：无效的大小格式:", sizeStr)
		os.Exit(1)
	}

	value, _ := strconv.ParseFloat(matches[1], 64)
	unit := strings.ToUpper(matches[3])

	switch unit {
	case "B":
		return int64(value)
	case "K":
		return int64(value * 1024)
	case "M":
		return int64(value * 1024 * 1024)
	case "G", "":
		return int64(value * 1024 * 1024 * 1024) // 默认 GB
	default:
		fmt.Println("错误：未知单位:", unit)
		os.Exit(1)
	}
	return 0
}

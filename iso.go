package main

import (
	"fmt"
	"os"
	"os/exec"
)

// 生成 ISO 檔案
func createISO() {
	cmd := exec.Command("genisoimage", "-o", isoName, "-J", "-R", tempDir)
	if _, err := exec.LookPath("genisoimage"); err != nil {
		cmd = exec.Command("mkisofs", "-o", isoName, "-J", "-R", tempDir)
	}

	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("错误：生成 ISO 失败:", err)
		os.Exit(1)
	}
}

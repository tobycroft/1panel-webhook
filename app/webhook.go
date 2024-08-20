package app

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
	"os"
	"os/exec"
	"path/filepath"
)

func WebhookController(c *gin.Context) {
	//http://127.0.0.1:888/hook?access_key=123&param=aaa
	access_key, ok := Input.Get("access_key", c, false)
	if !ok {
		return
	}
	re, err := ExecFile(access_key)
	if err != nil {
		RET.Fail(c, 500, re, err.Error())
	} else {
		RET.Success(c, 0, re, re)
	}
}

func ExecFile(filename string) (string, error) {
	// 获取当前程序运行的文件夹路径
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}

	// 获取当前程序运行的目录
	execDir := filepath.Dir(execPath)

	// 构建 exec 文件夹的路径
	execDirPath := filepath.Join(execDir, "exec")

	// 构建完整的.sh文件路径
	filePath := filepath.Join(execDirPath, filename)
	fmt.Println("exec:", filePath)

	// 检查文件是否存在
	if _, err := os.Stat(filePath + ".sh"); os.IsNotExist(err) {
		return "未找到文件", errors.New("未找到文件")
	}

	// 文件存在，执行.sh脚本
	out, err := exec.Command("/bin/bash", filePath).Output()
	if err != nil {
		return "", err
	}

	// 返回执行结果
	return string(out), nil
}

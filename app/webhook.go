package app

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
	"net/url"
	"os"
	"os/exec"
	"strings"
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
	filename = url.QueryEscape(filename)
	filename = strings.ReplaceAll(filename, ".", "")
	filename = strings.ReplaceAll(filename, "%", "")
	filename = strings.ReplaceAll(filename, "&", "")
	fmt.Println("./exec/" + filename + ".sh")
	if _, err := os.Stat("./exec/" + filename + ".sh"); os.IsNotExist(err) {
		return "未找到文件", errors.New("未找到文件")
	}
	// 文件存在，执行.sh脚本
	out, err := exec.Command("/bin/bash", "./exec/"+filename+".sh").Output()
	if err != nil {
		return "", err
	}

	// 返回执行结果
	return string(out), nil
}

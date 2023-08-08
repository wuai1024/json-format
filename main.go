package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	indentSpaces := 4

	args := os.Args
	if len(args) > 1 {
		indentSpaces, _ = strconv.Atoi(args[1])
	}

	// 获取剪切板内容
	clipboardContent, err := clipboard.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 判断是否为合法的 JSON 格式
	var jsonData interface{}
	err = json.Unmarshal([]byte(clipboardContent), &jsonData)
	if err != nil {
		fmt.Println("剪切板内容不是合法的 JSON 格式")
	} else {
		// 对 JSON 进行格式化（使用缩进空格数）
		indent := strings.Repeat(" ", indentSpaces)
		formattedJSON, err := json.MarshalIndent(jsonData, "", indent)
		if err != nil {
			log.Fatal(err)
		}

		// 将格式化后的 JSON 放回剪切板
		err = clipboard.WriteAll(string(formattedJSON))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("已将格式化后的 JSON 设置到剪切板（缩进：%d个空格）", indentSpaces)
	}
}

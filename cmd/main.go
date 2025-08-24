package main

import (
	"fmt"
	"github.com/KOU050223/flasgo/internal/filemaker"
	"github.com/KOU050223/flasgo/internal/help"
	"os"
)

func main() {
	args := os.Args[1:] // 引数の受け取り

	if len(args) == 0 {
		fmt.Println("コマンドを指定してください")
		help.Help()
		return
	}

	switch args[0] {
	case "create":
		filemaker.Generator()
	case "help":
		help.Help()
	default:
		fmt.Printf("不明なコマンド: %s\n", args[0])
		help.Help()
	}
}

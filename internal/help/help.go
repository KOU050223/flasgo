package help

import (
	"fmt"
	"github.com/KOU050223/flasgo/types"
)

func Help() {
	fmt.Println("ヘルプ一覧を表示します")
	commands := []types.Command{
		types.Command{"create", "flaskの標準的なフォルダ・ファイルを生成します"},
		types.Command{"help", "コマンド一覧を表示します"},
	}
	for i, command := range commands {
		fmt.Printf("%d. %s: %s\n", i+1, command.Name, command.Description)
	}
}

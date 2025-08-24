package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// プロンプト用の構造体
type Option struct {
	Label string
	Value string
}

// テキスト入力を求める
func PromptText(question string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	
	if defaultValue != "" {
		fmt.Printf("? %s (%s): ", question, defaultValue)
	} else {
		fmt.Printf("? %s: ", question)
	}
	
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	if input == "" && defaultValue != "" {
		return defaultValue
	}
	
	return input
}

// 選択肢から選ぶ
func PromptSelect(question string, options []Option) string {
	fmt.Printf("? %s\n", question)
	
	for i, option := range options {
		if i == 0 {
			fmt.Printf("  › %d) %s\n", i+1, option.Label)
		} else {
			fmt.Printf("    %d) %s\n", i+1, option.Label)
		}
	}
	
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("選択してください (1-" + strconv.Itoa(len(options)) + "): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if num, err := strconv.Atoi(input); err == nil {
			if num >= 1 && num <= len(options) {
				return options[num-1].Value
			}
		}
		
		fmt.Printf("無効な選択です。1-%d の番号を入力してください。\n", len(options))
	}
}

// 複数選択（チェックボックス形式）
func PromptMultiSelect(question string, options []Option) []string {
	fmt.Printf("? %s (スペース区切りで複数選択可。例: 1 3 4)\n", question)
	
	for i, option := range options {
		fmt.Printf("    %d) %s\n", i+1, option.Label)
	}
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("選択してください: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	if input == "" {
		return []string{}
	}
	
	var selected []string
	numbers := strings.Fields(input)
	
	for _, numStr := range numbers {
		if num, err := strconv.Atoi(numStr); err == nil {
			if num >= 1 && num <= len(options) {
				selected = append(selected, options[num-1].Value)
			}
		}
	}
	
	return selected
}

// 確認プロンプト
func PromptConfirm(question string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("? %s (y/N): ", question)
	
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))
	
	return input == "y" || input == "yes"
}
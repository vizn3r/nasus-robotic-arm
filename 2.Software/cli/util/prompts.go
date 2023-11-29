package util

import (
	"fmt"
	"strings"
)

func UserBoolPrompt(prompt string, options [][]string) bool {
	opts := " ["
	for i, opt := range options {
		for j, sub := range opt {
			opts += sub
			if j != len(opt) - 1 {
				opts += ", "
			}
		}
		if i != len(options) - 1 {
			opts += " / "
		}
	}
	opts += "]?: "
	fmt.Print(prompt + opts)
	var ans string
	fmt.Scanf("%s", ans)
	for _, opt := range options {
		for _, sub := range opt {
			if strings.EqualFold(sub, ans) {
				return true
			}
		}
	}
	return false
}

func UserIndexPrompt(prompt string, options []string) int {
	opts := ""
	for i, opt := range options {
		opts += fmt.Sprintf("%v", i) + ". - " + opt
		if i != len(options) - 1 {
			opts += "\n"
		}
	}
	opts += "\n"
	fmt.Print(opts + prompt + "?: ")
	var ans int
	fmt.Scanf("%d", ans)
	for _, opt := range options {
		if options[ans] == opt {
			return ans
		}
	}
	fmt.Println("Please choose one of the options")
	return -1
}

func UserStringPrompt(prompt string, options []string) string {
	opts := ""
	for i, opt := range options {
		opts += fmt.Sprintf("%v", i) + ". - " + opt
		if i != len(options) - 1 {
			opts += "\n"
		}
	}
	opts += "\n"
	fmt.Print(opts + prompt + "?: ")
	var ans int
	fmt.Scanf("%d", ans)
	for _, opt := range options {
		if options[ans] == opt {
			return opt
		}
	}
	fmt.Println("Please choose one of the options")
	return ""
}
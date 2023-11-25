package util

import (
	"fmt"
	"strings"
)

func UserPrompt(prompt string, options [][]string) bool {
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
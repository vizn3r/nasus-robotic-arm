package util

import "strings"

func StringArrayHas(arr []string, item string) bool {
	for _, i := range arr {
		if strings.Compare(i, item) == 0 {
			return true
		}
	}
	return false
}
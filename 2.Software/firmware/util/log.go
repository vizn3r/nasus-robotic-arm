package util

import "os"

func PrintExt(a ...string) {
	for _, s := range a {
		print(s + " ")
	}
	print("\n")
	os.Exit(1)
}
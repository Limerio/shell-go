package utils

import (
	"fmt"
	"os"
	"strings"
)

func Exit() {
	fmt.Println("Goodbye")
	os.Exit(0)
}

func Escape() {
	fmt.Println()
}

func FormatPathInShell(path string) string {

	homeEnvVar := os.Getenv("HOME")

	if strings.HasPrefix(path, homeEnvVar) {
		return fmt.Sprintf("~%s", strings.Split(path, homeEnvVar)[1])
	}

	return path
}

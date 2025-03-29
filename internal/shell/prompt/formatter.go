package prompt

import (
	"fmt"
	"os"
	"strings"
)

func FormatPathInShell(path string) string {
	homeEnvVar := os.Getenv("HOME")

	if strings.HasPrefix(path, homeEnvVar) {
		return fmt.Sprintf("~%s", strings.Split(path, homeEnvVar)[1])
	}

	return path
}

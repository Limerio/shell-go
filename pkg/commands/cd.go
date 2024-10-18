package commands

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Cd(args []string) {
	if len(args) == 0 {
		fmt.Println("The first argument needs to a path.")
		return
	}

	dir := args[0]

	if strings.TrimSpace(dir) == "~" {
		dir = os.Getenv("HOME")
	}

	isAbsolute := dir[0] == '/'
	if !isAbsolute {
		dir = filepath.Join(os.Getenv("PWD"), dir)
	}

	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%s: No such file or directory\n", dir)
		return
	}

	os.Setenv("PWD", dir)
}

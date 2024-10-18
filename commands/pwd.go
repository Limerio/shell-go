package commands

import (
	"fmt"
	"os"
)

func Pwd(args []string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dir)
}

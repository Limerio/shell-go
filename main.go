package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Limerio/shell-go/commands"
	"github.com/Limerio/shell-go/utils"
)

func main() {
	go utils.SignalC()

	fmt.Println("Hello in the best shell (jk)")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s > ", utils.FormatPathInShell(os.Getenv("PWD")))
		scanner.Scan()
		prompt := strings.Fields(scanner.Text())
		cmd := prompt[0]
		args := prompt[1:]

		cmdMap := map[string]func([]string){
			"cat":      commands.Cat,
			"cd":       commands.Cd,
			"clear":    commands.Clear,
			"echo":     commands.Echo,
			"exit":     commands.Exit,
			"hostname": commands.Hostname,
			"mkdir":    commands.Mkdir,
			"nslookup": commands.NsLookup,
			"pwd":      commands.Pwd,
			"ls":       commands.Ls,
			"rm":       commands.Rm,
			"rmdir":    commands.Rmdir,
			"touch":    commands.Touch,
		}

		command, ok := cmdMap[cmd]
		if !ok {
			defaultFunc()
			continue
		}
		command(args)
	}
}

func defaultFunc() {
	fmt.Println("Unknown command")
}

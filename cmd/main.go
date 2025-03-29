package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Limerio/shell-go/internal/shell/prompt"
	"github.com/Limerio/shell-go/internal/shell/signals"
	fsCommands "github.com/Limerio/shell-go/pkg/builtins/fs"
	ioCommands "github.com/Limerio/shell-go/pkg/builtins/io"
	networkCommands "github.com/Limerio/shell-go/pkg/builtins/network"
	systemCommands "github.com/Limerio/shell-go/pkg/builtins/system"
)

func main() {
	go signals.HandleInterrupt()

	fmt.Println("Hello in the best shell (jk)")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s > ", prompt.FormatPathInShell(os.Getenv("PWD")))
		scanner.Scan()
		prompt := strings.Fields(scanner.Text())
		cmd := prompt[0]
		args := prompt[1:]

		cmdMap := map[string]func([]string){
			"cat":      fsCommands.Cat,
			"cd":       fsCommands.Cd,
			"clear":    ioCommands.Clear,
			"echo":     ioCommands.Echo,
			"exit":     systemCommands.Exit,
			"hostname": systemCommands.Hostname,
			"mkdir":    fsCommands.Mkdir,
			"nslookup": networkCommands.NsLookup,
			"pwd":      fsCommands.Pwd,
			"ls":       fsCommands.Ls,
			"rm":       fsCommands.Rm,
			"rmdir":    fsCommands.Rmdir,
			"touch":    fsCommands.Touch,
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

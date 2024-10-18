package commands

import "fmt"

func Echo(args []string) {
	fmt.Println(args[0])
}

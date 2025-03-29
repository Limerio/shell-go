package exit

import (
	"fmt"
	"os"
)

func Exit() {
	fmt.Println("Goodbye")
	os.Exit(0)
}

func Escape() {
	fmt.Println()
}

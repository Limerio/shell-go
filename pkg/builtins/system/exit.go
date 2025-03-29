package commands

import (
	"github.com/Limerio/shell-go/internal/shell/exit"
)

func Exit(args []string) {
	exit.Exit()
}

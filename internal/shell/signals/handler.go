package signals

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Limerio/shell-go/internal/shell/exit"
)

func HandleInterrupt() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	exit.Escape()
	exit.Exit()
}

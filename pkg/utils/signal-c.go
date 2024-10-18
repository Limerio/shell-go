package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func SignalC() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	Escape()
	Exit()
}

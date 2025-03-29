package commands

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Limerio/shell-go/internal/shell/exit"
)

func NsLookup(args []string) {
	if len(args) == 0 {
		fmt.Println("The DNS name needs to provide")
		return
	}

	ns, err := net.LookupHost(args[0])
	if err != nil {
		fmt.Println(err)
	}

	data, _ := os.ReadFile("/etc/resolv.conf")
	server := strings.TrimSpace(strings.Split(string(data), "nameserver")[1])

	fmt.Printf("Server: \t%s", server)
	exit.Escape()
	fmt.Printf("Address: \t%s#53", server)
	exit.Escape()
	exit.Escape()

	fmt.Println("Non-authoritative answer:")
	for _, n := range ns {
		if !strings.ContainsAny(n, ":") {
			fmt.Printf("Name:   %s", args[0])
			exit.Escape()
			fmt.Printf("Address: %s", n)
			exit.Escape()
		}
	}
	exit.Escape()
}

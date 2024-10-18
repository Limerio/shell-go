package commands

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Limerio/shell-go/utils"
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
	utils.Escape()
	fmt.Printf("Address: \t%s#53", server)
	utils.Escape()
	utils.Escape()

	fmt.Println("Non-authoritative answer:")
	for _, n := range ns {
		if !strings.ContainsAny(n, ":") {
			fmt.Printf("Name:   %s", args[0])
			utils.Escape()
			fmt.Printf("Address: %s", n)
			utils.Escape()
		}
	}
	utils.Escape()
}

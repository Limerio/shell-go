package system

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Hostname(args []string) {
	hostnameCmd := &cobra.Command{
		Use:   "hostname",
		Short: "Print name of current host system",
		Run: func(cmd *cobra.Command, args []string) {
			hostname, err := os.Hostname()
			if err != nil {
				fmt.Println("Something went wrong")
				return
			}
			fmt.Printf("%s", hostname)
		},
	}

	if err := hostnameCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

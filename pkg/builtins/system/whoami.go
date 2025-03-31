package system

import (
	"fmt"
	"os/user"

	"github.com/spf13/cobra"
)

func Whoami(args []string) {
	whoamiCmd := &cobra.Command{
		Use: "whoami",
		Run: func(cmd *cobra.Command, args []string) {
			currentUser, err := user.Current()
			if err != nil {
				fmt.Println("Unable to get the information about the current user.")

				return
			}

			fmt.Println(currentUser.Username)
		},
	}

	if err := whoamiCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

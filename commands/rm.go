package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Rm(args []string) {
	rmCmd := &cobra.Command{
		Use: "rm",
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]
			if err := os.Remove(file); err != nil {
				fmt.Println(err)
				return
			}
		},
	}

	rmCmd.SetArgs(args)

	rmCmd.Flags().BoolP("R", "r", false, "Attempt to remove the file hierarchy rooted in each file argument.")

	if err := rmCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

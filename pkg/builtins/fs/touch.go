package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Touch(args []string) {
	touchCmd := &cobra.Command{
		Use: "touch",
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]
			if err := os.WriteFile(file, []byte{}, 0644); err != nil {
				fmt.Println(err)
				return
			}
		},
	}

	touchCmd.SetArgs(args)

	if err := touchCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

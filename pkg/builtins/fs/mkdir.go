package fs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Mkdir(args []string) {
	rmdirCmd := &cobra.Command{
		Use: "mkdir",
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]
			if err := os.Mkdir(file, 06); err != nil {
				fmt.Println(err)
				return
			}
		},
	}

	rmdirCmd.SetArgs(args)

	if err := rmdirCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

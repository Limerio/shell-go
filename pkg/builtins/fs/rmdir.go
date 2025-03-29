package fs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Rmdir(args []string) {
	rmdirCmd := &cobra.Command{
		Use: "rmdir",
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]
			if err := os.Remove(file); err != nil {
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

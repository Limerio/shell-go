package fs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func LinkRun(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("You need to have 2 paths as arguments.")

		return
	}

	if err := os.Symlink(args[0], args[1]); err != nil {
		fmt.Println(err)
	}
}

func Link(args []string) {
	linkCmd := &cobra.Command{
		Use:     "link",
		Aliases: []string{"ln"},
		Run:     LinkRun,
	}

	linkCmd.SetArgs(args)

	if err := linkCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

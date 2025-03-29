package commands

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/Limerio/shell-go/internal/shell/exit"
	"github.com/spf13/cobra"
)

func render(i int, e fs.DirEntry) {
	if i == 0 {
		fmt.Print(e.Name())
	} else {
		fmt.Printf("\t %s", e.Name())
	}
}

func run(cmd *cobra.Command, args []string) {
	allValue, err := cmd.Flags().GetBool("all")
	if err != nil {
		fmt.Println(err)
		return
	}

	dirsValue, err := cmd.Flags().GetBool("dirs")
	if err != nil {
		fmt.Println(err)
		return
	}

	filesValue, err := cmd.Flags().GetBool("files")
	if err != nil {
		fmt.Println(err)
		return
	}

	entries, err := os.ReadDir(os.Getenv("PWD"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, e := range entries {
		if dirsValue && e.IsDir() {
			render(i, e)
		} else if filesValue && !e.IsDir() {
			render(i, e)
		} else if allValue {
			render(i, e)
		} else if !dirsValue && !filesValue && !allValue && !strings.HasPrefix(e.Name(), ".") {
			render(i, e)
		}
	}
}

func Ls(args []string) {
	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "List all the files and directories (by default you can see the files which starts with a \".\")",
		Run:   run,
	}

	lsCmd.SetArgs(args)

	lsCmd.Flags().BoolP("all", "a", false, "View all files and directories")
	lsCmd.Flags().BoolP("dirs", "d", false, "View only all directories")
	lsCmd.Flags().BoolP("files", "f", false, "View only all files")

	if err := lsCmd.Execute(); err != nil {
		fmt.Println(err)
	}
	exit.Escape()
}

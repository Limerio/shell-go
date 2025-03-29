package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/Limerio/shell-go/internal/shell/exit"
	"github.com/spf13/cobra"
)

func run_cat(cmd *cobra.Command, args []string) {
	content, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println("Something went wrong when reading the file")
		return
	}

	splittedString := strings.Split(string(content), "\n")

	nValue, err := cmd.Flags().GetBool("number")
	if err != nil {
		fmt.Println(err)
		return
	}

	if nValue {
		for i, l := range splittedString {
			fmt.Printf("%d %s", i+1, l)
			exit.Escape()
		}
		return
	}

	eValue, err := cmd.Flags().GetBool("e")
	if err != nil {
		fmt.Println(err)
		return
	}

	if eValue {
		for i, l := range splittedString {

			if len(splittedString)-1 != i {
				fmt.Printf("%s$", l)
				exit.Escape()
				continue
			}
			fmt.Println(l)
		}
		return
	}

	fmt.Println(strings.Join(splittedString, "\n"))
}

func Cat(args []string) {
	catCmd := &cobra.Command{
		Use:   "cat",
		Long:  "The cat utility reads files sequentially, writing them to the standard output.",
		Short: "The cat utility reads files sequentially, writing them to the standard output.",
		Run:   run_cat,
	}

	catCmd.SetArgs(args)

	catCmd.Flags().BoolP("number", "n", false, "Number the output lines, starting at 1.")
	catCmd.Flags().BoolP("e", "e", false, "Display non-printing characters, and display a dollar sign ('$') at the end of each line.")

	if err := catCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

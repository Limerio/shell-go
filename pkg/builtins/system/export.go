package system

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func Export(args []string) {
	exportCmd := &cobra.Command{
		Use:   "export",
		Short: "Export a variable",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Usage: export <key>=<value>")
				return
			}

			parts := strings.Split(args[0], "=")
			os.Setenv(parts[0], parts[1])
		},
	}

	exportCmd.SetArgs(args)

	if err := exportCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

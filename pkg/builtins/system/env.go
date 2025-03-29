package system

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Env(args []string) {
	envCmd := &cobra.Command{
		Use:   "env",
		Short: "Print all environment variables",
		Run: func(cmd *cobra.Command, args []string) {
			for _, env := range os.Environ() {
				fmt.Println(env)
			}
		},
	}

	if err := envCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

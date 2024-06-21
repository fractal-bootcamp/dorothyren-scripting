package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "welcome",
	Short: "this is saying hello",
	Long:  `this is say hello for now`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("welcome")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

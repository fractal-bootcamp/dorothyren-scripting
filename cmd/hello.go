package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "printing hello",
	Long:  `printing hello`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("goodbye")
	},
}

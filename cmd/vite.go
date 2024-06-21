package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createVite)
}

var createVite = &cobra.Command{
	Use:   "vite",
	Short: "this will create a react-vite app",
	Long:  `this will create a react-vite app with typescript`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {

		// JS version
		// const args = ["test", "--template", "typescript"]
		// installArgs = ["create", "vite", ...args] // ["create", "vite", "test", "--template", "typescript"]

		installArgs := append([]string{"create", "vite"}, args...)
		installArgsWithTypescriptFlag := append(installArgs, "--template", "react-ts")
		installVite := exec.Command("bun", installArgsWithTypescriptFlag...)

		// binding the appropriate in/out
		installVite.Stdin = os.Stdin
		installVite.Stdout = os.Stdout

		err := installVite.Run()

		if err != nil {
			log.Println(err)
		}
	},
}

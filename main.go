package main

import (
	"dorothy/hello/cmd"
	"os"
	"os/exec"
)

func main() {
	cmd.Execute()

}

func _main() {
	installVite := exec.Command("bun", "create", "vite")

	installVite.Stdin = os.Stdin
	installVite.Stdout = os.Stdout

	installVite.Run()
}

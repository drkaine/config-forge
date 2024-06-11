package main

import (
	"config-forge/configs"
	"config-forge/runner"
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "-c":
		runner.CLI()
	case "--cli":
		runner.CLI()
	case "-a":
		runner.API()
	case "--api":
		runner.API()
	default:
		runner.Prompt()
	}

	fmt.Println(configs.GOODBY_RETURN)
}

package main

import (
	"config-forge/configs"
	"config-forge/runner"
	"fmt"
	"os"
)

func main() {

	switch len(os.Args) {
	case 1:
		runner.CLI()
	default:
		runner.Prompt()
	}

	fmt.Println(configs.GOODBY_RETURN)
}

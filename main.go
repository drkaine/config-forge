package main

import (
	"config-forge/configs"
	"config-forge/configurator"
	"config-forge/utils"
	"fmt"
	"os"
	"strings"
)

func main() {

	switch len(os.Args) {
	case 1:
		executeScript()
	default:
		executeCLI()
	}
}

func executeCLI() {
	if utils.CheckArguments(os.Args) {
		informations := map[string]string{
			"name":         os.Args[2],
			"serverName":   os.Args[3],
			"documentRoot": os.Args[4],
		}
		configu := configurator.ConfigBuilder(os.Args[1], informations)

		er := utils.EditFile(configu)

		if er != nil {
			fmt.Println(er)
		}

		fmt.Println("End")
	}
	fmt.Println(configs.WRONG_CHOICE_RETURN)
}

func executeScript() {
	fmt.Println(configs.HELLO_RETURN)
	fmt.Println(strings.Join(configs.ConfigAccepted, " "))
	toolName := utils.ListeningResponse(os.Stdin)
	analyseChoice := utils.AnalyseChoice(toolName)
	fmt.Println(analyseChoice)

	if analyseChoice == configs.WRONG_CHOICE_RETURN {
		executeScript()
	}

	fmt.Println(configs.ASK_NAME_FILE)

	name := utils.ListeningResponse(os.Stdin)

	fmt.Println(configs.ASK_NAME_SERVER)

	serverName := utils.ListeningResponse(os.Stdin)

	fmt.Println(configs.ASK_DOCUMENT_ROOT_PATH)

	documentRoot := utils.ListeningResponse(os.Stdin)

	informations := map[string]string{
		"name":         name,
		"serverName":   serverName,
		"documentRoot": documentRoot,
	}
	configu := configurator.ConfigBuilder(toolName, informations)

	er := utils.EditFile(configu)

	if er != nil {
		fmt.Println(er)
	}

	fmt.Println(configs.ASK_CONTINU)

	other := utils.ListeningResponse(os.Stdin)

	if other == "y" || other == "yes" {
		executeScript()
	}

	fmt.Println(configs.GOODBY_RETURN)
}

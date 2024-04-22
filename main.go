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
	Execute()
}

func Execute() {
	fmt.Println(configs.HELLO_RETURN)
	fmt.Println(strings.Join(configs.ConfigAccepted, " "))
	toolName := utils.ListeningResponse(os.Stdin)
	analyseChoice := utils.AnalyseChoice(toolName)
	fmt.Println(analyseChoice)

	if analyseChoice == configs.WRONG_CHOICE_RETURN {
		Execute()
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

	configu := configurator.InstanciationConfig(toolName, informations)

	configu.CustomiseConfigContent()
	fmt.Println(configu.Name)

	er := utils.EditFile(configu)
	fmt.Println(er)
}

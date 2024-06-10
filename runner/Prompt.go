package runner

import (
	"config-forge/configs"
	"config-forge/configurator"
	"config-forge/utils"
	"fmt"
	"os"
	"strings"
)

func Prompt() {
	fmt.Println(configs.HELLO_RETURN)
	fmt.Println(strings.Join(configs.ConfigAccepted, " "))
	toolName := utils.ListeningResponse(os.Stdin)
	analyseChoice := utils.AnalyseChoice(toolName)
	fmt.Println(analyseChoice)

	if analyseChoice == configs.WRONG_CHOICE_RETURN {
		Prompt()
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
		Prompt()
	}
}

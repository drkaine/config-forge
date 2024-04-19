package main

import (
	"fmt"
	"os"
	"strings"
)

func Execute() {
	fmt.Println(HELLO_RETURN)
	fmt.Println(strings.Join(configAccepted, " "))
	answer := ListeningResponse(os.Stdin)
	analyseResponse := AnalyseResponse(answer)
	fmt.Println(analyseResponse)

	if analyseResponse == WRONG_CHOICE_RETURN {
		Execute()
	}

	fmt.Println(ASK_NAME_FILE)

	name := ListeningResponse(os.Stdin)

	fmt.Println(ASK_NAME_SERVER)

	serverName := ListeningResponse(os.Stdin)

	fmt.Println(ASK_DOCUMENT_ROOT_PATH)

	documentRoot := ListeningResponse(os.Stdin)

	config := InstanciationConfig(name, serverName, documentRoot)

	config.ImplementConfigContent()
	fmt.Println(config.name)

	er := EditFile(config)
	fmt.Println(er)
}

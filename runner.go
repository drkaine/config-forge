package main

import (
	"fmt"
	"os"
	"strings"
)

func Runner() {
	fmt.Println(Hello)
	fmt.Println(strings.Join(configAccepted, " "))
	answer := ListeningResponse(os.Stdin)
	analyseResponse := AnalyseResponse(answer)
	fmt.Println(analyseResponse)

	if analyseResponse == WrongChoice {
		Runner()
	}

	fmt.Println(AskNameFile)

	name := ListeningResponse(os.Stdin)

	fmt.Println(AskServerName)

	serverName := ListeningResponse(os.Stdin)

	fmt.Println(AskDocumentRoot)

	documentRoot := ListeningResponse(os.Stdin)

	config := InstanciationConfig(name, serverName, documentRoot)

	config.ImplementConfigContent()
	fmt.Println(config.name)

	er := EditFile(config)
	fmt.Println(er)
}

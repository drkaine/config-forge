package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	AcceptedChoice       = "apache\n ngnix"
	Hello                = "Hello, choose your tools at configure : "
	PrepareConfiguration = "Now go to prepare the configuration"
	WrongChoice          = "Uncorrect choice !"
	AskNameFile          = "How would you name the configuration file ? :"
)

func main() {
	Runner()
}

func Runner() {
	fmt.Println(Hello)
	fmt.Println(AcceptedChoice)
	answer := ListeningResponse(os.Stdin)
	analyseResponse := AnalyseResponse(answer)
	fmt.Println(analyseResponse)

	if analyseResponse == WrongChoice {
		Runner()
	}

	fmt.Println(AskNameFile)

	name := ListeningResponse(os.Stdin)

	configurator := InstanciationConfigurator(name)

	fmt.Println(configurator.name)

}

func ListeningResponse(reader io.Reader) string {
	bufReader := bufio.NewReader(reader)
	response, _ := bufReader.ReadString('\n')
	return response
}

func AnalyseResponse(choice string) string {
	if InArray(choice, strings.Split(AcceptedChoice, "\n")) {
		return PrepareConfiguration
	}
	return WrongChoice
}

func InArray(search string, target []string) bool {
	for _, value := range target {
		if value == search {
			return true
		}
	}
	return false
}

type apache struct {
	name string
}

func InstanciationConfigurator(name string) apache {
	configurator := apache{
		name: name,
	}

	return configurator
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	Hello                = "Hello, choose your tools at configure : "
	PrepareConfiguration = "Now go to prepare the configuration"
	WrongChoice          = "Uncorrect choice !"
	AskNameFile          = "How would you name the configuration file ? :"
)

var choiceAccepted = []string{
	"apache\n",
	"ngnix\n",
}

func main() {
	Runner()
}

func Runner() {
	fmt.Println(Hello)
	fmt.Println(PresentChoices())
	answer := ListeningResponse(os.Stdin)
	analyseResponse := AnalyseResponse(answer)
	fmt.Println(analyseResponse)

	if analyseResponse == WrongChoice {
		Runner()
	}

	fmt.Println(AskNameFile)

	config := apache{
		name: answer,
	}

	fmt.Println(config.name)
}

func PresentChoices() string {
	return strings.Join(choiceAccepted, "")
}

func ListeningResponse(reader io.Reader) string {
	bufReader := bufio.NewReader(reader)
	response, _ := bufReader.ReadString('\n')
	return response
}

func AnalyseResponse(choice string) string {
	if InArray(choice, choiceAccepted) {
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

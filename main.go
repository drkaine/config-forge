package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var choiceAccepted = []string{
	"apache\n",
	"ngnix\n",
}

func main() {
	Runner()
}

func Runner() {
	fmt.Println(PresentGreeting())
	fmt.Println(PresentChoices())
	answer := ListeningResponse(os.Stdin)
	analyseResponse := AnalyseResponse(answer)
	fmt.Println(analyseResponse)
	if analyseResponse == "Uncorrect choice !" {
		Runner()
	}

	config := apache{
		name: answer,
	}

	fmt.Println(config.name)
}

func PresentGreeting() string {
	return "Hello, choose your tools : \n"
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
		return "Now go to prepare the configuration"
	}
	return "Uncorrect choice !"
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

package main

import "strings"

var choiceAccepted = []string{
	"1 apache",
	"2 ngnix",
}

func main() {

}

func Greeting() string {
	return "Hello, choose your tools : \n"
}

func GetChoices() string {
	return strings.Join(choiceAccepted, " ")
}

func AnalyseResponse(choice string) string {
	if InArray(choice, choiceAccepted) {
		return "Now go to prepare the configuration"
	}
	return "Wrong choice, choose 1 apache 2 ngnix"
}

func InArray(search string, target []string) bool {
	for _, value := range target {
		if value == search {
			return true
		}
	}
	return false
}

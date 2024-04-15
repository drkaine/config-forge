package main

import "strings"

func main() {
}

func Choices() string {
	choiceAccepted := []string{
		"1 apache",
		"2 ngnix",
	}

	return strings.Join(choiceAccepted, " ")
}

func ChoiceResponded(choice string) string {
	choiceAccepted := []string{
		"1 apache",
		"2 ngnix",
	}

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

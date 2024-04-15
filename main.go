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
	return "Wrong choice, choose 1 apache 2 ngnix"
}

package main

import (
	"bytes"
	"strings"
	"testing"
)

var configAccepted = []string{
	"1 apache\n",
	"2 ngnix\n",
}

func TestPresentGreeting(t *testing.T) {
	mainf := PresentGreeting()
	want := "Hello, choose your tools : \n"

	if mainf != want {
		t.Errorf("The greetings don't reach")
	}
}

func TestPresentChoices(t *testing.T) {
	mainf := PresentChoices()
	want := strings.Join(configAccepted, "")

	if mainf != want {
		t.Errorf("The choices don't correspond")
	}
}

func TestWrongAnalyseResponse(t *testing.T) {
	choose := "0"
	mainf := AnalyseResponse(choose)
	want := "Wrong choice, choose : " + strings.Join(configAccepted, "")

	if mainf != want {
		t.Errorf("Good choice return")
	}
}

func TestGoodAnalyseResponse(t *testing.T) {
	choose := "1 apache\n"
	mainf := AnalyseResponse(choose)
	want := "Now go to prepare the configuration"

	if mainf != want {
		t.Errorf("Wrong choice return")
	}
}

func TestListeningResponse(t *testing.T) {
	input := "1 apache\n"
	reader := bytes.NewReader([]byte(input))

	response := ListeningResponse(reader)

	if response != input {
		t.Errorf("The response %q isn't the same of the input %q", input, response)
	}
}

func TestInArray(t *testing.T) {
	want := "1 apache\n"

	if !InArray(want, choiceAccepted) {
		t.Errorf("Not in array")
	}
}

func TestNotInArray(t *testing.T) {
	want := "0 apache"

	if InArray(want, choiceAccepted) {
		t.Errorf("In array")
	}
}

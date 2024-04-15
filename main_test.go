package main

import (
	"strings"
	"testing"
)

var configAccepted = []string{
	"1 apache",
	"2 ngnix",
}

func TestGreeting(t *testing.T) {
	mainf := PresentGreeting()
	want := "Hello, choose your tools : \n"

	if mainf != want {
		t.Errorf("The greetings don't reach")
	}
}

func TestChoices(t *testing.T) {
	mainf := PresentChoices()
	want := strings.Join(configAccepted, " ")

	if mainf != want {
		t.Errorf("The choices don't correspond")
	}
}

func TestWrongAnalyseResponse(t *testing.T) {
	choose := "0"
	mainf := AnalyseResponse(choose)
	want := "Wrong choice, choose 1 apache 2 ngnix"

	if mainf != want {
		t.Errorf("Good choice return")
	}
}

func TestGoodAnalyseResponse(t *testing.T) {
	choose := "1 apache"
	mainf := AnalyseResponse(choose)
	want := "Now go to prepare the configuration"

	if mainf != want {
		t.Errorf("Good choice return")
	}
}

package main

import (
	"os"
	"strings"
	"testing"
)

var configAccepted []string

func TestMain(m *testing.M) {
	configAccepted = []string{
		"1 apache",
		"2 ngnix",
	}

	code := m.Run()

	os.Exit(code)
}

func TestChoices(t *testing.T) {
	mainf := Choices()
	want := strings.Join(configAccepted, " ")

	if mainf != want {
		t.Errorf("The choices don't correspond")
	}
}

func TestWrongChoiceResponded(t *testing.T) {
	choose := "0"
	mainf := ChoiceResponded(choose)
	want := "Wrong choice, choose 1 apache 2 ngnix"

	if mainf != want {
		t.Errorf("Good choice return")
	}
}

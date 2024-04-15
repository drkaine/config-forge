package main

import (
	"testing"
)

func TestChoices(t *testing.T) {
	mainf := Choices()
	want := "1 apache 2 ngnix"

	if mainf != want {
		t.Errorf("The choices don't correspond")
	}
}

func TestWrongChoice(t *testing.T) {
	choose := "0"
	mainf := ChoiceResponded(choose)
	want := "Wrong choice, choose 1 apache 2 ngnix"

	if mainf != want {
		t.Errorf("Good choice return")
	}
}

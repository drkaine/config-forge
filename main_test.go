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

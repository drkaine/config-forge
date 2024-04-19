package main

import (
	"testing"
)

func TestInArray(t *testing.T) {
	if !InArray(ApacheInput, configAccepted) {
		t.Errorf("Not in array")
	}
}

func TestNotInArray(t *testing.T) {
	if InArray(BadInput, configAccepted) {
		t.Errorf("In array")
	}
}

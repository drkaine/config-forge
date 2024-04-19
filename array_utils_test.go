package main

import (
	"testing"
)

func TestInArray(t *testing.T) {
	if !InArray(APACHE_INPUT, configAccepted) {
		t.Errorf("Not in array")
	}
}

func TestNotInArray(t *testing.T) {
	if InArray(BAD_INPUT, configAccepted) {
		t.Errorf("In array")
	}
}

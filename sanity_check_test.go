package main

import (
	"config-forge/utils"
	"testing"
)

func TestCheckArgumentsWithGoodArguments(t *testing.T) {
	arguments := []string{
		"",
		"apache",
		"test",
		"test.test",
		"dev/test/public",
	}
	if utils.CheckArguments(arguments) != "ok" {
		t.Errorf("Minimum one arguments don't pass the sanity check")
	}
}

func TestCheckArgumentsWithBadArguments(t *testing.T) {
	tests := []struct {
		command []string
		wanted  string
	}{
		{command: []string{"te", "te", "te", ".te"}},
		{command: []string{"te", "te", "te", ".te", "t"}},
		{command: []string{"te", "apache", "te", ".te", "t"}},
		{command: []string{"te", "apache", "test", ".te", "t"}},
		{command: []string{"te", "apache", "test", "test.test", "t"}},
		{command: []string{"te", "apache", "test", "test.test", "dev/test/public", "te"}},
	}

	for _, arguments := range tests {
		if utils.CheckArguments(arguments.command) == "ok" {
			t.Errorf("Minimum one arguments pass the sanity check")
		}
	}
}

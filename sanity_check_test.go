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
	if !utils.CheckArguments(arguments) {
		t.Errorf("Minimum one arguments don't pass the sanity check")
	}
}

func TestCheckArgumentsWithBadArguments(t *testing.T) {
	tests := []struct {
		command []string
		wanted  string
	}{
		{command: []string{"", "", "", "", ""}},
		{command: []string{"", "apache", "", "", ""}},
		{command: []string{"", "apache", "test", "", ""}},
		{command: []string{"", "apache", "test", "test.test", ""}},
		{command: []string{"", "apache", "test", "test.test", "dev/test/public", ""}},
	}

	for _, arguments := range tests {
		if utils.CheckArguments(arguments.command) {
			t.Errorf("Minimum one arguments pass the sanity check")
		}
	}
}

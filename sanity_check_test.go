package main

import (
	"config-forge/configs"
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

	sanityReturn := utils.CheckArguments(arguments)

	if sanityReturn != configs.SANITY_CHECK_PASSED {
		t.Errorf("Minimum one arguments don't pass the sanity check : %s", sanityReturn)
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
		if utils.CheckArguments(arguments.command) == configs.SANITY_CHECK_PASSED {
			t.Errorf("Minimum one arguments pass the sanity check")
		}
	}
}

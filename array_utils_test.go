package main

import (
	"config-forge/configs"
	"config-forge/utils"
	"testing"
)

func TestInArray(t *testing.T) {
	if !utils.InArray(configs.APACHE_INPUT, configs.ConfigAccepted) {
		t.Errorf("Not in array")
	}
}

func TestNotInArray(t *testing.T) {
	if utils.InArray(configs.BAD_INPUT, configs.ConfigAccepted) {
		t.Errorf("In array")
	}
}

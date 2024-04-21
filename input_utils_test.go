package main

import (
	"bytes"
	"config-forge/configs"
	"config-forge/utils"
	"testing"
)

func TestWrongAnalyseChoice(t *testing.T) {
	function := utils.AnalyseChoice(configs.BAD_INPUT)

	if function != configs.WRONG_CHOICE_RETURN {
		t.Errorf(configs.GOOD_CHOICE_ERROR)
	}
}

func TestGoodAnalyseChoice(t *testing.T) {
	function := utils.AnalyseChoice(configs.APACHE_INPUT)

	if function != configs.PREPARE_CONFIGURATION_RETURN {
		t.Errorf(configs.WRONG_CHOICE_ERROR)
	}
}

func TestListeningResponse(t *testing.T) {
	reader := bytes.NewReader([]byte(configs.APACHE_INPUT))

	response := utils.ListeningResponse(reader)

	if response != configs.APACHE_INPUT {
		t.Errorf("The response %q isn't the same of the input %q", configs.APACHE_INPUT, response)
	}
}

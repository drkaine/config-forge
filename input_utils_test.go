package main

import (
	"bytes"
	"testing"
)

func TestWrongAnalyseResponse(t *testing.T) {
	function := AnalyseResponse(BAD_INPUT)

	if function != WRONG_CHOICE_RETURN {
		t.Errorf(GOOD_CHOICE_ERROR)
	}
}

func TestGoodAnalyseResponse(t *testing.T) {
	function := AnalyseResponse(APACHE_INPUT)

	if function != PREPARE_CONFIGURATION_RETURN {
		t.Errorf(WRONG_CHOICE_ERROR)
	}
}

func TestListeningResponse(t *testing.T) {
	reader := bytes.NewReader([]byte(APACHE_INPUT))

	response := ListeningResponse(reader)

	if response != APACHE_INPUT {
		t.Errorf("The response %q isn't the same of the input %q", APACHE_INPUT, response)
	}
}

package main

import (
	"bytes"
	"testing"
)

func TestWrongAnalyseResponse(t *testing.T) {
	function := AnalyseResponse(BadInput)

	if function != WrongChoice {
		t.Errorf(GoodChoiceError)
	}
}

func TestGoodAnalyseResponse(t *testing.T) {
	function := AnalyseResponse(ApacheInput)

	if function != PrepareConfiguration {
		t.Errorf(WrongChoiceError)
	}
}

func TestListeningResponse(t *testing.T) {
	reader := bytes.NewReader([]byte(ApacheInput))

	response := ListeningResponse(reader)

	if response != ApacheInput {
		t.Errorf("The response %q isn't the same of the input %q", ApacheInput, response)
	}
}
